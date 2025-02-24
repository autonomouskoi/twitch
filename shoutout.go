package twitch

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/autonomouskoi/akcore/bus"
	"github.com/nicklaw5/helix/v2"
)

const (
	shoutoutDelay      = time.Second * (120 * 5)   // 2 minutes per SO, +headroom
	shoutoutExpiration = time.Hour + time.Second*5 // two hours for repeat SO, +headroom
)

func (t *Twitch) handleRequestSendShoutout(reqMsg *bus.BusMessage) *bus.BusMessage {
	reply := &bus.BusMessage{
		Topic: reqMsg.GetTopic(),
		Type:  reqMsg.GetType() + 1,
	}
	ssr := &SendShoutoutRequest{}
	if reply.Error = t.UnmarshalMessage(reqMsg, ssr); reply.Error != nil {
		return reply
	}
	t.shoutouts.request(ssr)
	return reply
}

type shoutouts struct {
	clients map[string]*client
	queue   chan *SendShoutoutRequest
	recent  map[string]time.Time
	log     *slog.Logger
	lock    sync.Mutex
}

func newShoutouts(ctx context.Context, log *slog.Logger, clients map[string]*client) *shoutouts {
	so := &shoutouts{
		log:     log,
		clients: clients,
		queue:   make(chan *SendShoutoutRequest, 64),
		recent:  map[string]time.Time{},
	}
	go so.handleQueue(ctx)
	return so
}

func (so *shoutouts) handleQueue(ctx context.Context) {
	queue := so.queue
	for req := range queue {
		so.send(req)
		select {
		case <-ctx.Done():
			so.lock.Lock()
			so.queue = nil
			so.lock.Unlock()
			close(queue)
			bus.Drain(queue)
			return
		case <-time.After(shoutoutDelay):
		}
	}
}

func (so *shoutouts) send(req *SendShoutoutRequest) {
	so.lock.Lock()
	defer so.lock.Unlock()
	client := so.clients[req.FromProfile]
	if client == nil {
		so.log.Error("no matching client", "profile", req.FromProfile)
		return
	}
	now := time.Now()
	for k, expires := range so.recent {
		if expires.Before(now) {
			delete(so.recent, k)
		}
	}
	if _, present := so.recent[req.ToBroadcasterId]; present {
		return
	}
	so.recent[req.ToBroadcasterId] = now.Add(shoutoutExpiration)
	resp, err := client.SendShoutout(&helix.SendShoutoutParams{
		FromBroadcasterID: client.token.UserId,
		ToBroadcasterID:   req.ToBroadcasterId,
		ModeratorID:       req.ModeratorId,
	})
	if err := extractError(err, resp.ResponseCommon); err != nil {
		so.log.Error("sending shoutout request", "error", err.Error())
	}
}

func (so *shoutouts) request(req *SendShoutoutRequest) {
	so.lock.Lock()
	defer so.lock.Unlock()
	if so.queue == nil {
		return
	}
	so.queue <- req
}
