package twitch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/nicklaw5/helix/v2"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"

	"github.com/autonomouskoi/akcore/bus"
	"github.com/autonomouskoi/akcore/modules/modutil"
	"github.com/autonomouskoi/datastruct/slices"
)

const (
	defaultWebsocketURL = "wss://eventsub.wss.twitch.tv/ws"

	callbackMethod = "websocket"
)

var (
	topicEventSubEvent string
)

func init() {
	topicEventSubEvent = MessageTypeEventSub_TYPE_EVENTSUB_EVENT.String()
}

type twitchClient interface {
	UserID() string
	CreateEventSubSubscription(payload *helix.EventSubSubscription) (*helix.EventSubSubscriptionsResponse, error)
}

func (t *Twitch) handleEventSub(ctx context.Context) error {
	if t.cfg.EsConfig != nil && t.cfg.EsConfig.Enabled {
		t.startEventSub()
	}
	<-ctx.Done()
	t.stopEventSub()
	return t.eventSub.eg.Wait()
}

func (t *Twitch) startEventSub() {
	if t.eventSub.cancel != nil {
		return
	}
	t.eventSub.eg.Go(func() error {
		var client *client
		if t.cfg != nil && t.cfg.EsConfig != nil {
			client = t.clients[t.cfg.EsConfig.Profile]
		}
		if client == nil {
			return errors.New("invalid profile")
		}
		t.eventSub.twitch = client

		ctx, cancel := context.WithCancel(t.eventSub.parentCtx)
		t.eventSub.cancel = cancel

		t.eventSub.setStatus(EventSubStatus_EVENT_SUB_STATUS_UNKNOWN, "Connecting")
		if err := t.eventSub.start(ctx); err != nil {
			t.eventSub.setStatus(EventSubStatus_EVENT_SUB_STATUS_ERROR, err.Error())
			return fmt.Errorf("staring: %w", err)
		}
		return nil
	})
}

func (t *Twitch) stopEventSub() {
	if t.eventSub.cancel != nil {
		t.eventSub.cancel()
	}
}

type eventSub struct {
	modutil.ModuleBase
	bus          *bus.Bus
	c            *websocket.Conn
	parentCtx    context.Context
	twitch       twitchClient
	seen         *seenIDs
	cancel       func()
	eg           errgroup.Group
	status       EventSubStatus
	statusDetail string
}

func (es *eventSub) connect(ctx context.Context, websocketURL string) error {
	if es.c != nil {
		es.c.CloseNow()
	}
	es.Log.Debug("dialing eventsub websocket", "url", websocketURL)
	c, _, err := websocket.Dial(ctx, websocketURL, nil)
	if err != nil {
		return fmt.Errorf("dialing websocket: %w", err)
	}

	defer func() {
		// if es.c is nil, we didn't get set up properly. Close the websocket
		if es.c == nil {
			es.Log.Debug("no client, closing")
			c.CloseNow()
		}
	}()

	var msg Message
	if err := wsjson.Read(ctx, c, &msg); err != nil {
		return fmt.Errorf("reading welcome message: %w", err)
	}

	welcome, err := msg.AsSessionWelcome()
	if err != nil {
		return fmt.Errorf("getting welcome message: %w", err)
	}
	es.Log.Debug("got welcome", "msg", *welcome.Session)
	es.setStatus(EventSubStatus_EVENT_SUB_STATUS_CONNECTED, "Connected!")
	transport := helix.EventSubTransport{
		Method:    callbackMethod,
		SessionID: welcome.Session.ID,
	}

	for eventType, eventVersion := range map[string]string{
		EventTypeChannelPointRedeem: EventVersionChannelPointRedeem,
		EventTypeCheer:              EventVersionCheer,
		EventTypeHypeTrainBegin:     EventVersionHypTrainBegin,
		EventTypeHypeTrainProgress:  EventVersionHypTrainProgress,
		EventTypeHypeTrainEnd:       EventVersionHypTrainEnd,
		EventTypeSubscription:       EventVersionSubscription,
	} {
		_, err := es.twitch.CreateEventSubSubscription(&helix.EventSubSubscription{
			Type:    eventType,
			Version: eventVersion,
			Condition: helix.EventSubCondition{
				BroadcasterUserID: es.twitch.UserID(),
			},
			Transport: transport,
		})
		if err != nil {
			return fmt.Errorf("creating %s/%s sub: %w", eventType, eventVersion, err)
		}
	}
	_, err = es.twitch.CreateEventSubSubscription(&helix.EventSubSubscription{
		Type:    EventTypeRaid,
		Version: EventVersionRaid,
		Condition: helix.EventSubCondition{
			ToBroadcasterUserID: es.twitch.UserID(),
		},
		Transport: transport,
	})
	if err != nil {
		return fmt.Errorf("creating %s/%s sub: %w", EventTypeRaid, EventVersionRaid, err)
	}
	_, err = es.twitch.CreateEventSubSubscription(&helix.EventSubSubscription{
		Type:    EventTypeFollow,
		Version: EventVersionFollow,
		Condition: helix.EventSubCondition{
			BroadcasterUserID: es.twitch.UserID(),
			ModeratorUserID:   es.twitch.UserID(),
		},
		Transport: transport,
	})
	if err != nil {
		return fmt.Errorf("creating %s/%s sub: %w", EventTypeFollow, EventVersionFollow, err)
	}
	_, err = es.twitch.CreateEventSubSubscription(&helix.EventSubSubscription{
		Type:    helix.EventSubTypeChannelChatMessage,
		Version: EventVersionChannelChatMessage,
		Condition: helix.EventSubCondition{
			BroadcasterUserID: es.twitch.UserID(),
			UserID:            es.twitch.UserID(),
		},
		Transport: transport,
	})
	if err != nil {
		return fmt.Errorf("creating %s/%s sub: %w", helix.EventSubTypeChannelChatMessage, EventVersionChannelChatMessage, err)
	}
	es.c = c
	return nil
}

