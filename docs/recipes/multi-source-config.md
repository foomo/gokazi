# Multi-source config

`gokazi` accepts repeated `-c` flags. Sources are merged in order, later wins.

## Pattern: shared + per-developer overrides

`gokazi.yaml` (committed):

```yaml
version: '1.0'
tasks:
  api:
    name: api-server
    path: $PWD/bin
    cwd: $PWD
    args: [serve]
```

`gokazi.local.yaml` (gitignored, per-developer):

```yaml
version: '1.0'
tasks:
  api:
    args: [serve, --port=9090]
  scratch:
    name: scratch
    cwd: $PWD/scratch
```

Run with both:

```shell
gokazi -c gokazi.yaml -c gokazi.local.yaml list
```

Result: `api` is overridden to use port 9090, and a new `scratch` task is added. Verify with `gokazi config` to see the post-merge config.

::: tip .gitignore
Add `gokazi.local.yaml` to your `.gitignore` so each developer can have their own without committing it.
:::

## Pattern: stdin as a source

The literal `-` source reads from stdin. Useful for templating with another tool:

```shell
envsubst < gokazi.template.yaml | gokazi -c - list
```

Or for one-shot adhoc tasks:

```shell
cat <<'EOF' | gokazi -c gokazi.yaml -c - stop scratch
version: '1.0'
tasks:
  scratch:
    name: my-experiment
EOF
```

## Pattern: layering by environment

`gokazi.yaml` for defaults, `gokazi.$ENV.yaml` for overrides:

```shell
gokazi -c gokazi.yaml -c gokazi.$ENV.yaml list
```

Set `ENV=ci` in CI, leave it unset locally — `gokazi.yaml` alone is the local default.
