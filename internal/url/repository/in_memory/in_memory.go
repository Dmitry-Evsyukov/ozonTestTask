package urlRepository

import (
	"context"
	"main/pkg/err"
	"sync"
	"time"
)

type internal struct {
	timeCreation time.Time
	url          string
}

type InMemory struct {
	mutex        sync.RWMutex
	stopCleaning chan struct{}
	shortToFull  map[string]internal
	timeToLive   time.Duration
}

func NewInMemory(ttl time.Duration) *InMemory {
	result := &InMemory{shortToFull: make(map[string]internal), timeToLive: ttl, stopCleaning: make(chan struct{})}
	go result.cleanUp()
	return result
}

func (im *InMemory) SaveShortToFullUrl(ctx context.Context, short, full string) error {
	im.mutex.Lock()
	defer im.mutex.Unlock()
	im.shortToFull[short] = internal{
		timeCreation: time.Now(),
		url:          full,
	}
	return nil
}

func (im *InMemory) GetFullUrl(ctx context.Context, short string) (string, error) {
	im.mutex.RLock()
	defer im.mutex.RUnlock()
	full, ok := im.shortToFull[short]
	if !ok {
		return "", httpError.NoOriginalUrl
	}

	return full.url, nil
}

func (im *InMemory) cleanUp() {
	ticker := time.NewTicker(im.timeToLive)
	for {
		select {
		case <-ticker.C:
			im.mutex.Lock()
			for short, data := range im.shortToFull {
				if time.Since(data.timeCreation) > im.timeToLive {
					delete(im.shortToFull, short)
				}
			}
			im.mutex.Unlock()
		case <-im.stopCleaning:
			ticker.Stop()
			return
		}
	}
}

func (im *InMemory) stopCleanUp() {
	im.stopCleaning <- struct{}{}
}