func (es *eventSub) start(ctx context.Context) error {
	defer func() {
		es.cancel()
		es.cancel = nil
	}()
	websocketURL := defaultWebsocketURL
	if v := os.Getenv("TWITCH_WS_URL"); v != "" {
		websocketURL = v
	}
	if err := es.connect(ctx, websocketURL); err != nil {
		return err
	}

	es.handleLoop(ctx)

	es.close()
	es.setStatus(EventSubStatus_EVENT_SUB_STATUS_OFF, "Disconnected")

	return nil
}

func (es *eventSub) close() {
	es.Log.Debug("closing on request")
	if err := es.c.Close(websocket.StatusNormalClosure, ""); err != nil {
		es.Log.Error("closing websocket", "error", err)
	}
}

func (es *eventSub) handleLoop(ctx context.Context) {
	for {
		if ctx.Err() != nil {
			break
		}
		var msg Message
		if err := wsjson.Read(ctx, es.c, &msg); err != nil {
			if !errors.Is(err, context.DeadlineExceeded) {
				es.Log.Error("receiving", "error", err)
				break
			}
			continue
		}
		if es.seen.seenID(msg.Metadata.ID) {
			es.Log.Debug("duplicate event ID", "id", msg.Metadata.ID)
			continue
		}
		switch msg.Metadata.Type {
		case MessageTypeNotification:
			notification := &Notification{}
			if err := json.Unmarshal(msg.Payload, &notification); err != nil {
				es.Log.Error("unmarshalling notification", "error", err)
				continue
			}
			go es.handleNotification(notification)
		case MessageTypeSessionKeepalive:
			// TODO: watch for keepalives and bail when they stop
		case MessageTypeSessionReconnect:
			welcome, err := msg.AsSessionWelcome()
			if err != nil {
				es.Log.Error("unmarshalling reconnect", "error", err)
				continue
			}
			es.Log.Info("received reconnect")
			if err := es.connect(ctx, *welcome.Session.ReconnectURL); err != nil {
				es.Log.Error("reconnecting", "error", err)
				continue
			}
		default:
			es.Log.Debug("unhandled message", "type", msg.Metadata.Type, "payload", string(msg.Payload))
		}
	}
}

func (es *eventSub) handleNotification(n *Notification) {
	var msg *bus.BusMessage
	switch n.Subscription.Type {
	case helix.EventSubTypeChannelChatMessage:
		msg = es.handleChannelChatMessage(n.Event)
	case EventTypeChannelPointRedeem:
		msg = es.handleChannePointRedeem(n.Event)
	case EventTypeCheer:
		msg = es.handleChannelCheer(n.Event)
	case EventTypeFollow:
		msg = es.handleChannelFollow(n.Event)
	case EventTypeRaid:
		msg = es.handleChannelRaid(n.Event)
	default:
		es.Log.Info("unhandled notification", "type", n.Subscription.Type)
	}
	if msg != nil {
		es.bus.Send(msg)
	}
}

