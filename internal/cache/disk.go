package cache

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync/atomic"
)

// diskEntry stores SVG + ETag as JSON on disk.
type diskEntry struct {
	SVG  []byte `json:"s"`
	ETag string `json:"e"`
}

// DiskCache is a content-addressable filesystem cache (L2).
// Files are stored as {basePath}/{key[:2]}/{key}.json
type DiskCache struct {
	basePath string
	hits     int64
	misses   int64
}

// NewDiskCache creates a disk-backed cache. Creates the base dir if needed.
func NewDiskCache(basePath string) (*DiskCache, error) {
	if err := os.MkdirAll(basePath, 0750); err != nil {
		return nil, err
	}
	return &DiskCache{basePath: basePath}, nil
}

func (c *DiskCache) path(key string) string {
	// Shard into 256 subdirs to avoid filesystem limits
	prefix := key[:2]
	return filepath.Join(c.basePath, prefix, key+".json")
}

// Get retrieves from disk.
func (c *DiskCache) Get(key string) ([]byte, string, bool) {
	data, err := os.ReadFile(c.path(key))
	if err != nil {
		atomic.AddInt64(&c.misses, 1)
		return nil, "", false
	}

	var e diskEntry
	if err := json.Unmarshal(data, &e); err != nil {
		atomic.AddInt64(&c.misses, 1)
		return nil, "", false
	}

	atomic.AddInt64(&c.hits, 1)
	return e.SVG, e.ETag, true
}

// Set writes to disk. Fire-and-forget — errors are silently ignored.
func (c *DiskCache) Set(key string, svg []byte, etag string) {
	p := c.path(key)
	os.MkdirAll(filepath.Dir(p), 0750)

	data, _ := json.Marshal(diskEntry{SVG: svg, ETag: etag})
	os.WriteFile(p, data, 0640)
}

// Stats returns hit/miss counts.
func (c *DiskCache) Stats() (hits, misses int64) {
	return atomic.LoadInt64(&c.hits), atomic.LoadInt64(&c.misses)
}

// Purge removes all cached files.
func (c *DiskCache) Purge() {
	os.RemoveAll(c.basePath)
	os.MkdirAll(c.basePath, 0750)
}
