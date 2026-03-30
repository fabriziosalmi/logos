package middleware

import (
	"compress/gzip"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/fabriziosalmi/logos/internal/config"
)

// CORS adds cross-origin headers.
func CORS(cfg config.CORSConfig) func(http.Handler) http.Handler {
	origins := strings.Join(cfg.AllowedOrigins, ", ")
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", origins)
			w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")
			if cfg.AllowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// Security adds hardening headers.
func Security(cfg config.SecurityConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("X-Frame-Options", "SAMEORIGIN")
			w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
			w.Header().Set("Permissions-Policy", "camera=(), microphone=(), geolocation=()")
			if cfg.CSP != "" {
				w.Header().Set("Content-Security-Policy", cfg.CSP)
			}
			if cfg.HSTS != "" {
				w.Header().Set("Strict-Transport-Security", cfg.HSTS)
			}
			next.ServeHTTP(w, r)
		})
	}
}

// CacheSWR sets Cache-Control with Stale-While-Revalidate strategy.
func CacheSWR(cfg config.CacheConfig) func(http.Handler) http.Handler {
	val := fmt.Sprintf("public, max-age=%d", cfg.MaxAge)
	if cfg.StaleWhileRevalidate > 0 {
		val += fmt.Sprintf(", stale-while-revalidate=%d", cfg.StaleWhileRevalidate)
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", val)
			next.ServeHTTP(w, r)
		})
	}
}

// CacheImmutable sets aggressive caching for truly immutable static assets.
func CacheImmutable(maxAge int) func(http.Handler) http.Handler {
	val := fmt.Sprintf("public, max-age=%d, immutable", maxAge)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", val)
			next.ServeHTTP(w, r)
		})
	}
}

// ETag handles If-None-Match validation. The handler must set the ETag header before writing.
func ETag(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		// ETag comparison is done in the handler itself for efficiency
		// since we need the hash before writing the body
	})
}

// Gzip compresses responses that accept gzip encoding.
func Gzip(next http.Handler) http.Handler {
	pool := sync.Pool{
		New: func() any {
			gz, _ := gzip.NewWriterLevel(io.Discard, gzip.BestSpeed)
			return gz
		},
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz := pool.Get().(*gzip.Writer)
		defer pool.Put(gz)
		gz.Reset(w)

		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Del("Content-Length")
		w.Header().Set("Vary", "Accept-Encoding")

		gzw := &gzipWriter{ResponseWriter: w, Writer: gz}
		next.ServeHTTP(gzw, r)
		gz.Close()
	})
}

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w *gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// RateLimit implements a simple token bucket rate limiter.
func RateLimit(cfg config.RateLimitConfig) func(http.Handler) http.Handler {
	if !cfg.Enabled {
		return func(next http.Handler) http.Handler { return next }
	}

	type bucket struct {
		tokens    float64
		lastCheck time.Time
	}

	var mu sync.Mutex
	buckets := make(map[string]*bucket)
	rate := cfg.RequestsPerSecond
	burst := float64(cfg.Burst)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			if i := strings.LastIndex(ip, ":"); i != -1 {
				ip = ip[:i]
			}

			mu.Lock()
			b, ok := buckets[ip]
			if !ok {
				b = &bucket{tokens: burst, lastCheck: time.Now()}
				buckets[ip] = b
			}

			now := time.Now()
			elapsed := now.Sub(b.lastCheck).Seconds()
			b.lastCheck = now
			b.tokens += elapsed * rate
			if b.tokens > burst {
				b.tokens = burst
			}

			if b.tokens < 1 {
				mu.Unlock()
				w.Header().Set("Retry-After", "1")
				http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			b.tokens--
			mu.Unlock()

			next.ServeHTTP(w, r)
		})
	}
}

// Logger logs requests with duration.
func Logger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			sw := &statusWriter{ResponseWriter: w, status: 200}
			next.ServeHTTP(sw, r)
			logger.Info("request",
				"method", r.Method,
				"path", r.URL.Path,
				"status", sw.status,
				"duration", time.Since(start).String(),
			)
		})
	}
}

// Recover catches panics and returns 500.
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("panic recovered", "error", err, "path", r.URL.Path)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}
