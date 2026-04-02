# Contributing

Contributions are welcome!

## Getting Started

```bash
git clone https://github.com/fabriziosalmi/logos.git
cd logos
make dev
```

Open `http://localhost:3000` to see the dashboard.

## Adding Animations

Animations are defined in [animation.go](file:///Users/fab/Documents/git/logos/internal/svg/animation.go) as CSS fragments applied inside the SVG `<style>` block. Target these classes:

- `.core` — the inner symbol
- `.orbit` — the secondary ring/shape
- `.outer` — the outer boundary

## Adding Colors

Named colors live in [palette.go](file:///Users/fab/Documents/git/logos/internal/svg/palette.go). Values are 6-digit hex codes without `#`.

## Adding Shapes

- Built-in shapes are implemented in [shape.go](file:///Users/fab/Documents/git/logos/internal/svg/shape.go).
- Icon packs are embedded from [web/icons](file:///Users/fab/Documents/git/logos/web/icons) and loaded at startup; see [iconpack.go](file:///Users/fab/Documents/git/logos/internal/svg/iconpack.go).

## Pull Requests

1. Fork the repo and create a branch from `main`
2. Run `go test ./...` and keep the dashboard usable
3. Verify SVG output renders in browsers
4. Submit a PR with a clear description
