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
	t.bus.HandleTypes(ctx, BusTopics_TWITCH_REQUEST.String(), 4,
		map[int32]bus.MessageHandler{
			int32(MessageTypeRequest_TYPE_REQUEST_LIST_PROFILES_REQ):    t.handleRequestListProfiles,
			int32(MessageTypeRequest_TYPE_REQUEST_GET_USER_REQ):         t.handleRequestGetUser,
			int32(MessageTypeRequest_TYPE_REQUEST_CHAT_GET_CONFIG_REQ):  t.handleChatGetConfigRequest,
			int32(MessageTypeRequest_TYPE_REQUEST_EVENT_GET_CONFIG_REQ): t.handleEventSubGetConfigRequest,
			int32(MessageTypeRequest_TYPE_REQUEST_EVENT_GET_STATUS_REQ): t.handleEventSubGetStatusRequest,
			int32(MessageTypeRequest_TYPE_REQUEST_SEND_SHOUTOUT_REQ):    t.handleRequestSendShoutout,
		},
		nil,
	)
	return nil
}

func (t *Twitch) handleRequestListProfiles(reqMsg *bus.BusMessage) *bus.BusMessage {
	reply := &bus.BusMessage{
		Topic: reqMsg.GetTopic(),
		Type:  reqMsg.GetType() + 1,
	}
	t.lock.Lock()
	defer t.lock.Unlock()
	lpr := &ListProfilesResponse{}
	for profileName, profile := range t.cfg.Profiles {
		lpr.Names = append(lpr.Names, profileName)
		lpr.Profiles = append(lpr.Profiles, &ListProfilesResponse_ProfileListing{
			Name:    profileName,
			UserId:  profile.GetUserId(),
			Scopes:  profile.GetToken().GetScopes(),
			Expires: profile.GetToken().GetExpires(),
		})
	}
	t.MarshalMessage(reply, lpr)
	return reply
}

func (t *Twitch) handleRequestGetUser(reqMsg *bus.BusMessage) *bus.BusMessage {
	reply := &bus.BusMessage{
		Topic: reqMsg.GetTopic(),
		Type:  reqMsg.GetType() + 1,
	}
	gur := &GetUserRequest{}
	if reply.Error = t.UnmarshalMessage(reqMsg, gur); reply.Error != nil {
		return reply
	}
	t.lock.Lock()
	client, present := t.clients[gur.Profile]
	t.lock.Unlock()
	if !present {
		reply.Error = errInvalidClient
		return reply
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
		if err != nil {
			reply.Error = &bus.Error{
				Detail: proto.String("getting user: " + err.Error()),
			}
			return reply
		}
		err = extractError(err, uResp.ResponseCommon)
		if err != nil {
			reply.Error = &bus.Error{
				Detail:         proto.String(err.Error()),
				NotCommonError: true,
			}
			return reply
		}
		if len(uResp.Data.Users) == 0 {
			reply.Error = &bus.Error{
				Code: int32(bus.CommonErrorCode_NOT_FOUND),
			}
			return reply
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

	t.MarshalMessage(reply, guResp)
	return reply
}

func (t *Twitch) handleChatGetConfigRequest(reqMsg *bus.BusMessage) *bus.BusMessage {
	reply := &bus.BusMessage{
		Topic: reqMsg.GetTopic(),
		Type:  reqMsg.GetType() + 1,
	}
	cfg := &ChatConfig{}
	if t.cfg != nil && t.cfg.ChatConfig != nil {
		t.lock.Lock()
		cfg = t.cfg.ChatConfig
		t.lock.Unlock()
	}
	t.MarshalMessage(reply, &ChatGetConfigResponse{Config: cfg})
	return reply
}

func (t *Twitch) handleEventSubGetConfigRequest(reqMsg *bus.BusMessage) *bus.BusMessage {
	reply := &bus.BusMessage{
		Topic: reqMsg.GetTopic(),
		Type:  reqMsg.GetType() + 1,
	}
	cfg := &EventSubConfig{}
	if t.cfg != nil && t.cfg.EsConfig != nil {
		t.lock.Lock()
		cfg = t.cfg.EsConfig
		t.lock.Unlock()
	}
	t.MarshalMessage(reply, &EventSubGetConfigResponse{
		Config: cfg,
	})
	return reply
}

func (t *Twitch) handleEventSubGetStatusRequest(reqMsg *bus.BusMessage) *bus.BusMessage {
	b, _ := proto.Marshal(&EventSubGetStatusResponse{
		Status: t.eventSub.status,
		Detail: t.eventSub.statusDetail,
	})
	return &bus.BusMessage{
		Topic:   reqMsg.GetTopic(),
		Type:    reqMsg.GetType() + 1,
		Message: b,
	}
}

// GetUser is a convenience function for GET_USER_REQ
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
		return nil, reply.Error
	}
	gur := &GetUserResponse{}
	if err := proto.Unmarshal(reply.GetMessage(), gur); err != nil {
		return nil, err
	}
	return gur.GetUser(), nil
}
