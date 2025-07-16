package twitch

import (
	"context"

	"google.golang.org/protobuf/proto"

	"github.com/autonomouskoi/akcore/bus"
	"github.com/nicklaw5/helix/v2"
)

// handle chat request messages
func (t *Twitch) handleChat(ctx context.Context) error {
	t.bus.HandleTypes(ctx, BusTopics_TWITCH_CHAT_REQUEST.String(), 8,
		map[int32]bus.MessageHandler{
			int32(MessageTypeTwitchChatRequest_TWITCH_CHAT_REQUEST_TYPE_SEND_REQ): t.handleChatSendRequest,
		},
		nil)
	return nil
}

// handle a request to send a message to twitch chat
func (t *Twitch) handleChatSendRequest(msg *bus.BusMessage) *bus.BusMessage {
	reply := &bus.BusMessage{
		Topic: msg.GetTopic(),
		Type:  msg.GetType() + 1,
	}
	/*
		if !t.cfg.ChatConfig.Enabled {
			reply.Error = &bus.Error{
				Code:        int32(bus.CommonErrorCode_INVALID_TYPE),
				UserMessage: proto.String("not enabled"),
			}
			return reply
		}
	*/
	sr := &TwitchChatRequestSendRequest{}
	if reply.Error = t.UnmarshalMessage(msg, sr); reply.Error != nil {
		return reply
	}
	if sr.Text == "" {
		return nil
	}
	profile := defaultValue(sr.GetProfile(), t.cfg.GetChatConfig().GetProfile())
	t.lock.Lock()
	client := t.clients[profile]
	t.lock.Unlock()
	if client == nil {
		reply.Error = &bus.Error{
			Code:        int32(bus.CommonErrorCode_NOT_FOUND),
			UserMessage: proto.String("no profile: " + profile),
		}
		return reply
	}
	channelClient := t.clients[defaultValue(sr.GetChannel(), client.UserID())]
	if channelClient == nil {
		reply.Error = &bus.Error{
			Code:        int32(bus.CommonErrorCode_NOT_FOUND),
			UserMessage: proto.String("no profile: " + profile),
		}
		return reply
	}
	resp, err := client.SendChatMessage(&helix.SendChatMessageParams{
		BroadcasterID: channelClient.UserID(),
		SenderID:      client.UserID(),
		Message:       t.cfg.ChatConfig.MessagePrefix + sr.GetText(),
	})
	if err != nil {
		reply.Error = &bus.Error{
			Detail: proto.String("sending message: " + err.Error()),
		}
		return reply
	}
	if err := extractError(err, resp.ResponseCommon); err != nil {
		reply.Error = &bus.Error{
			Detail: proto.String("sending message: " + err.Error()),
		}
		return reply
	}
	t.MarshalMessage(reply, &TwitchChatRequestSendResponse{})
	return reply
}
