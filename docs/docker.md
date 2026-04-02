# Docker

## Build

```bash
docker build -t logos:local .
```

## Run

```bash
docker run --rm -p 3000:3000 logos:local
```

The container ships with `config.sample.yaml` copied to `/config.yaml`. Override configuration by mounting a file and setting `CONFIG_PATH`.

Example:

```bash
docker run --rm -p 3000:3000 \
  -e CONFIG_PATH=/config.yaml \
  -v "$PWD/config.yaml:/config.yaml:ro" \
  logos:local
```
