# Local dev servers

Scenario: you have a Go API and a Vite front-end. You want to start them however you already do, then use `gokazi` to see what's running and stop them selectively.

## `gokazi.yaml`

```yaml
version: '1.0'

tasks:
  web:
    name: node
    description: Vite dev server
    cwd: $PWD/web
    args: [vite]
  api:
    name: api-server
    description: Go API
    path: $PWD/bin
    cwd: $PWD
    args: [serve]
```

A few things to note:

- `web.name: node` because the dev server runs as `node /path/to/vite`. The cmdline contains `vite`, so we match on the arg, not the binary.
- `api.path: $PWD/bin` narrows to your local build, so a globally-installed binary by the same name elsewhere won't match.

## Workflow

```shell
# Start your servers however you already do.
(cd web && npm run dev &)
go run ./cmd/api serve &

# See what gokazi found.
gokazi list

# Restart just the API.
gokazi stop api
go run ./cmd/api serve &
```

If `gokazi list` shows `running: false` for something you can see in `ps`, walk the [troubleshooting checklist](/guide/troubleshooting).
