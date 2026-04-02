package cache

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/fabriziosalmi/logos/internal/config"
)

// Tiered is the multi-layer cache orchestrator.
// L1 (memory) → L2 (disk) → L3 (redis) → miss → render
type Tiered struct {
	L1  *MemoryCache
	L2  *DiskCache  // nil if disabled
	L3  *RedisCache // nil if disabled
	cdn config.CDNConfig

	writeCh chan cacheWriteOp
}

type cacheWriteOp struct {
	tier int
	key  string
	svg  []byte
	etag string
}

// NewTiered builds the cache stack from config.
func NewTiered(cfg config.CacheConfig) *Tiered {
	t := &Tiered{
		writeCh: make(chan cacheWriteOp, 10000), // Buffer to absorb spikes
	}

	// Start async write workers
	for i := 0; i < 5; i++ {
		go func() {
			for op := range t.writeCh {
				if op.tier == 2 && t.L2 != nil {
					t.L2.Set(op.key, op.svg, op.etag)
				} else if op.tier == 3 && t.L3 != nil {
					t.L3.Set(op.key, op.svg, op.etag)
				}
			}
		}()
	}

	// L1: always on
	t.L1 = NewMemoryCache(cfg.L1Memory.MaxEntries)
	slog.Info("cache L1 (memory) enabled", "max_entries", cfg.L1Memory.MaxEntries)

	// L2: optional disk
	if cfg.L2Disk.Enabled {
		dc, err := NewDiskCache(cfg.L2Disk.Path)
		if err != nil {
			slog.Warn("cache L2 (disk) failed to init", "error", err)
		} else {
			t.L2 = dc
			slog.Info("cache L2 (disk) enabled", "path", cfg.L2Disk.Path)
		}
	}

	// L3: optional redis
	if cfg.L3Redis.Enabled {
		ttl := time.Duration(cfg.L3Redis.TTL) * time.Second
		rc, err := NewRedisCache(cfg.L3Redis.Addr, cfg.L3Redis.Password, cfg.L3Redis.DB, ttl)
		if err != nil {
			slog.Warn("cache L3 (redis) failed to connect", "error", err, "addr", cfg.L3Redis.Addr)
		} else {
			t.L3 = rc
			slog.Info("cache L3 (redis) enabled", "addr", cfg.L3Redis.Addr)
		}
	}

	t.cdn = cfg.CDN

	return t
}

func (t *Tiered) asyncWrite(tier int, key string, svg []byte, etag string) {
	select {
	case t.writeCh <- cacheWriteOp{tier: tier, key: key, svg: svg, etag: etag}:
	default:
		// Queue full, drop write to avoid blocking
	}
}

// Get looks up the cache key through all tiers.
// Returns svg, etag, tier ("L1"/"L2"/"L3"/"miss").
func (t *Tiered) Get(key string) ([]byte, string, string) {
	// L1
	if svg, etag, ok := t.L1.Get(key); ok {
		return svg, etag, "L1"
	}

	// L2
	if t.L2 != nil {
		if svg, etag, ok := t.L2.Get(key); ok {
			// Promote to L1
			t.L1.Set(key, svg, etag)
			return svg, etag, "L2"
		}
	}

	// L3
	if t.L3 != nil {
		if svg, etag, ok := t.L3.Get(key); ok {
			// Promote to L1 + L2
			t.L1.Set(key, svg, etag)
			if t.L2 != nil {
				t.asyncWrite(2, key, svg, etag)
			}
			return svg, etag, "L3"
		}
	}

	return nil, "", "miss"
}

// Set stores the result in all active tiers.
func (t *Tiered) Set(key string, svg []byte, etag string) {
	t.L1.Set(key, svg, etag)

	if t.L2 != nil {
		t.asyncWrite(2, key, svg, etag)
	}
	if t.L3 != nil {
		t.asyncWrite(3, key, svg, etag)
	}
}

// SetCDNHeaders writes CDN-specific cache headers based on config.
func (t *Tiered) SetCDNHeaders(w http.ResponseWriter, tier string, appSlug string) {
	// X-Cache: show which tier served the response
	w.Header().Set("X-Cache", tier)

	if !t.cdn.Enabled {
		return
	}

	// CDN-Cache-Control (Cloudflare)
	if t.cdn.Provider == "cloudflare" || t.cdn.Provider == "auto" {
		w.Header().Set("CDN-Cache-Control", "max-age=86400")
	}

	// Surrogate-Control (Fastly/Varnish)
	if t.cdn.Provider == "fastly" || t.cdn.Provider == "auto" {
		w.Header().Set("Surrogate-Control", "max-age=86400")
	}

	// Surrogate-Key / Cache-Tag for targeted purge
	if t.cdn.SurrogateKeys {
		tags := "logos-svg"
		if appSlug != "" {
			tags += " app-" + appSlug
		}
		w.Header().Set("Surrogate-Key", tags)   // Fastly
		w.Header().Set("Cache-Tag", tags)        // Cloudflare
	}
}

// Stats returns a JSON-serializable stats summary.
func (t *Tiered) Stats() map[string]any {
	l1Hits, l1Misses, l1Size := t.L1.Stats()
	stats := map[string]any{
		"l1_hits":   l1Hits,
		"l1_misses": l1Misses,
		"l1_size":   l1Size,
	}

	if t.L2 != nil {
		l2Hits, l2Misses := t.L2.Stats()
		stats["l2_hits"] = l2Hits
		stats["l2_misses"] = l2Misses
	}
	if t.L3 != nil {
		l3Hits, l3Misses := t.L3.Stats()
		stats["l3_hits"] = l3Hits
		stats["l3_misses"] = l3Misses
	}

	return stats
}

// StatsJSON returns stats as JSON bytes.
func (t *Tiered) StatsJSON() []byte {
	data, _ := json.Marshal(t.Stats())
	return data
}

// Purge clears all tiers.
func (t *Tiered) Purge() {
	t.L1.Purge()
	if t.L2 != nil {
		t.L2.Purge()
	}
	slog.Info("cache purged (all tiers)")
}
