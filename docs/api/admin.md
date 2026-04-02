# Admin endpoints

Admin endpoints are disabled by default. To enable them:

1. Set `server.admin_key` in `config.yaml`
2. Send `X-Admin-Key: <value>` with each request

## Cache stats

```
GET /cache/stats
```

Returns cache tier counters (hits/misses) and basic information.

## Cache purge

```
POST /cache/purge
```

Clears all active cache tiers.
