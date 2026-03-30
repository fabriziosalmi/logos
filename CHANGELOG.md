# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [4.0.0] - 2026-03-30

### Added

- **V4 Cinematic API** with scene composition, themes, and typography support
- 50 animation styles: static, zen, breathe, spin, pulse, vortex, aurora, nebula, and more
- 23 color palettes: 8 base colors + 15 creative colors (neon, matrix, cyber, laser, etc.)
- 4 output formats: favicon (32x32), avatar (512x512), og-card (1200x630), hero (1920x1080)
- 4 scene modes: pure, spotlight, grid, split
- 5 theme variants: dark, light, solid, glass, auto
- Text overlay support with title and subtitle via query parameters
- Custom hex color input support
- Interactive dashboard with live preview at root `/`
- Real-time URL generation and clipboard copy
- V3 legacy API compatibility at `/api/favicon/`
- Static favicon generator script (`generate-favicons.js`)
- Pre-generated animated favicon variants in `/animated/`
- Cache headers for production use (1 year max-age)

## [0.0.1] - 2026-03-30

### Added

- Initial project setup
