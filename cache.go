package cache

import (
	"time"
)

type ICache interface {
	Put(key string, content string, duration time.Duration)
	Get(key string) (string, bool)
	Remove(key string)
}
