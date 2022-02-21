package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data map[string]string
	mux  sync.RWMutex
}

var (
	once  sync.Once
	cache *Cache
)

func New() *Cache {
	once.Do(func() {
		cache = &Cache{}
	})
	return cache
}

func (cache *Cache) get() (map[string]string, error) {
	cache.mux.RLock()
	defer cache.mux.RUnlock()
	return cache.data, nil
}

func (cache *Cache) sync() error {
	timer := time.NewTimer(60 * time.Second)
	for {
		select {
		case <-timer.C:
			if _, err := cache.get(); err != nil {
				fmt.Println("auth sync: %v", err)
			}
		}
		timer.Reset(60 * time.Second)
	}
}