func (es *eventSub) handleChannelChatMessage(eventJSON json.RawMessage) *bus.BusMessage {
	msg := &bus.BusMessage{
		Topic: BusTopics_TWITCH_EVENTSUB_EVENT.String(),
		Type:  int32(MessageTypeEventSub_TYPE_CHANNEL_CHAT_MESSAGE),
	}
	event := &helix.EventSubChannelChatMessageEvent{}
	if err := json.Unmarshal(eventJSON, event); err != nil {
		es.Log.Error("unmarshalling chat message event", "error", err.Error())
		return nil
	}
	ecm := &EventChannelChatMessage{
		Id: event.MessageID,
		Broadcaster: &EventUser{
			Id:    event.BroadcasterUserID,
			Login: event.BroadcasterUserLogin,
			Name:  event.BroadcasterUserName,
		},
		Chatter: &EventUser{
			Id:    event.ChatterUserID,
			Login: event.ChatterUserLogin,
			Name:  event.ChatterUserName,
		},
		Message: &ChatMessage{
			Text: event.Message.Text,
			Fragments: slices.Map(event.Message.Fragments,
				func(in helix.EventSubChatMessageFragment) *ChatMessageFragment {
					cmf := &ChatMessageFragment{
						Text: in.Text,
					}
					switch in.Type {
					case helix.EventSubChatMessageFragmentTypeText:
						cmf.Type = ChatMessageFragmentType_text
					case helix.EventSubChatMessageFragmentTypeCheermote:
						cmf.Type = ChatMessageFragmentType_cheermote
						cmf.Cheermote = &ChatMessageCheermote{
							Prefix: in.Cheermote.Prefix,
							Bits:   in.Cheermote.Bits,
							Tier:   int32(in.Cheermote.Tier),
						}
					case helix.EventSubChatMessageFragmentTypeEmote:
						cmf.Type = ChatMessageFragmentType_emote
						cmf.Emote = &ChatMessageEmote{
							Id:         in.Emote.ID,
							EmoteSetId: in.Emote.EmoteSetID,
							OwnerId:    in.Emote.OwnerID,
							Format:     in.Emote.Format,
						}
					case helix.EventSubChatMessageFragmentTypeMention:
						cmf.Type = ChatMessageFragmentType_mention
						cmf.Mention = &EventUser{
							Id:    in.Mention.UserID,
							Login: in.Mention.UserLogin,
							Name:  in.Mention.UserName,
						}
					}
					return cmf
				}),
		},
		MessageType: func(t helix.EventSubChatMessageType) ChatMessageType {
			switch t {
			case helix.EventSubChatMessageTypeText:
				return ChatMessageType_chat_message_type_text
			case helix.EventSubChatMessageTypeChannelPointsHighlighted:
				return ChatMessageType_chat_message_type_channel_points_highlighted
			case helix.EventSubChatMessageTypeChannelPointsSubOnly:
				return ChatMessageType_chat_message_type_channel_points_sub_only
			case helix.EventSubChatMessageTypeUserIntro:
				return ChatMessageType_chat_message_type_user_intro
			}
			return -1
		}(event.MessageType),
		Badges: slices.Map(event.Badges, func(badge helix.EventSubChatBadge) *ChatBadge {
			return &ChatBadge{
				SetId: badge.SetID,
				Id:    badge.ID,
				Info:  badge.Info,
			}
		}),
		Cheer: func(cheer helix.EventSubChatMessageCheer) *ChatMessageCheer {
			if cheer.Bits == 0 {
				return nil
			}
			return &ChatMessageCheer{Bits: cheer.Bits}
		}(event.Cheer),
		Color: event.Color,
		Reply: func(reply helix.EventSubChatMessageReply) *ChatMessageReply {
			if reply.ParentMessageID == "" {
				return nil
			}
			return &ChatMessageReply{
				ParentMessageId:   event.Reply.ParentMessageID,
				ParentMessageBody: event.Reply.ParentMessageBody,
				ParentUser: &EventUser{
					Id:    event.Reply.ParentUserID,
					Login: event.Reply.ParentUserLogin,
					Name:  event.Reply.ParentUserName,
				},
				ThreadMessageId: event.Reply.ThreadMessageID,
				ThreadUser: &EventUser{
					Id:    event.Reply.ThreadUserID,
					Login: event.Reply.ThreadUserLogin,
					Name:  event.Reply.ThreadUserName,
				},
			}
		}(event.Reply),
		ChannelPointsCustomRewardId: event.ChannelPointsCustomRewardID,
		Enrichments: &EventChannelChatMessage_Enrichments{
			IsMod: slices.Matches(event.Badges, func(b helix.EventSubChatBadge) bool {
				return b.SetID == "broadcaster" || b.SetID == "moderator"
			}),
		},
	}

	es.MarshalMessage(msg, ecm)
	if msg.Error != nil {
		return nil
	}
	return msg
}

