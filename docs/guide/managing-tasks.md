# Managing tasks

## Listing

```shell
gokazi list
```

Output is JSON with one entry per configured task:

```json
{
  "list": {
    "config": { "name": "gokazi", "path": "...", "cwd": "...", "args": ["list"] },
    "running": true,
    "pid": 31415
  },
  "list-debug": {
    "config": { "name": "gokazi", "path": "...", "cwd": "...", "args": ["list", "--debug"] },
    "running": false,
    "pid": 0
  }
}
```

`running: false` and `pid: 0` mean the matcher found nothing — see [Troubleshooting](/guide/troubleshooting) if that surprises you.

## Stopping

```shell
gokazi stop <id>
```

Where `<id>` is the YAML key under `tasks:` (so `list` for the example above, **not** the process name). The kill signal runs with a 5-second context timeout; processes that ignore it will return an error.

## Inspecting merged config

When you layer multiple `-c` files (or pipe one in via `-`), use:

```shell
gokazi config
```

It prints the resolved configuration after all sources are merged and env-var expansion is applied. Useful when a task isn't matching and you want to confirm the final `path`/`cwd`/`args`.

## Multiple tasks at once

`gokazi` does not start processes — it observes and stops them. Pair it with a Makefile or justfile that owns the start side. See [Make / Just integration](/recipes/make-just-integration).

Next: [Troubleshooting](/guide/troubleshooting).
