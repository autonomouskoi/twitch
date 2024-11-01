package twitch

import (
	"context"
	"strings"
	"time"

	"github.com/nicklaw5/helix/v2"
	"google.golang.org/protobuf/proto"

	"github.com/autonomouskoi/akcore/bus"
)

func (t *Twitch) handleRequest(ctx context.Context) error {
	in := make(chan *bus.BusMessage, 16)
	t.bus.Subscribe(BusTopics_TWITCH_REQUEST.String(), in)
	go func() {
		<-ctx.Done()
		t.bus.Unsubscribe(BusTopics_TWITCH_REQUEST.String(), in)
		bus.Drain(in)
	}()
	for msg := range in {
		var reply *bus.BusMessage
		switch msg.Type {
		case int32(MessageTypeRequest_TYPE_REQUEST_LIST_PROFILES_REQ):
			reply = t.handleRequestListProfiles(msg)
		case int32(MessageTypeRequest_TYPE_REQUEST_GET_USER_REQ):
			reply = t.handleRequestGetUser(msg)
		case int32(MessageTypeRequest_TYPE_REQUEST_IRC_GET_CONFIG_REQ):
			reply = t.handleIRCGetConfigRequest(msg)
		case int32(MessageTypeRequest_TYPE_REQUEST_IRC_GET_STATUS_REQ):
			reply = t.handleIRCGetStatusRequest(msg)
		case int32(MessageTypeRequest_TYPE_REQUEST_EVENT_GET_CONFIG_REQ):
			reply = t.handleEventSubGetConfigRequest(msg)
		case int32(MessageTypeRequest_TYPE_REQUEST_EVENT_GET_STATUS_REQ):
			reply = t.handleEventSubGetStatusRequest(msg)
		}
		if reply != nil {
			t.bus.SendReply(msg, reply)
		}
	}
	return ctx.Err()
}

func (t *Twitch) handleRequestListProfiles(reqMsg *bus.BusMessage) *bus.BusMessage {
	t.lock.Lock()
	defer t.lock.Unlock()
	lpr := &ListProfilesResponse{}
	for profileName := range t.cfg.Profiles {
		lpr.Names = append(lpr.Names, profileName)
	}
	b, err := proto.Marshal(lpr)
	if err != nil {
		t.log.Error("marshalling", "type", "ListProfilesResponse", "error", err.Error())
		return nil
	}
	return &bus.BusMessage{
		Topic:   reqMsg.GetTopic(),
		Type:    int32(MessageTypeRequest_TYPE_REQUEST_LIST_PROFILES_RESP),
		Message: b,
	}
}

func (t *Twitch) handleRequestGetUser(reqMsg *bus.BusMessage) *bus.BusMessage {
	gur := &GetUserRequest{}
	if err := proto.Unmarshal(reqMsg.GetMessage(), gur); err != nil {
		t.log.Error("unmarshalling", "type", "GetUserRequest", "error", err.Error())
		return nil
	}
	resp := &bus.BusMessage{
		Topic: reqMsg.GetTopic(),
		Type:  int32(MessageTypeRequest_TYPE_REQUEST_GET_USER_RESP),
	}
	t.lock.Lock()
	client, present := t.clients[gur.Profile]
	t.lock.Unlock()
	if !present {
		resp.Error = errInvalidClient
		return resp
	}

	guResp := &GetUserResponse{
		Login: gur.Login,
	}

	login := strings.TrimPrefix(strings.TrimSpace(gur.Login), "@")
	user, present := t.cacheUsers.Get(login)
	if present {
		guResp.User = user
	} else {
		uResp, err := client.GetUsers(&helix.UsersParams{Logins: []string{login}})
		err = extractError(err, uResp.ResponseCommon)
		if err != nil {
			resp.Error = &bus.Error{
				Detail:         proto.String(err.Error()),
				NotCommonError: true,
			}
			return resp
		}
		if len(uResp.Data.Users) == 0 {
			resp.Error = &bus.Error{
				Code: int32(bus.CommonErrorCode_NOT_FOUND),
			}
			return resp
		}
		user = &User{
			Id:              uResp.Data.Users[0].ID,
			Login:           uResp.Data.Users[0].Login,
			DisplayName:     uResp.Data.Users[0].DisplayName,
			Type:            uResp.Data.Users[0].Type,
			BroadcasterType: uResp.Data.Users[0].BroadcasterType,
			Description:     uResp.Data.Users[0].Description,
			ProfileImageUrl: uResp.Data.Users[0].ProfileImageURL,
			OfflineImageUrl: uResp.Data.Users[0].OfflineImageURL,
			ViewCount:       uint32(uResp.Data.Users[0].ViewCount),
			Email:           uResp.Data.Users[0].Email,
			CreatedAt:       uint32(uResp.Data.Users[0].CreatedAt.Unix()),
		}
		t.cacheUsers.Set(login, user)
		guResp.User = user
	}

	b, err := proto.Marshal(guResp)
	if err == nil {
		resp.Message = b
	} else {
		resp.Error = &bus.Error{}
		t.log.Error("marshalling", "type", "GetUserResponse")
	}
	return resp
}

