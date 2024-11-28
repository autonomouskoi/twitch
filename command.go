package twitch

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"google.golang.org/protobuf/proto"

	"github.com/autonomouskoi/akcore/bus"
)

var (
	oauthChallenge = rand.Int63()

	oauthScopes = []string{
		"bits:read",
		"channel:manage:redemptions",
		"channel:read:hype_train",
		"channel:read:subscriptions",
		"chat:edit",
		"chat:read",
		"moderator:manage:banned_users",
		"moderator:read:followers",
		"user:read:chat",
		"user:write:chat",
	}
)

func (t *Twitch) handleCommand(ctx context.Context) error {
	in := make(chan *bus.BusMessage, 16)
	t.bus.Subscribe(BusTopics_TWITCH_COMMAND.String(), in)
	go func() {
		<-ctx.Done()
		t.bus.Unsubscribe(BusTopics_TWITCH_COMMAND.String(), in)
		bus.Drain(in)
	}()
	for msg := range in {
		var resp *bus.BusMessage
		switch msg.Type {
		case int32(MessageTypeCommand_TYPE_COMMAND_GET_OAUTH_URL_REQ):
			resp = t.handleCommandGetOAuthURL(msg)
		case int32(MessageTypeCommand_TYPE_COMMAND_WRITE_PROFILE_REQ):
			resp = t.handleCommandWriteProfile(msg)
		case int32(MessageTypeCommand_TYPE_COMMAND_DELETE_PROFILE_REQ):
			resp = t.handleCommandDeleteProfile(msg)
		case int32(MessageTypeCommand_TYPE_COMMAND_CHAT_SET_CONFIG_REQ):
			resp = t.handleChatSetConfigRequest(msg)
		case int32(MessageTypeCommand_TYPE_COMMAND_EVENT_SET_CONFIG_REQ):
			resp = t.handleEventSubSetConfigRequest(msg)
		}
		if resp != nil {
			t.bus.SendReply(msg, resp)
		}
	}
	return ctx.Err()
}

func (t *Twitch) handleCommandGetOAuthURL(reqMsg *bus.BusMessage) *bus.BusMessage {
	query := url.Values{}
	query.Set("client_id", clientID)
	query.Set("redirect_uri", "http://localhost:8011/m/twitch/oauth.html")
	query.Set("response_type", "token")
	query.Set("scope", strings.Join(oauthScopes, " "))
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, rand.Int63())
	state := t.oauthStateToken(hex.EncodeToString(buf.Bytes()))
	query.Set("state", state)
	redirURL := url.URL{
		Scheme:   "https",
		Host:     "id.twitch.tv",
		Path:     "/oauth2/authorize",
		RawQuery: query.Encode(),
	}
	b, err := proto.Marshal(&GetOAuthURLResponse{
		Url: redirURL.String(),
	})
	msg := &bus.BusMessage{
		Topic: reqMsg.GetTopic(),
		Type:  int32(MessageTypeCommand_TYPE_COMMAND_GET_OAUTH_URL_RES),
	}
	if err != nil {
		t.Log.Error("marshalling", "type", "GetOAuthURLResposne", "error", err.Error())
		msg.Error = &bus.Error{
			Detail: proto.String("marshalling: " + err.Error()),
		}
		return msg
	}
	msg.Message = b
	return msg
}

func (t *Twitch) handleCommandWriteProfile(reqMsg *bus.BusMessage) *bus.BusMessage {
	wpr := &WriteProfileRequest{}
	msg := &bus.BusMessage{
		Topic: reqMsg.GetTopic(),
	}
	if err := proto.Unmarshal(reqMsg.GetMessage(), wpr); err != nil {
		t.Log.Error("unmarshalling", "type", "WriteProfileRequest", "error", err.Error())
		msg.Error = &bus.Error{
			Code:   int32(bus.CommonErrorCode_INVALID_TYPE),
			Detail: proto.String("unmarshalling: " + err.Error()),
		}
		return msg
	}

	profile := wpr.GetProfile()
	token := profile.GetToken()
	state := token.ClientId
	nonce, _, _ := strings.Cut(state, ":")
	if t.oauthStateToken(nonce) != state {
		t.Log.Debug("state validation failed", "state", state,
			"computed", t.oauthStateToken(nonce),
		)
		msg.Error = &bus.Error{
			Code:   int32(bus.CommonErrorCode_INVALID_TYPE),
			Detail: proto.String("State validation failed"),
		}
		return msg
	}
	token.ClientId = clientID
	client, err := newClient(token)
	if err != nil {
		t.Log.Error("creating client", "error", err.Error())
		msg.Error = &bus.Error{
			Code:   int32(bus.CommonErrorCode_INVALID_TYPE),
			Detail: proto.String("creating client: " + err.Error()),
		}
		return msg
	}
	valid, validationResp, err := client.ValidateToken(token.Access)
	profile.Name = validationResp.Data.Login
	profile.Token.UserId = validationResp.Data.UserID
	if err != nil {
		t.Log.Error("validating new token", "error", err.Error())
		msg.Error = &bus.Error{
			Code:   int32(bus.CommonErrorCode_INVALID_TYPE),
			Detail: proto.String("validating token: " + err.Error()),
		}
		return msg
	}
	if !valid {
		msg.Error = &bus.Error{
			Code:   int32(bus.CommonErrorCode_INVALID_TYPE),
			Detail: proto.String("invalid token"),
		}
		return msg
	}

	if err := t.addProfile(profile); err != nil {
		t.Log.Error("adding profile", "error", err.Error())
		msg.Error = &bus.Error{
			Code:   int32(bus.CommonErrorCode_INVALID_TYPE),
			Detail: proto.String("invalid profile: " + err.Error()),
		}
		return msg
	}
	t.lock.Lock()
	if t.cfg.Profiles == nil {
		t.cfg.Profiles = map[string]*Profile{}
	}
	t.cfg.Profiles[wpr.Profile.Name] = wpr.GetProfile()
	t.lock.Unlock()
	t.Log.Info("twitch profile saved", "name", wpr.Profile.Name)
	b, err := proto.Marshal(&WriteProfileResponse{})
	if err != nil {
		t.Log.Error("marshalling", "type", "WriteProfileResponse", "error", err.Error())
		return nil
	}
	msg.Type = int32(MessageTypeCommand_TYPE_COMMAND_WRITE_PROFILE_RESP)
	msg.Message = b
	return msg
}

