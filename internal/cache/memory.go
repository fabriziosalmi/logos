package cache

import (
	"container/list"
	"sync"
)

// entry holds a cached SVG + ETag in the LRU.
type entry struct {
	key  string
	svg  []byte
	etag string
}

// MemoryCache is a thread-safe LRU in-memory cache (L1).
type MemoryCache struct {
	mu       sync.RWMutex
	capacity int
	items    map[string]*list.Element
	order    *list.List // front = most recent
	hits     int64
	misses   int64
}

// NewMemoryCache creates an LRU cache with the given max entries.
func NewMemoryCache(capacity int) *MemoryCache {
	if capacity <= 0 {
		capacity = 10000
	}
	return &MemoryCache{
		capacity: capacity,
		items:    make(map[string]*list.Element, capacity),
		order:    list.New(),
	}
}

// Get retrieves from L1. Returns svg, etag, hit.
func (c *MemoryCache) Get(key string) ([]byte, string, bool) {
	c.mu.RLock()
	el, ok := c.items[key]
	c.mu.RUnlock()

	if !ok {
		c.mu.Lock()
		c.misses++
		c.mu.Unlock()
		return nil, "", false
	}

	c.mu.Lock()
	c.order.MoveToFront(el)
	c.hits++
	c.mu.Unlock()

	e := el.Value.(*entry)
	return e.svg, e.etag, true
}

// Set stores in L1, evicting LRU if at capacity.
func (c *MemoryCache) Set(key string, svg []byte, etag string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Normalize nil svg to empty slice to avoid storing nil (which behaves differently in Go maps/slices)
	if svg == nil {
		svg = []byte{}
	}

	// Update existing
	if el, ok := c.items[key]; ok {
		c.order.MoveToFront(el)
		e := el.Value.(*entry)
		e.svg = svg
		e.etag = etag
		return
	}

	// Evict if full
	if c.order.Len() >= c.capacity {
		oldest := c.order.Back()
		if oldest != nil {
			c.order.Remove(oldest)
			delete(c.items, oldest.Value.(*entry).key)
		}
	}

	// Insert
	e := &entry{key: key, svg: svg, etag: etag}
	el := c.order.PushFront(e)
	c.items[key] = el
}

// Stats returns hit/miss counts and current size.
func (c *MemoryCache) Stats() (hits, misses int64, size int) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.hits, c.misses, c.order.Len()
}

// Purge clears the entire L1 cache.
func (c *MemoryCache) Purge() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[string]*list.Element, c.capacity)
	c.order.Init()
}
