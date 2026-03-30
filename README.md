# Logos API

High-performance dynamic SVG asset generation service built in Go. Generate customizable, animated logos, favicons, avatars, OG cards, and hero images on-the-fly.

**Zero runtime dependencies** - single static binary with embedded dashboard.

## Quick Start

```bash
# Server
make run

# CLI: render an app from config.yaml
./logos render minerva > minerva-favicon.svg

# CLI: render custom SVG
./logos render --color=neon --animation=vortex --texture=glitch --format=avatar > logo.svg

# Docker
docker compose up -d
```

Server starts at `http://localhost:3000` (configurable via `config.yaml` or `PORT` env var).

## API

### V4 Cinematic (recommended)

```
GET /api/v4/render/{format}/{scene}/{color1}/{animation}.svg
GET /api/v4/render/{format}/{scene}/{color1}/{color2}/{animation}.svg
```

**Formats**: `favicon` (32x32), `avatar` (512x512), `og-card` (1200x630), `hero` (1920x1080)

**Scenes**: `pure`, `spotlight`, `grid`, `split`

**Themes**: `dark`, `light`, `solid`, `glass`, `auto`

#### Query Parameters

| Param | Type | Description |
|---|---|---|
| `theme` | string | dark, light, solid, glass, auto |
| `title` | string | Title text overlay |
| `subtitle` | string | Subtitle text overlay |
| `texture` | string | grain, glass, noise, glitch, shadow, neon |
| `stroke` | float | Stroke width override (0-10) |
| `padding` | float | Extra viewBox padding (0-20) |
| `alpha` | float | Global opacity (0.0-1.0) |
| `badge` | int | Notification badge number (1-999) |
| `variant` | string | `outline` for transparent fill |
| `hover` | bool | Inject `:hover` brightness CSS |
| `decorative` | bool | Set `aria-hidden="true"` (skip screen readers) |

#### Examples

```
/api/v4/render/favicon/pure/blue/vortex.svg
/api/v4/render/og-card/spotlight/amber/cyan/pulse.svg?theme=dark&title=MyApp
/api/v4/render/hero/grid/neon/laser/spin.svg?texture=glitch&hover=true
/api/v4/render/avatar/pure/rose/breathe.svg?badge=3&stroke=2
/api/v4/render/favicon/pure/green/static.svg?decorative=true&alpha=0.5
```

### App Shortcuts

Register apps in `config.yaml`, access via:

```
GET /app/{name}/{format}.svg
```

```yaml
apps:
  minerva:
    color: amber
    animation: breathe
    theme: dark
    title: Minerva
    texture: grain
```

```
/app/minerva/favicon.svg     -> amber favicon, breathe, grain texture
/app/minerva/og-card.svg     -> amber OG card with title "Minerva"
/app/unknown/favicon.svg     -> gray placeholder (200, not 404)
```

### V3 Legacy

```
GET /api/favicon/{color1}/{animation}.svg
GET /api/favicon/{color1}/{color2}/{animation}.svg
```

### Health Check

```
GET /healthz
```

## CLI

```bash
# Render registered app
logos render minerva > favicon.svg
logos render rainlogs --format=og-card > og.svg

# Render custom
logos render --color=cyber --animation=vortex --texture=neon --badge=5 > logo.svg

# Version
logos version
```

## Configuration

All settings in `config.yaml`:

```yaml
server:
  host: "0.0.0.0"
  port: 3000
  read_timeout: 5s
  write_timeout: 10s

delivery:
  compression:
    gzip: true

cache:
  strategy: stale-while-revalidate
  max_age: 86400
  stale_while_revalidate: 604800
  etag: true

security:
  hsts: "max-age=63072000; includeSubDomains; preload"
  rate_limit:
    enabled: false
    requests_per_second: 50
    burst: 100

accessibility:
  generate_title: true
  generate_desc: true
  respect_reduced_motion: true
  respect_save_data: true

defaults:
  animation: breathe
  theme: dark
  texture: none

apps:
  my-app:
    color: cyan
    animation: zen
    title: My App
    texture: glass
```

## Colors

**Base**: amber, blue, cyan, green, indigo, orange, purple, rose

**Creative**: white, gray, black, gold, platinum, champagne, neon, matrix, cyber, laser, plasma, void, emerald, sapphire, ruby, ocean, sunset, magma, mint, peach, lavender

Custom hex values also supported (e.g., `ff6600`).

## Animations

50 styles: static, zen, breathe, levitate, glimmer, spin, spin-fast, smooth-spin, orbit-chase, compass, gyro, satellite, eclipse, pulse, heartbeat, pulse-ring, strobe, nova, elastic, flip, orbit-tilt, vortex, harmony, sync, sway, radar, radar-sweep, signal, glow, aurora, nebula, corona, ripple-core, morph, morph-blob, morph-crystal, bounce, bounce-drop, trampoline, shake, jitter, earthquake, swing, pendulum, wave-swing, zoom-in, zoom-out, zoom-pulse, slide-in, slide-loop, typewriter-blink.

## Textures

SVG filter effects applied to the logo core:

| Texture | Effect |
|---|---|
| `grain` | Analog film noise (`feTurbulence` fractal) |
| `glass` | Frosted glass refraction (`feGaussianBlur` + `feDisplacementMap`) |
| `noise` | Digital noise overlay |
| `glitch` | Chromatic aberration (RGB channel split) |
| `shadow` | Multi-layer soft drop shadow (organic glow) |
| `neon` | Neon outer glow (triple gaussian blur merge) |

## Accessibility

Every SVG includes by default:
- `role="img"` + `aria-labelledby`
- `<title>` and `<desc>` with semantic context
- `@media(prefers-reduced-motion:reduce)` — kills animations
- `@media(forced-colors:active)` — Windows High Contrast support
- `vector-effect="non-scaling-stroke"` — constant stroke width at any zoom
- `Save-Data` header detection — strips textures on slow networks
- `?decorative=true` — sets `aria-hidden="true"` for ornamental use

## Architecture

```
cmd/logos/              Entry + CLI + graceful shutdown
internal/
  config/               YAML config + validation (fail-fast)
  svg/
    engine.go           Master renderer (sync.Pool, ETag, badge, params)
    animation.go        50 animations
    palette.go          23 colors + hex resolver
    scene.go            Scene composition
    theme.go            Theme engine
    typography.go       Text overlay
    texture.go          6 SVG filters (grain/glass/noise/glitch/shadow/neon)
    format.go           Output dimensions
  handler/
    api.go              V3/V4/App handlers + fallback placeholder
    dashboard.go        Embedded dashboard
  middleware/
    middleware.go       CORS, SWR cache, ETag, gzip, rate limit, security, logging
web/
  dashboard/            Modular ES6 frontend
```

## Docker

Distroless image, non-root, read-only filesystem, no capabilities:

```bash
docker compose up -d
```

## Performance

- 6.2MB static binary, zero runtime deps
- SVG generation: <100µs per request
- `sync.Pool` buffer reuse
- Gzip compression (stdlib, pooled writers)
- ETag + `304 Not Modified`
- Stale-While-Revalidate cache headers
- Pre-warm registered apps on startup
- Embedded assets (no disk I/O)

## License

[MIT](LICENSE)
