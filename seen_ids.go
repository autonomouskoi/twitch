package twitch

import (
	"sync"
	"time"
)

const seenIDLifetime = time.Minute * 10

// seenIDs holds eventsub message IDs for deduplication. Each message ID is
// dropped in the background after 10 minutes
type seenIDs struct {
	lock sync.Mutex
	ids  map[string]int64
}

func newSeenIDs() *seenIDs {
	seen := &seenIDs{
		lock: sync.Mutex{},
		ids:  map[string]int64{},
	}

	go func() { // in the background, expire IDs
		for {
			time.Sleep(seenIDLifetime / 2)
			now := time.Now().Unix()
			seen.lock.Lock()
			for k, expires := range seen.ids {
				if expires < now {
					delete(seen.ids, k)
				}
			}
			seen.lock.Unlock()
		}
	}()

	return seen
}

// seenID returns whether or not an ID has been seen. It will return true even
// if the ID is "expired", but that still indicates a duplicate message. If the
// ID isn't present it's added to the map
func (seen *seenIDs) seenID(id string) bool {
	seen.lock.Lock()
	defer seen.lock.Unlock()
	if _, present := seen.ids[id]; present {
		return true
	}
	seen.ids[id] = time.Now().Add(seenIDLifetime).Unix()
	return false
}
