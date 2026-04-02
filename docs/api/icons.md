# Icon packs

## Endpoint

```
GET /api/v4/icons?pack=lucide&q=rocket
```

Query parameters:

- `pack` (optional): `lucide` or `hero`
- `q` (optional): search string

## Using an icon as a shape

Use the `shape` query parameter in the Render API:

```
/api/v4/render/favicon/pure/white/static.svg?shape=lucide:rocket
/api/v4/render/favicon/pure/white/static.svg?shape=hero:fire
```

## Notes

Icon keys are returned as `{pack}:{name}`. The dashboard provides browsing and search.
