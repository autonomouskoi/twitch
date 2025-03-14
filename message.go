package twitch

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/nicklaw5/helix/v2"
)

const (
	// scope: channel:manage:redemptions
	EventTypeChannelPointRedeem    = "channel.channel_points_custom_reward_redemption.add"
	EventVersionChannelPointRedeem = "1"

	// scope: bits:read
	EventTypeCheer    = "channel.cheer"
	EventVersionCheer = "1"

	// scope: moderator:read:followers
	EventTypeFollow    = "channel.follow"
	EventVersionFollow = "2"

	// scope: channel:read:hype_train
	EventTypeHypeTrainBegin      = "channel.hype_train.begin"
	EventVersionHypTrainBegin    = "1"
	EventTypeHypeTrainProgress   = "channel.hype_train.progress"
	EventVersionHypTrainProgress = "1"
	EventTypeHypeTrainEnd        = "channel.hype_train.end"
	EventVersionHypTrainEnd      = "1"

	// no permissions
	EventTypeRaid    = "channel.raid"
	EventVersionRaid = "1"

	// scope: channel:read:subscriptions
	EventTypeSubscription    = "channel.subscribe"
	EventVersionSubscription = "1"

	// scope: user:read:chat
	EventVersionChannelChatMessage = "1"

	MessageTypeNotification     = "notification"
	MessageTypeSessionKeepalive = "session_keepalive"
	MessageTypeSessionReconnect = "session_reconnect"
	MessageTypeSessionWelcome   = "session_welcome"
)

var (
	// ErrWrongMsgType indicates the Message doesn't carry the type you want
	ErrWrongMsgType = errors.New("wrong message type")
)

// Message unmarshalls a messages received over the websocket
type Message struct {
	Metadata MessageMetadata `json:"metadata"`
	Payload  json.RawMessage `json:"payload"`
}

// MessageMetadata has fields common to all websocket messages received
type MessageMetadata struct {
	ID        string    `json:"message_id"`
	Type      string    `json:"message_type"`
	Timestamp time.Time `json:"message_timestamp"`
}

// Notification carries the events of interest
type Notification struct {
	Subscription *helix.EventSubSubscription `json:"subscription"`
	Event        json.RawMessage             `json:"event"`
}

// SessionWelcome is the initial message received from the server or indicates
// that you should Reconnect at a different URL
type SessionWelcome struct {
	Session *struct {
		ID                      string    `json:"id"`
		Status                  string    `json:"status"`
		ConnectedAt             time.Time `json:"connected_at"`
		KeepaliveTimeoutSeconds *int      `json:"keepalive_timeout_seconds"`
		ReconnectURL            *string   `json:"reconnect_url"`
	} `json:"session"`
}

// Type returns the type of the message
func (msg *Message) Type() string {
	return msg.Metadata.Type
}

// AsSessionWelcome returns its payload as a SessionWelcome
func (msg *Message) AsSessionWelcome() (*SessionWelcome, error) {
	return msgAsType[SessionWelcome](msg, MessageTypeSessionWelcome)
}

func msgAsType[T any](msg *Message, want string) (*T, error) {
	if msg.Metadata.Type != want {
		return nil, fmt.Errorf("%w: got %s", ErrWrongMsgType, msg.Metadata.Type)
	}
	v := new(T)
	if err := json.Unmarshal(msg.Payload, v); err != nil {
		return nil, fmt.Errorf("unmarshalling: %w", err)
	}
	return v, nil
}
