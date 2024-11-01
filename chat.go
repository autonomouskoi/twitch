package twitch

import (
	"context"
	"fmt"
	"log/slog"

	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"

	"github.com/autonomouskoi/akcore/bus"
)

type chat struct {
	bus          *bus.Bus
	cancel       func()
	cfg          *IRCConfig
	eg           errgroup.Group
	log          *slog.Logger
	parentCtx    context.Context
	status       ChatStatus
	statusDetail string
}

func (t *Twitch) handleChat(ctx context.Context) error {
	if t.cfg.IrcConfig != nil && t.cfg.IrcConfig.Enabled {
		t.startChat(t.cfg.IrcConfig)
	}
	<-ctx.Done()
	t.stopChat()
	return t.chat.eg.Wait()
}

func (t *Twitch) startChat(cfg *IRCConfig) {
	c := t.chat
	c.cfg = cfg
	if c.status == ChatStatus_CHAT_STATUS_CONNECTED {
		return
	}
	ctx, cancel := context.WithCancel(c.parentCtx)
	c.cancel = cancel
	client, present := t.clients[cfg.Profile]
	if !present {
		t.log.Error("invalid profile for chat", "profile", cfg.Profile)
		return
	}
	accessToken := client.GetUserAccessToken()
	c.eg.Go(func() error {
		c.setStatus(ChatStatus_CHAT_STATUS_UNKNOWN, "Connecting")

		ircC := newIRCClient(c.cfg, c.bus, c.log, accessToken)
		if err := ircC.ConnectContext(ctx); err != nil {
			c.setStatus(ChatStatus_CHAT_STATUS_ERROR, err.Error())
			return fmt.Errorf("connecting to IRC: %w", err)
		}
		c.setStatus(ChatStatus_CHAT_STATUS_CONNECTED, "Connected!")
		c.bus.HandleTypes(ctx, BusTopics_TWITCH_CHAT_REQUEST.String(), 4,
			map[int32]bus.MessageHandler{
				int32(MessageTypeTwitchChatRequest_TWITCH_CHAT_REQUEST_TYPE_SEND_REQ): func(msg *bus.BusMessage) *bus.BusMessage {
					reply := &bus.BusMessage{
						Topic: msg.Topic,
						Type:  msg.Type + 1,
					}
					cmo := &TwitchChatRequestSendRequest{}
					if err := proto.Unmarshal(msg.GetMessage(), cmo); err != nil {
						c.log.Error("unmarshalling", "type", "TwitchChatRequestSendRequest")
						reply.Error = &bus.Error{
							Code:   int32(bus.CommonErrorCode_INVALID_TYPE),
							Detail: proto.String("marshalling TwitchChatRequestSendRequest: " + err.Error()),
						}
						return reply
					}
					text := cmo.Text
					if c.cfg.MessagePrefix != "" {
						text = c.cfg.MessagePrefix + text
					}
					ircC.Privmsg("#"+c.cfg.GetProfile(), text)
					reply.Message, _ = proto.Marshal(&TwitchChatRequestSendResponse{})
					return reply
				},
			},
			nil,
		)

		if err := ircC.Close(); err != nil {
			return fmt.Errorf("closing IRC: %w", err)
		}
		c.setStatus(ChatStatus_CHAT_STATUS_OFF, "Disconnected")

		return nil
	})
}

func (t *Twitch) stopChat() {
	if t.chat.cancel != nil {
		t.chat.cancel()
	}
}

func (c *chat) setStatus(status ChatStatus, detail string) {
	c.status = status
	c.statusDetail = detail
	b, _ := proto.Marshal(&ChatEvent{
		Type:   ChatEventType_EVENT_TYPE_STATUS,
		Detail: detail,
		Status: c.status,
	})
	c.bus.Send(&bus.BusMessage{
		Topic:   BusTopics_TWITCH_CHAT_EVENT.String(),
		Message: b,
	})
}
