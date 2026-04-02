# Configuration file

## Overview

Logos reads configuration from YAML:

- `config.yaml` (local, not committed)
- `config.sample.yaml` (repository sample and fallback)

Set `CONFIG_PATH` to load a different file path.

## server

- `host`, `port`: bind address and port
- `admin_key`: enables admin endpoints; requests must include `X-Admin-Key`
- timeouts: read/write/idle/shutdown

## apps

Apps are exposed through:

```
GET /app/{slug}/{format}.svg
```

Example:

```yaml
apps:
  example-app:
    color: purple
    animation: vortex
    theme: dark
    title: Example App
    subtitle: Generative Assets
    shape: atom
```

## defaults

Default values applied when an app entry omits a field (format, scene, theme, texture, shape, animation).

## cache

See [Caching](/caching) for strategy, tiers (L1/L2/L3/L4) and related headers.

## cors

`allowed_origins` controls the `Access-Control-Allow-Origin` header. If you list multiple origins, the server matches against the request `Origin` and reflects a single allowed origin.

## security

See [Security headers](/security) for CSP, HSTS and rate limiting.
