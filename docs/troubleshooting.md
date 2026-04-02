# Troubleshooting

## Docs build fails on GitHub Pages

The Pages workflow assumes `docs/package-lock.json` exists and uses `npm ci`.

- Ensure `docs/package.json` and `docs/package-lock.json` are committed.
- Ensure the workflow builds `docs/.vitepress/dist`.

## 403 on /cache/stats

Admin endpoints are disabled by default. Set `server.admin_key` and send `X-Admin-Key`.

## CORS issues

If you configure multiple origins, the server reflects the request origin only when it matches an allowed origin.

## Reverse proxy rate limiting

If you run behind a reverse proxy/CDN, ensure it forwards `X-Forwarded-For` or `X-Real-IP` so rate limiting can identify clients correctly.
