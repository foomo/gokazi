# Configuration

`gokazi` reads YAML configuration from one or more `-c` sources (use `-` for stdin). Sources merge top-down: later sources override earlier ones.

## Top level

| Field | Type | Required | Description |
|---|---|---|---|
| `version` | string | yes | Must equal `"1.0"`. Loading fails otherwise. |
| `tasks` | map<string, [Task](#task)> | yes | Task ID → task definition. The ID is the key passed to `gokazi stop <id>`. |

## Task

| Field | Type | Required | Env expansion | Description |
|---|---|---|---|---|
| `name` | string | yes | no | Process name to match. Typically the binary basename (e.g. `gokazi`, not `/usr/bin/gokazi`). |
| `description` | string | no | no | Human-readable description. Surfaced in `gokazi list` JSON. |
| `path` | string | no | yes | Directory containing the executable. Compared against `path.Dir(exe)`. |
| `cwd` | string | no | yes | Working directory of the running process. Compared after `path.Clean` + `os.ExpandEnv`. |
| `args` | []string | no | yes | Subset match: every entry must appear somewhere in `cmdline[1:]`. Order and extras don't matter. |

## Multi-source config

`gokazi -c shared.yaml -c overrides.yaml list` reads `shared.yaml` first, then `overrides.yaml` on top. Maps are merged by key; scalars and arrays are replaced.

```shell
# Read from stdin in addition to a file
cat <<'EOF' | gokazi -c gokazi.yaml -c - list
version: '1.0'
tasks:
  experimental:
    name: my-server
    args: [--port=9090]
EOF
```

See [Multi-source config](/recipes/multi-source-config) for layered-team-config patterns.

## Env-var expansion

Fields marked with env expansion run through Go's `os.ExpandEnv` at match time, so `$PWD`, `$HOME`, etc. resolve from the *current* gokazi invocation's environment — not the environment of the matched process.

```yaml
tasks:
  api:
    name: api-server
    path: $PWD/bin
    cwd: $PWD
    args: [--config=$ENV_NAME.json]
```

If the variable is unset, it expands to the empty string.

## Schema

The JSON schema is generated from the Go config types and lives at [`gokazi.schema.json`](https://raw.githubusercontent.com/foomo/gokazi/main/gokazi.schema.json). Wire it into your editor:

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/foomo/gokazi/main/gokazi.schema.json
version: '1.0'
tasks:
  ...
```
