# Security headers

The server sets a baseline set of headers and allows configuring CSP and HSTS via `config.yaml`.

## Default headers

- `X-Content-Type-Options: nosniff`
- `X-Frame-Options: DENY`
- `Referrer-Policy: no-referrer`
- `Permissions-Policy`: conservative defaults

## Content Security Policy (CSP)

Configure `security.csp`. The default policy is strict enough for the embedded dashboard and does not require external domains.

## HSTS

Configure `security.hsts` when serving over HTTPS.

## Rate limiting

Configure `security.rate_limit`. The limiter uses IP extraction compatible with reverse proxies (supports `X-Forwarded-For` and `X-Real-IP`).
