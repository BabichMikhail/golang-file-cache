package cache

import (
	"time"
)

type DummyCache struct {
	ICache
}

func NewDummyCache() ICache {
	return new(DummyCache)
}

func (c *DummyCache) Put(key string, content string, duration time.Duration) {
}

func (c *DummyCache) Get(key string) (string, bool) {
	return "", false
}

func (c *DummyCache) Remove(key string) {
}