func (t *Twitch) revokeToken(profileName string) {
	profile, present := t.cfg.Profiles[profileName]
	if !present {
		t.Log.Debug("profile not present to revoke", "name", profile)
		return
	}
	values := url.Values{}
	values.Set("client_id", clientID)
	values.Set("token", profile.Token.Access)
	req, err := http.NewRequest(
		http.MethodPost,
		"https://id.twitch.tv/oauth2/revoke",
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		t.Log.Error("creating revoke request", "error", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cache-Control", "no-store")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log.Error("sending revoke request", "error", err.Error())
		return
	}
	if resp.StatusCode != http.StatusOK {
		t.Log.Error("non-200 status revoking token", "code", resp.StatusCode, "status", resp.Status)
		return
	}
	t.Log.Info("revoked token", "name", profileName)
}

func (t *Twitch) oauthStateToken(nonce string) string {
	h := sha256.New()
	h.Write([]byte(nonce))
	binary.Write(h, binary.BigEndian, oauthChallenge)
	hash := h.Sum(nil)
	return fmt.Sprintf("%s:%s", nonce, hex.EncodeToString(hash[0:8]))
}

func (t *Twitch) handleCommandDeleteProfile(reqMsg *bus.BusMessage) *bus.BusMessage {
	dpr := &DeleteProfileRequest{}
	msg := &bus.BusMessage{
		Topic: reqMsg.GetTopic(),
		Type:  int32(MessageTypeCommand_TYPE_COMMAND_DELETE_PROFILE_RESP),
	}
	if err := proto.Unmarshal(reqMsg.GetMessage(), dpr); err != nil {
		t.Log.Error("unmarshalling", "type", "DeleteProfileRequest", "error", err.Error())
		msg.Error = &bus.Error{
			Code:   int32(bus.CommonErrorCode_INVALID_TYPE),
			Detail: proto.String("unmarshalling: " + err.Error()),
		}
		return msg
	}
	msg.Message, _ = proto.Marshal(&DeleteProfileResponse{})
	t.lock.Lock()
	defer t.lock.Unlock()
	if _, ok := t.cfg.Profiles[dpr.GetName()]; !ok {
		return msg
	}
	t.revokeToken(dpr.GetName())
	delete(t.cfg.Profiles, dpr.GetName())
	delete(t.clients, dpr.GetName())
	return msg
}

func (t *Twitch) handleChatSetConfigRequest(msg *bus.BusMessage) *bus.BusMessage {
	reply := &bus.BusMessage{
		Topic: msg.GetTopic(),
		Type:  msg.GetType() + 1,
	}
	scr := &ChatSetConfigRequest{}
	if reply.Error = t.UnmarshalMessage(msg, scr); reply.Error != nil {
		return reply
	}
	t.lock.Lock()
	t.cfg.ChatConfig = scr.Config
	t.lock.Unlock()
	/*
		if scr.Config.Enabled {
			t.startChat(scr.Config)
		} else {
			t.stopChat()
		}
	*/
	reply.Message, _ = proto.Marshal(&ChatSetConfigResponse{})
	return reply
}

func (t *Twitch) handleEventSubSetConfigRequest(reqMsg *bus.BusMessage) *bus.BusMessage {
	scr := &EventSubSetConfigRequest{}
	msg := &bus.BusMessage{
		Topic: reqMsg.GetTopic(),
		Type:  int32(MessageTypeCommand_TYPE_COMMAND_EVENT_SET_CONFIG_RESP),
	}
	if err := proto.Unmarshal(reqMsg.GetMessage(), scr); err != nil {
		t.Log.Error("unmarshalling", "type", "EventSubSetConfigRequest", "error", err.Error())
		msg.Error = &bus.Error{
			Code:   int32(bus.CommonErrorCode_INVALID_TYPE),
			Detail: proto.String("unmarshalling: " + err.Error()),
		}
		return msg
	}
	t.lock.Lock()
	t.cfg.EsConfig = scr.Config
	t.lock.Unlock()
	if scr.Config.Enabled {
		t.startEventSub()
	} else {
		t.stopEventSub()
	}
	msg.Message, _ = proto.Marshal(&ChatSetConfigResponse{})
	return msg
}
