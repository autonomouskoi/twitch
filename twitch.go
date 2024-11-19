package twitch

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/nicklaw5/helix/v2"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"

	"github.com/autonomouskoi/akcore"
	"github.com/autonomouskoi/akcore/bus"
	"github.com/autonomouskoi/akcore/modules"
	"github.com/autonomouskoi/akcore/modules/modutil"
	"github.com/autonomouskoi/akcore/storage/kv"
	"github.com/autonomouskoi/akcore/web/webutil"
	"github.com/autonomouskoi/datastruct/ttlcache"
)

const (
	EnvLocalContentPath = "TWITCH_WEB_CONTENT"

	clientID = "zqciq1diwsv0xqizn7m2gxbke6ez2v"
)

var (
	cfgKVKey = []byte("config")

	errInvalidClient = &bus.Error{
		Code:        int32(bus.CommonErrorCode_NOT_FOUND),
		Detail:      proto.String("invalid profile"),
		UserMessage: proto.String("no such Twitch profile"),
	}
)

func init() {
	manifest := &modules.Manifest{
		Id:          "27cff9c0fb8385a6",
		Name:        "twitch",
		Description: "Integration with Twitch APIs",
		WebPaths: []*modules.ManifestWebPath{
			{
				Path:        "https://autonomouskoi.org/module-twitch.html",
				Type:        modules.ManifestWebPathType_MANIFEST_WEB_PATH_TYPE_HELP,
				Description: "Help!",
			},
			{
				Path:        "/m/twitch/embed_ctrl.js",
				Type:        modules.ManifestWebPathType_MANIFEST_WEB_PATH_TYPE_EMBED_CONTROL,
				Description: "Controls for Twitch",
			},
			{
				Path:        "/m/twitch/index.html",
				Type:        modules.ManifestWebPathType_MANIFEST_WEB_PATH_TYPE_CONTROL_PAGE,
				Description: "Controls for Twitch",
			},
		},
	}
	modules.Register(manifest, &Twitch{})
}

type Twitch struct {
	http.Handler
	bus     *bus.Bus
	lock    sync.Mutex
	log     *slog.Logger
	kv      *kv.KVPrefix
	cfg     *Config
	clients map[string]*client

	cacheUsers ttlcache.Cache[string, *User]

	chat     *chat
	eventSub *eventSub
}

//go:embed web.zip
var webZip []byte

func (t *Twitch) Start(ctx context.Context, deps *modutil.ModuleDeps) error {
	t.log = deps.Log
	t.bus = deps.Bus
	t.kv = &deps.KV
	t.clients = map[string]*client{}
	t.chat = &chat{
		bus:       deps.Bus,
		log:       deps.Log.With("module", "twitchchat"),
		parentCtx: ctx,
	}
	t.eventSub = &eventSub{
		bus:       deps.Bus,
		parentCtx: ctx,
		seen:      newSeenIDs(),
	}
	t.eventSub.Log = deps.Log.With("module", "twitcheventsub")

	t.cacheUsers = ttlcache.New[string, *User](ctx, time.Minute*15, time.Minute)

	fs, err := webutil.ZipOrEnvPath(EnvLocalContentPath, webZip)
	if err != nil {
		return fmt.Errorf("get web FS %w", err)
	}
	t.Handler = http.StripPrefix("/m/twitch", http.FileServer(fs))

	t.cfg = &Config{}
	if err := t.kv.GetProto(cfgKVKey, t.cfg); err != nil && !errors.Is(err, akcore.ErrNotFound) {
		return fmt.Errorf("retrieving config: %w", err)
	}
	defer t.writeCfg()

	for name, profile := range t.cfg.Profiles {
		if err := t.addProfile(profile); err != nil {
			return fmt.Errorf("adding profile %s: %w", name, err)
		}
		t.log.Debug("loaded profile", "name", name)
	}

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error { return t.handleRequest(ctx) })
	eg.Go(func() error { return t.handleCommand(ctx) })
	eg.Go(func() error { return t.handleChat(ctx) })
	eg.Go(func() error { return t.handleEventSub(ctx) })

	return eg.Wait()
}

func (t *Twitch) writeCfg() {
	if err := t.kv.SetProto(cfgKVKey, t.cfg); err != nil {
		t.log.Error("writing config", "error", err.Error())
	}
}

func extractError(err error, rc helix.ResponseCommon) error {
	if err != nil {
		return err
	}
	if rc.ErrorMessage == "" {
		return nil
	}
	return errors.New(rc.ErrorMessage)
}
