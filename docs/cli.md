# CLI

The binary provides a `render` subcommand.

## Render from flags

```bash
./logos render --format=favicon --scene=pure --color=neon --animation=vortex > logo.svg
```

Common flags:

- `--format`, `--scene`
- `--color`, `--color2`
- `--animation`, `--theme`
- `--shape`, `--texture`, `--variant`
- `--title`, `--subtitle`

## Render a configured app

```bash
./logos render example-app > favicon.svg
```

This reads from `config.yaml` (or `config.sample.yaml` if `config.yaml` does not exist). Use `--config` to set a path.
