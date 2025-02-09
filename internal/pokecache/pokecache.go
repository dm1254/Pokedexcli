package pokecache

import (
	"time"
	"sync"

)
type cacheEntry struct{
	CreateAt time.Time
	Val []byte 


}

type Cache struct{
	Cache map[string]cacheEntry
	Interval time.Duration
	mu *sync.Mutex 
	
 
}

func NewCache(interval time.Duration) Cache{
	new_cache := Cache{
		Cache: make(map[string]cacheEntry),
		Interval: interval,
		mu: &sync.Mutex{},	

	}
	go new_cache.reapLoop()
	return new_cache
}

func(c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Cache[key] = cacheEntry{
		CreateAt: time.Now(),
		Val: val,

	}
	
}

func(c *Cache) Get(key string) ([]byte, bool){
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, exists := c.Cache[key]; exists{
		return entry.Val, true	
	}
	return nil,false

}

func(c *Cache) reapLoop(){
	ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()
		
	for range ticker.C{
		current_time := time.Now()
		c.mu.Lock()
		for key,entry := range c.Cache{
			if current_time.Sub(entry.CreateAt) > c.Interval{
				delete(c.Cache,key)
			}

		}
		c.mu.Unlock()
	}
	
}
