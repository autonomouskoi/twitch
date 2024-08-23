package twitch

import (
	"sync"
	"time"
)

const lifetime = 60 * 10

type seenIDs struct {
	lock sync.Mutex
	ids  map[string]int64
}

func newSeenIDs() *seenIDs {
	seen := &seenIDs{
		lock: sync.Mutex{},
		ids:  map[string]int64{},
	}

	go func() {
		for {
			time.Sleep(lifetime / 2)
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

func (seen *seenIDs) seenID(id string) bool {
	seen.lock.Lock()
	defer seen.lock.Unlock()
	if _, present := seen.ids[id]; present {
		return true
	}
	seen.ids[id] = time.Now().Unix()
	return false
}
