# App shortcuts

## Endpoint

```
GET /app/{slug}/{format}.svg
```

The `slug` must exist in the `apps:` section of your configuration file.

## Example

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

Then request:

```
/app/example-app/favicon.svg
```

## Fallback

If the app is unknown, the server returns a placeholder SVG and sets `X-Logos-Fallback: true`.
