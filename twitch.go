package twitch

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/nicklaw5/helix/v2"
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
	// EnvLocalContentPath specifies an env var with a path to serve web content
	// from instead of the embedded zip file, for development
	EnvLocalContentPath = "AK_CONTENT_TWITCH"

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
		Title:       "Twitch",
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

// Twitch provides an interface to the Twitch API
type Twitch struct {
	http.Handler
	modutil.ModuleBase
	bus         *bus.Bus
	lock        sync.Mutex
	kv          *kv.KVPrefix
	cfg         *Config
	clients     map[string]*client
	storagePath string

	cacheUsers ttlcache.Cache[string, *User]

	eventSub  *eventSub
	shoutouts *shoutouts
}

//go:embed web.zip
var webZip []byte

// Start the module
func (t *Twitch) Start(ctx context.Context, deps *modutil.ModuleDeps) error {
	t.Log = deps.Log
	t.bus = deps.Bus
	t.kv = &deps.KV
	t.storagePath = deps.StoragePath
	t.clients = map[string]*client{}
	t.eventSub = &eventSub{
		bus:       deps.Bus,
		parentCtx: ctx,
		seen:      newSeenIDs(),
	}
	t.eventSub.Log = t.Log

	t.cacheUsers = ttlcache.New[string, *User](ctx, time.Minute*15, time.Minute)

	fs, err := webutil.ZipOrEnvPath(EnvLocalContentPath, webZip)
	if err != nil {
		return fmt.Errorf("get web FS %w", err)
	}
	t.Handler = http.FileServer(fs)

	t.cfg = &Config{}
	if err := t.kv.GetProto(cfgKVKey, t.cfg); err != nil && !errors.Is(err, akcore.ErrNotFound) {
		return fmt.Errorf("retrieving config: %w", err)
	}
	defer t.writeCfg()

	for name, profile := range t.cfg.Profiles {
		if err := t.addProfile(profile); err != nil {
			t.Log.Error("adding profile", "profile", profile.Name, "error", err)
			continue
		}
		t.Log.Debug("loaded profile", "name", name)
	}

	t.shoutouts = newShoutouts(ctx, t.Log, t.clients)

	t.Go(func() error { return t.handleRequest(ctx) })
	t.Go(func() error { return t.handleCommand(ctx) })
	t.Go(func() error { return t.handleChat(ctx) })
	t.Go(func() error { return t.handleEventSub(ctx) })

	return t.Wait()
}

func (t *Twitch) writeCfg() {
	if err := t.kv.SetProto(cfgKVKey, t.cfg); err != nil {
		t.Log.Error("writing config", "error", err.Error())
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

func defaultValue[T comparable](a, b T) T {
	var v T
	if a == v {
		return b
	}
	return a
}

//go:embed icon.svg
var icon []byte

func (*Twitch) Icon() ([]byte, string, error) {
	return icon, "image/svg+xml", nil
}
