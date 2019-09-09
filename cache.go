package cache

type ICache interface {
	Put(key string, content string)
	Get(key string) (string, bool)
}
