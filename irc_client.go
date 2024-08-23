package twitch

import (
	"crypto/tls"
	"fmt"
	"log/slog"
	"runtime"
	"strconv"
	"strings"

	"github.com/autonomouskoi/akcore/bus"
	irc "github.com/fluffle/goirc/client"
	"google.golang.org/protobuf/proto"
)

const (
	ircServerAddr = "irc.chat.twitch.tv"
	ircServerPort = "6697"
)

var (
	ircCapabilities = []string{
		"twitch.tv/membership",
		"twitch.tv/tags",
	}
)

func newIRCClient(cfg *IRCConfig, mBus *bus.Bus, log *slog.Logger, accessToken string) *irc.Conn {
	ircCfg := irc.NewConfig(cfg.Profile)
	ircCfg.EnableCapabilityNegotiation = true
	ircCfg.Capabilites = ircCapabilities
	ircCfg.Recover = func(c *irc.Conn, l *irc.Line) {
		if err := recover(); err != nil {
			_, f, l, _ := runtime.Caller(2)
			log.Error("IRC Panic",
				"file", f,
				"line", l,
				"error", err,
			)
		}
	}
	ircCfg.SSL = true
	ircCfg.SSLConfig = &tls.Config{ServerName: ircServerAddr}
	ircCfg.Server = ircServerAddr + ":" + ircServerPort
	ircCfg.NewNick = func(n string) string { return n + "^" }
	ircCfg.Pass = "oauth:" + accessToken
	ircC := irc.Client(ircCfg)

	ircC.HandleFunc(irc.CONNECTED, func(conn *irc.Conn, _ *irc.Line) {
		log.Info("IRC connected")
		conn.Join("#selfdrivingcarp")
	})
	ircC.HandleFunc(irc.PING, func(conn *irc.Conn, line *irc.Line) {
		conn.Pong(line.Text())
	})
	ircC.HandleFunc(irc.PONG, func(_ *irc.Conn, _ *irc.Line) {
		// do nothing
	})
	ircC.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {
		log.Debug("RECV", "text", line.Text())
		/*
			if line.Nick != "" {
				bot.seen(line.Nick)
			}
		*/
		b, err := proto.Marshal(pb(line))
		if err != nil {
			log.Error("marshaling MessageIn proto", "error", err.Error())
			return
		}
		mBus.Send(
			&bus.BusMessage{
				Topic:   BusTopics_TWITCH_CHAT_RECV.String(),
				Message: b,
			})
	})

	// send these to the event bus
	ircC.HandleFunc(irc.JOIN, func(_ *irc.Conn, line *irc.Line) {
		b, err := proto.Marshal(&ChatEvent{
			Type:   ChatEventType_EVENT_TYPE_JOIN,
			Detail: line.Nick,
		})
		if err != nil {
			log.Error("marshalling join proto", "error", err.Error())
			return
		}
		mBus.Send(
			&bus.BusMessage{
				Topic:   BusTopics_TWITCH_CHAT_EVENT.String(),
				Type:    int32(ChatEventType_EVENT_TYPE_JOIN),
				Message: b,
			})
	})

	for _, msgType := range []string{
		irc.REGISTER,
		//irc.CONNECTED,
		irc.DISCONNECTED,
		irc.ACTION,
		irc.AUTHENTICATE,
		irc.AWAY,
		irc.CAP,
		irc.CTCP,
		irc.CTCPREPLY,
		irc.ERROR,
		irc.INVITE,
		//irc.JOIN,
		irc.KICK,
		irc.MODE,
		irc.NICK,
		irc.NOTICE,
		irc.OPER,
		irc.PART,
		irc.PASS,
		//irc.PING,
		//irc.PONG,
		//irc.PRIVMSG,
		irc.QUIT,
		irc.TOPIC,
		irc.USER,
		irc.VERSION,
		irc.VHOST,
		irc.WHO,
		irc.WHOIS,
	} {
		msgType := msgType
		ircC.HandleFunc(msgType, func(c *irc.Conn, l *irc.Line) {
			log.Debug("unhandled message", "type", msgType, "values", l)
		})
	}

	return ircC
}

func pb(line *irc.Line) *ChatMessageIn {
	mi := &ChatMessageIn{
		Text:   line.Text(),
		Args:   line.Args,
		Nick:   line.Nick,
		Emotes: emotes(line),
		Tags:   line.Tags,
	}
	if isMod(line) {
		t := true
		mi.IsMod = &t
	}

	return mi
}

func isMod(line *irc.Line) bool {
	if line.Tags["user-id"] == line.Tags["room-id"] && line.Tags["room-id"] != "" {
		return true
	}
	if line.Tags["mod"] == "1" {
		return true
	}
	return false
}

func EmoteURL(e *Emote) string {
	return fmt.Sprintf("https://static-cdn.jtvnw.net/emoticons/v2/%s/static/dark/3.0", e.Id)
}

func emotes(line *irc.Line) []*Emote {
	tag := line.Tags["emotes"]
	if tag == "" {
		return nil
	}

	var emotes []*Emote

	for _, emoteStr := range strings.Split(tag, "/") {
		id, ranges, matched := strings.Cut(emoteStr, ":")
		if !matched {
			continue
		}
		var positions []*EmotePosition
		for _, rangeStr := range strings.Split(ranges, ",") {
			startStr, endStr, matched := strings.Cut(rangeStr, "-")
			if !matched {
				continue
			}
			start, err := strconv.Atoi(startStr)
			if err != nil {
				continue
			}
			end, err := strconv.Atoi(endStr)
			if err != nil {
				continue
			}
			positions = append(positions,
				&EmotePosition{
					Start: uint32(start),
					End:   uint32(end),
				},
			)
		}
		if len(positions) == 0 {
			continue
		}
		emotes = append(emotes, &Emote{
			Id:        id,
			Positions: positions,
		})
	}

	return emotes
}
