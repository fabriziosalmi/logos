# API

## Render (V4)

```
GET /api/v4/render/{format}/{scene}/{color1}/{animation}.svg
GET /api/v4/render/{format}/{scene}/{color1}/{color2}/{animation}.svg
```

Parametri query comuni:

- `shape`: built-in (`atom`, `shield`, …) o pack (`lucide:rocket`, `hero:fire`)
- `theme`: preset (`dark`, `dracula`, `vercel`, …)
- `texture`: `grain`, `glass`, `glitch`, …
- `variant`: `outline`, `badge`, `glow`, …

## Icons

```
GET /api/v4/icons?pack=lucide&q=rocket
```
