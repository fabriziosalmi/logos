package cache

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

// RedisCache is a minimal Redis client for L3 caching.
// Uses raw RESP protocol — zero external dependencies.
type RedisCache struct {
	addr     string
	password string
	db       int
	ttl      time.Duration
	hits     int64
	misses   int64
}

// NewRedisCache creates a Redis L3 cache. Tests connectivity on creation.
func NewRedisCache(addr, password string, db int, ttl time.Duration) (*RedisCache, error) {
	rc := &RedisCache{addr: addr, password: password, db: db, ttl: ttl}

	// Test connection
	conn, err := rc.dial()
	if err != nil {
		return nil, fmt.Errorf("redis connect: %w", err)
	}
	conn.Close()

	return rc, nil
}

func (rc *RedisCache) dial() (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", rc.addr, 2*time.Second)
	if err != nil {
		return nil, err
	}
	conn.SetDeadline(time.Now().Add(3 * time.Second))

	if rc.password != "" {
		if err := rc.cmd(conn, "AUTH", rc.password); err != nil {
			conn.Close()
			return nil, err
		}
	}
	if rc.db != 0 {
		if err := rc.cmd(conn, "SELECT", strconv.Itoa(rc.db)); err != nil {
			conn.Close()
			return nil, err
		}
	}
	return conn, nil
}

func (rc *RedisCache) cmd(conn net.Conn, args ...string) error {
	// Write RESP array
	fmt.Fprintf(conn, "*%d\r\n", len(args))
	for _, a := range args {
		fmt.Fprintf(conn, "$%d\r\n%s\r\n", len(a), a)
	}

	// Read response (expect +OK or -ERR)
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "-") {
		return fmt.Errorf("redis: %s", line[1:])
	}
	return nil
}

// Get retrieves a value from Redis.
func (rc *RedisCache) Get(key string) ([]byte, string, bool) {
	conn, err := rc.dial()
	if err != nil {
		atomic.AddInt64(&rc.misses, 1)
		return nil, "", false
	}
	defer conn.Close()

	// GET logos:{key}
	rkey := "logos:" + key
	fmt.Fprintf(conn, "*2\r\n$3\r\nGET\r\n$%d\r\n%s\r\n", len(rkey), rkey)

	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		atomic.AddInt64(&rc.misses, 1)
		return nil, "", false
	}
	line = strings.TrimSpace(line)

	if line == "$-1" {
		atomic.AddInt64(&rc.misses, 1)
		return nil, "", false
	}

	if !strings.HasPrefix(line, "$") {
		atomic.AddInt64(&rc.misses, 1)
		return nil, "", false
	}

	size, _ := strconv.Atoi(line[1:])
	if size <= 0 || size > 10*1024*1024 { // 10MB max to prevent memory exhaustion
		atomic.AddInt64(&rc.misses, 1)
		return nil, "", false
	}
	data := make([]byte, size+2) // +2 for \r\n
	n := 0
	for n < len(data) {
		r, err := reader.Read(data[n:])
		if err != nil {
			atomic.AddInt64(&rc.misses, 1)
			return nil, "", false
		}
		n += r
	}

	// Format: etag\n{svg bytes}
	payload := string(data[:size])
	parts := strings.SplitN(payload, "\n", 2)
	if len(parts) != 2 {
		atomic.AddInt64(&rc.misses, 1)
		return nil, "", false
	}

	atomic.AddInt64(&rc.hits, 1)
	return []byte(parts[1]), parts[0], true
}

// Set stores a value in Redis with TTL.
func (rc *RedisCache) Set(key string, svg []byte, etag string) {
	conn, err := rc.dial()
	if err != nil {
		return
	}
	defer conn.Close()

	rkey := "logos:" + key
	value := etag + "\n" + string(svg)
	ttlSec := strconv.Itoa(int(rc.ttl.Seconds()))

	// SETEX logos:{key} {ttl} {value}
	fmt.Fprintf(conn, "*4\r\n$5\r\nSETEX\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n",
		len(rkey), rkey,
		len(ttlSec), ttlSec,
		len(value), value)

	// Read OK
	reader := bufio.NewReader(conn)
	reader.ReadString('\n')
}

// Stats returns hit/miss counts.
func (rc *RedisCache) Stats() (hits, misses int64) {
	return atomic.LoadInt64(&rc.hits), atomic.LoadInt64(&rc.misses)
}
