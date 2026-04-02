# Generative API

## Endpoints

```
GET /api/v4/gen/render?prompt=...&format=...
GET /api/v4/gen/resolve?prompt=...
```

- `resolve` returns the resolved parameters as JSON (debug).
- `render` returns an SVG using the resolved parameters.

## Parameters

- `prompt` (required): free text used to derive format/theme/colors/shape/scene/animation
- `format` (optional, render only): forces a target format

## Notes

The generative endpoint is intended for convenience and prototyping. For stable asset pipelines, prefer the explicit Render API endpoints.
