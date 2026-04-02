# Getting started

## Run locally

```bash
make run
```

The server listens on `http://localhost:3000/`.

- Dashboard: `/`
- Health check: `/healthz`
- Render API: `/api/v4/render/...`

## Configuration file

The server loads `config.yaml` if present; otherwise it falls back to `config.sample.yaml`.

To create a local configuration:

```bash
cp config.sample.yaml config.yaml
```

Then edit `config.yaml` to register your apps and set optional features.

## Airgapped environments

The dashboard is served from embedded assets and uses local CSS/JS. The only runtime network calls are to the same server origin (the API itself).
