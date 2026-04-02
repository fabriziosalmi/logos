package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server        ServerConfig          `yaml:"server"`
	Delivery      DeliveryConfig        `yaml:"delivery"`
	Cache         CacheConfig           `yaml:"cache"`
	CORS          CORSConfig            `yaml:"cors"`
	Security      SecurityConfig        `yaml:"security"`
	Accessibility A11yConfig            `yaml:"accessibility"`
	Defaults      DefaultsConfig        `yaml:"defaults"`
	Apps          map[string]*AppConfig `yaml:"apps"`
}

type ServerConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	AdminKey        string        `yaml:"admin_key"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
	IdleTimeout     time.Duration `yaml:"idle_timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type DeliveryConfig struct {
	Compression  CompressionConfig  `yaml:"compression"`
	Optimization OptimizationConfig `yaml:"optimization"`
}

type CompressionConfig struct {
	Gzip bool `yaml:"gzip"`
}

type OptimizationConfig struct {
	Minify    bool `yaml:"minify"`
	Precision int  `yaml:"precision"`
}

type CacheConfig struct {
	Strategy             string `yaml:"strategy"`
	MaxAge               int    `yaml:"max_age"`
	StaleWhileRevalidate int    `yaml:"stale_while_revalidate"`
	ImmutableStatic      bool   `yaml:"immutable_static"`
	ETag                 bool   `yaml:"etag"`

	// Multi-tier cache
	L1Memory L1MemoryConfig `yaml:"l1_memory"`
	L2Disk   L2DiskConfig   `yaml:"l2_disk"`
	L3Redis  L3RedisConfig  `yaml:"l3_redis"`
	CDN      CDNConfig      `yaml:"cdn"`
}

type L1MemoryConfig struct {
	MaxEntries int `yaml:"max_entries"`
}

type L2DiskConfig struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
}

type L3RedisConfig struct {
	Enabled  bool   `yaml:"enabled"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	TTL      int    `yaml:"ttl"` // seconds
}

type CDNConfig struct {
	Enabled       bool   `yaml:"enabled"`
	Provider      string `yaml:"provider"` // auto, cloudflare, fastly, generic
	SurrogateKeys bool   `yaml:"surrogate_keys"`
}

type CORSConfig struct {
	AllowedOrigins   []string `yaml:"allowed_origins"`
	AllowCredentials bool     `yaml:"allow_credentials"`
}

type SecurityConfig struct {
	CSP       string          `yaml:"csp"`
	HSTS      string          `yaml:"hsts"`
	RateLimit RateLimitConfig `yaml:"rate_limit"`
}

type RateLimitConfig struct {
	Enabled           bool    `yaml:"enabled"`
	RequestsPerSecond float64 `yaml:"requests_per_second"`
	Burst             int     `yaml:"burst"`
}

type A11yConfig struct {
	GenerateTitle         bool `yaml:"generate_title"`
	GenerateDesc          bool `yaml:"generate_desc"`
	RespectReducedMotion  bool `yaml:"respect_reduced_motion"`
	RespectSaveData       bool `yaml:"respect_save_data"`
}

type DefaultsConfig struct {
	Animation string `yaml:"animation"`
	Theme     string `yaml:"theme"`
	Format    string `yaml:"format"`
	Scene     string `yaml:"scene"`
	Texture   string `yaml:"texture"`
	Shape     string `yaml:"shape"`
}

type AppConfig struct {
	Color     string `yaml:"color"`
	Color2    string `yaml:"color2"`
	Animation string `yaml:"animation"`
	Theme     string `yaml:"theme"`
	Scene     string `yaml:"scene"`
	Title     string `yaml:"title"`
	Subtitle  string `yaml:"subtitle"`
	Texture   string `yaml:"texture"`
	Shape     string `yaml:"shape"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	cfg := &Config{
		Server: ServerConfig{
			Host:            "0.0.0.0",
			Port:            3000,
			ReadTimeout:     5 * time.Second,
			WriteTimeout:    10 * time.Second,
			IdleTimeout:     120 * time.Second,
			ShutdownTimeout: 10 * time.Second,
		},
		Delivery: DeliveryConfig{
			Compression:  CompressionConfig{Gzip: true},
			Optimization: OptimizationConfig{Minify: true, Precision: 2},
		},
		Cache: CacheConfig{
			Strategy:             "stale-while-revalidate",
			MaxAge:               86400,
			StaleWhileRevalidate: 604800,
			ETag:                 true,
			L1Memory:             L1MemoryConfig{MaxEntries: 10000},
			L2Disk:               L2DiskConfig{Path: "/tmp/logos-cache"},
			L3Redis:              L3RedisConfig{Addr: "localhost:6379", TTL: 86400},
			CDN:                  CDNConfig{Provider: "auto"},
		},
		Accessibility: A11yConfig{
			GenerateTitle:        true,
			GenerateDesc:         true,
			RespectReducedMotion: true,
			RespectSaveData:      true,
		},
		Defaults: DefaultsConfig{
			Animation: "breathe",
			Theme:     "dark",
			Format:    "favicon",
			Scene:     "pure",
			Texture:   "none",
			Shape:     "circle",
		},
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if envPort := os.Getenv("PORT"); envPort != "" {
		var port int
		if _, err := fmt.Sscanf(envPort, "%d", &port); err == nil {
			cfg.Server.Port = port
		}
	}

	if envKey := os.Getenv("ADMIN_KEY"); envKey != "" {
		cfg.Server.AdminKey = envKey
	}

	return cfg, nil
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

// Validate checks the config for invalid values. Fail fast on startup.
func Validate(cfg *Config) error {
	if cfg.Server.Port < 1 || cfg.Server.Port > 65535 {
		return fmt.Errorf("invalid port: %d", cfg.Server.Port)
	}
	if cfg.Server.ReadTimeout <= 0 {
		return fmt.Errorf("read_timeout must be positive")
	}
	if cfg.Server.WriteTimeout <= 0 {
		return fmt.Errorf("write_timeout must be positive")
	}

	validThemes := map[string]bool{"dark": true, "light": true, "solid": true, "glass": true, "auto": true, "system": true}
	if !validThemes[cfg.Defaults.Theme] {
		return fmt.Errorf("invalid default theme: %q", cfg.Defaults.Theme)
	}

	validFormats := map[string]bool{"favicon": true, "avatar": true, "og-card": true, "hero": true}
	if !validFormats[cfg.Defaults.Format] {
		return fmt.Errorf("invalid default format: %q", cfg.Defaults.Format)
	}

	for name, app := range cfg.Apps {
		if app.Color == "" {
			return fmt.Errorf("app %q: color is required", name)
		}
	}

	if cfg.Security.RateLimit.Enabled {
		if cfg.Security.RateLimit.RequestsPerSecond <= 0 {
			return fmt.Errorf("rate_limit.requests_per_second must be positive when enabled")
		}
		if cfg.Security.RateLimit.Burst <= 0 {
			return fmt.Errorf("rate_limit.burst must be positive when enabled")
		}
	}

	return nil
}