func (es *eventSub) handleChannePointRedeem(eventJSON json.RawMessage) *bus.BusMessage {
	msg := &bus.BusMessage{
		Topic: topicEventSubEvent,
		Type:  int32(MessageTypeEventSub_TYPE_CHANNEL_POINT_CUSTOM_REDEEM),
	}
	event := &helix.EventSubChannelPointsCustomRewardRedemptionEvent{}
	if err := json.Unmarshal(eventJSON, event); err != nil {
		es.Log.Error("unmarshalling channel point redeem event", "error", err.Error())
		return nil
	}
	es.MarshalMessage(msg, &EventChannelPointsCustomRewardRedemption{
		Id: event.ID,
		Broadcaster: &EventUser{
			Id:    event.BroadcasterUserID,
			Login: event.BroadcasterUserLogin,
			Name:  event.BroadcasterUserName,
		},
		User: &EventUser{
			Id:    event.UserID,
			Login: event.UserLogin,
			Name:  event.UserName,
		},
		Input:  event.UserInput,
		Status: event.Status,
		Reward: &Reward{
			Id:     event.Reward.ID,
			Title:  event.Reward.Title,
			Cost:   int32(event.Reward.Cost),
			Prompt: event.Reward.Prompt,
		},
		RedeemedAt: event.RedeemedAt.Unix(),
	})
	if msg.Error != nil {
		return nil
	}
	return msg
}

func (es *eventSub) handleChannelCheer(eventJSON json.RawMessage) *bus.BusMessage {
	msg := &bus.BusMessage{
		Topic: topicEventSubEvent,
		Type:  int32(MessageTypeEventSub_TYPE_CHANNEL_CHEER),
	}
	event := &helix.EventSubChannelCheerEvent{}
	if err := json.Unmarshal(eventJSON, event); err != nil {
		es.Log.Error("unmarshalling cheer event", "error", err.Error())
		return nil
	}
	es.MarshalMessage(msg,
		&EventChannelCheer{
			IsAnonymous: &event.IsAnonymous,
			From: &EventUser{
				Id:    event.UserID,
				Login: event.UserLogin,
				Name:  event.UserName,
			},
			Broadcaster: &EventUser{
				Id:    event.BroadcasterUserID,
				Login: event.BroadcasterUserLogin,
				Name:  event.BroadcasterUserName,
			},
			Message: &event.Message,
			Bits:    uint32(event.Bits),
		})
	if msg.Error != nil {
		return nil
	}
	return msg
}

func (es *eventSub) handleChannelFollow(eventJSON json.RawMessage) *bus.BusMessage {
	msg := &bus.BusMessage{
		Topic: topicEventSubEvent,
		Type:  int32(MessageTypeEventSub_TYPE_CHANNEL_FOLLOW),
	}
	event := &helix.EventSubChannelFollowEvent{}
	if err := json.Unmarshal(eventJSON, event); err != nil {
		es.Log.Error("unmarshalling follow event", "error", err.Error())
		return nil
	}
	es.MarshalMessage(msg, &EventChannelFollow{
		Broadcaster: &EventUser{
			Id:    event.BroadcasterUserID,
			Login: event.BroadcasterUserLogin,
			Name:  event.BroadcasterUserName,
		},
		Follower: &EventUser{
			Id:    event.UserID,
			Login: event.UserLogin,
			Name:  event.UserName,
		},
		At: event.FollowedAt.Unix(),
	})
	if msg.Error != nil {
		return nil
	}
	return msg
}

func (es *eventSub) handleChannelRaid(eventJSON json.RawMessage) *bus.BusMessage {
	msg := &bus.BusMessage{
		Topic: topicEventSubEvent,
		Type:  int32(MessageTypeEventSub_TYPE_CHANNEL_RAID),
	}
	event := &helix.EventSubChannelRaidEvent{}
	if err := json.Unmarshal(eventJSON, event); err != nil {
		es.Log.Error("unmarshalling raid event", "error", err.Error())
		return nil
	}
	ecr := &EventChannelRaid{
		FromBroadcaster: &EventUser{
			Id:    event.FromBroadcasterUserID,
			Login: event.FromBroadcasterUserLogin,
			Name:  event.FromBroadcasterUserName,
		},
		ToBroadcaster: &EventUser{
			Id:    event.ToBroadcasterUserID,
			Login: event.ToBroadcasterUserLogin,
			Name:  event.ToBroadcasterUserName,
		},
		Viewers: uint32(event.Viewers),
	}
	if es.MarshalMessage(msg, ecr); msg.Error != nil {
		return nil
	}
	return msg
}

func (es *eventSub) setStatus(status EventSubStatus, detail string) {
	es.status = status
	es.statusDetail = detail
	b, _ := proto.Marshal(&EventSubStatusEvent{
		Detail: detail,
		Status: es.status,
	})
	es.bus.Send(&bus.BusMessage{
		Topic:   BusTopics_TWITCH_EVENTSUB_EVENT.String(),
		Type:    int32(MessageTypeEventSub_TYPE_EVENTSUB_EVENT),
		Message: b,
	})
}
