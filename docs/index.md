# Logos API

Generatore di asset SVG dinamici (favicon, avatar, OG cards, hero) con dashboard embedded e deployment airgapped.

## Quick start

```bash
make run
```

Apri `http://localhost:3000/`.

## Config

- Crea `config.yaml` partendo da `config.sample.yaml`
- Imposta `server.admin_key` per abilitare gli endpoint admin (`/cache/stats`, `/cache/purge`)
