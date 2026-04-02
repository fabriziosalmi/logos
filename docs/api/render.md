# Render API

## Endpoints

```
GET /api/v4/render/{format}/{scene}/{color1}/{animation}.svg
GET /api/v4/render/{format}/{scene}/{color1}/{color2}/{animation}.svg
```

Path parameters:

- `format`: output size/preset (favicon, avatar, og-card, hero, and others)
- `scene`: background/layout preset (pure, spotlight, grid, split, and others)
- `color1`: primary color name or hex
- `color2`: optional gradient color name or hex
- `animation`: animation preset

## Query parameters

Common parameters:

- `shape`: built-in (`atom`, `shield`, `hexagon`, ...) or icon pack key (`lucide:rocket`, `hero:fire`)
- `theme`: theme preset (dark, light, and others)
- `texture`: texture preset (grain, glass, glitch, and others)
- `variant`: variant preset (outline, badge, glow, and others)
- `title`, `subtitle`: optional text overlay

Additional parameters:

- `stroke` (float): stroke width override
- `padding` (float): viewBox padding
- `alpha` (float): opacity (0.0–1.0)
- `badge` (int): badge count (1–999)
- `speed` (float): animation speed multiplier
- `direction` (string): reverse, alternate, alternate-reverse
- `hover` (bool): inject a `:hover` brightness rule
- `decorative` (bool): mark SVG as decorative (`aria-hidden="true"`)

## Examples

```
/api/v4/render/favicon/pure/blue/vortex.svg
/api/v4/render/og-card/spotlight/amber/cyan/pulse.svg?theme=dracula&title=MyApp
/api/v4/render/twitter-card/aurora-bg/rose-400/breathe.svg?shape=lucide:heart&variant=glow
```

## Response headers

- `ETag`: enables conditional requests (`If-None-Match`) and 304 responses
- `Cache-Control`: based on cache strategy
- `Vary`: includes `Accept-Encoding` and `Save-Data` (when enabled)