func (t *Twitch) handleIRCGetConfigRequest(reqMsg *bus.BusMessage) *bus.BusMessage {
	cfg := &IRCConfig{}
	if t.cfg != nil && t.cfg.IrcConfig != nil {
		t.lock.Lock()
		cfg = t.cfg.IrcConfig
		t.lock.Unlock()
	}
	b, err := proto.Marshal(&IRCGetConfigResponse{
		Config: cfg,
	})
	resp := &bus.BusMessage{
		Topic: reqMsg.Topic,
		Type:  int32(MessageTypeRequest_TYPE_REQUEST_IRC_GET_CONFIG_RESP),
	}
	if err != nil {
		t.log.Error("marshalling", "type", "IRCGetConfigResponse", "error", err.Error())
		resp.Error = &bus.Error{
			Detail: proto.String("marshalling IRCGetConfigResponse: " + err.Error()),
		}
		return resp
	}
	resp.Message = b
	return resp
}

func (t *Twitch) handleIRCGetStatusRequest(reqMsg *bus.BusMessage) *bus.BusMessage {
	b, _ := proto.Marshal(&IRCGetStatusResponse{
		Status: t.chat.status,
		Detail: t.chat.statusDetail,
	})
	return &bus.BusMessage{
		Topic:   reqMsg.Topic,
		Type:    int32(MessageTypeRequest_TYPE_REQUEST_IRC_GET_STATUS_RESP),
		Message: b,
	}
}

func (t *Twitch) handleEventSubGetConfigRequest(reqMsg *bus.BusMessage) *bus.BusMessage {
	cfg := &EventSubConfig{}
	if t.cfg != nil && t.cfg.EsConfig != nil {
		t.lock.Lock()
		cfg = t.cfg.EsConfig
		t.lock.Unlock()
	}
	b, err := proto.Marshal(&EventSubGetConfigResponse{
		Config: cfg,
	})
	resp := &bus.BusMessage{
		Topic: reqMsg.Topic,
		Type:  int32(MessageTypeRequest_TYPE_REQUEST_EVENT_GET_CONFIG_RESP),
	}
	if err != nil {
		t.log.Error("marshalling", "type", "EventSubGetConfigResponse", "error", err.Error())
		resp.Error = &bus.Error{
			Detail: proto.String("marshalling EventSubGetConfigResponse: " + err.Error()),
		}
		return resp
	}
	resp.Message = b
	return resp
}

func (t *Twitch) handleEventSubGetStatusRequest(reqMsg *bus.BusMessage) *bus.BusMessage {
	b, _ := proto.Marshal(&EventSubGetStatusResponse{
		Status: t.eventSub.status,
		Detail: t.eventSub.statusDetail,
	})
	return &bus.BusMessage{
		Topic:   reqMsg.Topic,
		Type:    int32(MessageTypeRequest_TYPE_REQUEST_EVENT_GET_STATUS_RESP),
		Message: b,
	}
}

func GetUser(ctx context.Context, b *bus.Bus, twitchProfile, login string) (*User, error) {
	msg := &bus.BusMessage{
		Topic: BusTopics_TWITCH_REQUEST.String(),
		Type:  int32(MessageTypeRequest_TYPE_REQUEST_GET_USER_REQ),
	}
	var err error
	msg.Message, err = proto.Marshal(&GetUserRequest{
		Profile: twitchProfile,
		Login:   login,
	})
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	reply := b.WaitForReply(ctx, msg)
	cancel()
	if reply.Error != nil {
		return nil, err
	}
	gur := &GetUserResponse{}
	if err := proto.Unmarshal(reply.GetMessage(), gur); err != nil {
		return nil, err
	}
	return gur.GetUser(), nil
}
