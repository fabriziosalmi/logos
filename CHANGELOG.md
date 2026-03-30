# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [5.2.0] - 2026-03-30

### Added

- **CLI render mode**: `logos render minerva > favicon.svg` or `logos render --color=neon --animation=vortex`
- **CLI version**: `logos version`
- **Query params**: `?stroke=`, `?padding=`, `?alpha=`, `?badge=N`, `?variant=outline`, `?hover=true`, `?decorative=true`
- **Badge notification**: red dot with count (1-999) in corner, scales per format
- **Variant=outline**: forces transparent fill with visible strokes
- **Hover state injection**: `:hover{filter:brightness(1.2)}` CSS in SVG
- **`aria-hidden="true"`**: `?decorative=true` marks SVG as decorative for screen readers
- **`forced-colors:active`**: Windows High Contrast media query in every SVG
- **`vector-effect="non-scaling-stroke"`**: constant stroke width at any scale
- **`preserveAspectRatio="xMidYMid meet"`**: explicit on every SVG root
- **Textures**: `glitch` (chromatic aberration), `shadow` (soft multi-layer drop shadow), `neon` (outer glow)
- **Fallback placeholder**: unknown apps return 200 with gray placeholder SVG + `X-Logos-Fallback: true` header
- **Pre-warm cache**: all registered app favicons rendered into memory on startup
- **Config validation**: fail-fast on invalid port, theme, format, missing app color, bad rate limit
- **`Vary: Accept-Encoding, Save-Data`**: correct cache key differentiation

### Changed

- SVG root: removed fixed `width`/`height`, uses `preserveAspectRatio` + `viewBox` only
- Fallback for unknown apps: 200 with placeholder instead of 404

## [5.1.0] - 2026-03-30

### Added

- **Accessibility (a11y)**: every SVG now has `role="img"`, `<title>`, `<desc>` with semantic context
- **prefers-reduced-motion**: `@media` query injected into all animated SVGs, kills animations for motion-sensitive users
- **Save-Data awareness**: strips textures and filters when client sends `Save-Data: on` header
- **Stale-While-Revalidate cache**: latency-zero cache strategy (serve stale, regenerate in background)
- **ETag generation**: SHA-256 content hash, returns 304 Not Modified on match
- **Gzip compression**: stdlib gzip middleware with pooled writers
- **Rate limiting**: token bucket algorithm (configurable, disabled by default)
- **Texture filters**: `grain` (analog film noise), `glass` (frosted glass refraction), `noise` (digital noise)
- **Texture query param**: `?texture=grain|glass|noise` on V4 API
- **HSTS header**: `max-age=63072000; includeSubDomains; preload`
- **Permissions-Policy**: camera, microphone, geolocation blocked
- **CSP tightened**: explicit `script-src`, `connect-src`, `base-uri`, `object-src 'none'`
- **App `texture` field**: per-app texture configuration in `config.yaml`
- **App `subtitle` as role**: used in `<desc>` for semantic a11y context

### Changed

- Cache middleware: simple max-age -> SWR with stale-while-revalidate
- Security middleware: added HSTS, Permissions-Policy
- SVG engine: returns `RenderResult{SVG, ETag}` instead of raw bytes
- Handlers: shared config init via `InitHandlers()`, ETag/304 support

## [5.0.0] - 2026-03-30

### Added

- **Go rewrite**: entire backend rewritten in Go for maximum performance (<100µs per SVG)
- **Modular architecture**: clean separation into `svg/`, `handler/`, `middleware/`, `config/`
- **App registry**: `config.yaml` with app shortcuts (`/app/{name}/{format}.svg`)
- **Modular frontend**: ES6 modules (theme, grid, preview, url-builder, toast)
- **Health endpoint**: `/healthz` for container orchestration
- **Graceful shutdown**: clean SIGTERM handling with configurable timeout
- **Security middleware**: CSP, X-Content-Type-Options, X-Frame-Options, Referrer-Policy
- **CORS middleware**: configurable allowed origins
- **Cache middleware**: configurable max-age per route group
- **Structured logging**: JSON slog with request duration tracking
- **Panic recovery**: middleware catches panics, returns 500
- **Docker distroless**: multi-stage build, nonroot user, read-only FS, no capabilities
- **docker-compose**: resource limits, tmpfs, health checks, security hardening
- **Embedded assets**: dashboard + static files baked into binary via `embed`
- **sync.Pool buffers**: zero-allocation SVG rendering on hot path
- **Makefile**: build, run, dev, docker, test targets
- **config.yaml**: full server tuning, cache, CORS, security, defaults, app registry

### Changed

- Server language: Node.js -> Go (chi router + net/http)
- Dashboard: monolithic JS -> 6 ES6 modules
- CSS: inline styles -> extracted `style.css`
- Version bump: 4.0.0 -> 5.0.0

## [4.0.0] - 2026-03-30

### Added

- V4 Cinematic API with scene composition, themes, and typography support
- 50 animation styles
- 23 color palettes
- 4 output formats, 4 scene modes, 5 theme variants
- Interactive dashboard with live preview
- V3 legacy API compatibility

## [0.0.1] - 2026-03-30

### Added

- Initial project setup
