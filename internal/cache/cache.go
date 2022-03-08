package cache

import (
	"sync"
	"time"
)

type item struct {
	value   []string
	expires time.Time
}

func (i item) IsExpired(now time.Time) bool {
	return i.expires.UnixNano() < now.UnixNano()
}

type Cache struct {
	items         map[string]*item
	commonExpires time.Duration
	mu            sync.Mutex
}

// Set keyがすでに存在していた場合は上書きされます
func (c *Cache) Set(key string, v []string) {
	c.mu.Lock()

	// 保持しているcacheの中身を変更できないようにcopyを返す
	copied := make([]string, len(v))
	copy(copied, v)

	c.items[key] = &item{
		value:   copied,
		expires: time.Now().Add(c.commonExpires),
	}

	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.items[key]
	if !ok {
		return nil, false
	}

	if v.IsExpired(time.Now()) {
		return nil, false
	}

	// 保持しているcacheの中身を変更できないようにcopyを返す
	result := make([]string, len(v.value))
	copy(result, v.value)

	return result, true
}

const expireCheckDuration = 10 * time.Second

func New(expires time.Duration) *Cache {
	c := &Cache{
		items:         make(map[string]*item),
		commonExpires: expires,
	}

	go autoDeleteExpired(c, expireCheckDuration)

	return c
}

func autoDeleteExpired(c *Cache, d time.Duration) {
	t := time.NewTicker(d)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			c.mu.Lock()
			for k, v := range c.items {
				if v.IsExpired(time.Now()) {
					delete(c.items, k)
				}
			}
			c.mu.Unlock()
		}
	}
}
