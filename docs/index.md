# Logos API

Logos is an SVG asset generation API (favicons, avatars, OG cards, heroes) with an embedded dashboard. It runs as a single Go binary and can be deployed without runtime external web dependencies.

## Capabilities

- Render SVG assets by format, scene, colors, animation, and theme
- Built-in shapes plus embedded icon packs (Lucide, Heroicons)
- Optional textures and visual variants
- HTTP caching (ETag + stale-while-revalidate) with multi-tier server cache
- Optional generative endpoint that resolves a text prompt into render parameters

## Quick start (local)

```bash
make run
```

Open `http://localhost:3000/`.

## Configuration

- Create `config.yaml` starting from `config.sample.yaml`
- Set `server.admin_key` to enable admin endpoints (`/cache/stats`, `/cache/purge`)

## Documentation development

```bash
cd docs
npm ci
npm run docs:dev
```

## Where to start

- [Getting started](/getting-started)
- [Render API](/api/render)
- [Configuration file](/configuration)
