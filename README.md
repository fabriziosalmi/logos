# Logos API

Dynamic SVG asset generation service. Generate customizable, animated logos, favicons, avatars, OG cards, and hero images on-the-fly via a simple HTTP API.

**Zero dependencies** - pure Node.js.

## Quick Start

```bash
node server.js
```

Server starts at `http://localhost:3000` (override with `PORT` env var).

Open `http://localhost:3000` for the interactive dashboard.

## API

### V4 Cinematic (recommended)

```
GET /api/v4/render/:format/:scene/:color1/:color2?/:animation.svg
```

Query parameters: `theme`, `title`, `subtitle`

**Formats**: `favicon` (32x32), `avatar` (512x512), `og-card` (1200x630), `hero` (1920x1080)

**Scenes**: `pure`, `spotlight`, `grid`, `split`

**Themes**: `dark`, `light`, `solid`, `glass`, `auto`

#### Examples

```
/api/v4/render/favicon/pure/blue/vortex.svg
/api/v4/render/og-card/spotlight/amber/cyan/pulse.svg?theme=dark&title=MyApp
/api/v4/render/hero/grid/neon/laser/spin.svg?theme=light&subtitle=tagline
/api/v4/render/avatar/pure/rose/breathe.svg
```

### V3 Legacy

```
GET /api/favicon/:color1/:color2?/:animation.svg
```

### Static Favicons

Pre-generated favicons are available in `/animated/` and as `favicon-{color}.svg` at root.

Regenerate with:

```bash
node generate-favicons.js
```

## Colors

**Base**: amber, blue, cyan, green, indigo, orange, purple, rose

**Creative**: white, gray, black, gold, platinum, champagne, neon, matrix, cyber, laser, plasma, void, emerald, sapphire, ruby, ocean, sunset, magma, mint, peach, lavender

Custom hex values are also supported (e.g., `ff6600`).

## Animations

50 animation styles organized by category:

| Category | Animations |
|---|---|
| Static | static |
| Zen | zen, breathe, levitate, glimmer |
| Spin | spin, spin-fast, smooth-spin, orbit-chase, compass, gyro, satellite, eclipse |
| Pulse | pulse, heartbeat, pulse-ring, strobe, nova, elastic |
| Flip | flip, orbit-tilt, vortex, harmony, sync, sway |
| Radar | radar, radar-sweep, signal |
| Glow | glow, aurora, nebula, corona, ripple-core |
| Morph | morph, morph-blob, morph-crystal |
| Bounce | bounce, bounce-drop, trampoline |
| Shake | shake, jitter, earthquake |
| Swing | swing, pendulum, wave-swing |
| Zoom | zoom-in, zoom-out, zoom-pulse |
| Slide | slide-in, slide-loop |
| Typewriter | typewriter-blink |

## License

[MIT](LICENSE)
