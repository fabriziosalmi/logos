# Caching

## HTTP caching

The Render API supports:

- `ETag` and `If-None-Match` (304 responses)
- `Cache-Control` suitable for CDN and browser caching
- stale-while-revalidate strategy (when enabled)

## Server-side tiers

The server supports multiple cache tiers configured in `config.yaml`:

- L1: in-process memory cache (always enabled)
- L2: filesystem cache (optional)
- L3: Redis/Valkey cache (optional)
- L4: CDN cache via response headers (optional)

## Admin operations

See [Admin endpoints](/api/admin) for `/cache/stats` and `/cache/purge`.
