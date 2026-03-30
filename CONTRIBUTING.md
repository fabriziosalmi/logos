# Contributing

Contributions are welcome!

## Getting Started

```bash
git clone https://github.com/fabriziosalmi/logos.git
cd logos
node server.js
```

Open `http://localhost:3000` to see the dashboard.

## Adding Animations

Animations are defined in `server.js` in the `animations` object. Each animation is a CSS string applied inside the SVG `<style>` block. Target these classes:

- `.core` - the central dot
- `.orbit` - the orbital ring
- `.outer` - the outer ring

## Adding Colors

Add entries to `baseColors` or `creativeColors` in `server.js`. Values are 6-digit hex codes without `#`.

## Pull Requests

1. Fork the repo and create a branch from `main`
2. Test your changes locally
3. Verify SVG output renders in browsers
4. Submit a PR with a clear description
