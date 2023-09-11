package cache

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

var (
	cacheClient *memcache.Client
)

func InitCache() {
	cacheClient = memcache.New("localhost:11211")
}

func Get(key string) []byte {
	item, err := cacheClient.Get(key)
	if err != nil {
		fmt.Println("Error getting item from cache", err)
		return nil
	}
	return item.Value
}

func Set(key string, value []byte) {
	if err := cacheClient.Set(&memcache.Item{
		Key:   key,
		Value: value,
	}); err != nil {
		fmt.Println("Error setting item in cache", err)
	}
}
