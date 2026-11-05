
package pokecache

import (
	"sync"
	"encoding/json"
	"time"
	"fmt"
)


type Maps struct {
	Count    int    
	Next     string 
	Previous string 
	Results  []struct {
		Name string 
		URL  string 	
	} 
}




type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}


func (c *Cache) Find(location string) (string, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for _, value := range c.cache {
			maps := Maps{}
			err := json.Unmarshal(value.val, &maps)
			if err != nil {
				fmt.Println(err)
			}
			for _, item := range maps.Results{
				if item.Name == location {
					return item.URL, true
				}
			}
	}
	return "", false
}
