# Troubleshooting

## "I configured a task but `gokazi list` says it isn't running"

`findProcess` matches on **all** of `name`, `path`, `cwd`, and `args` simultaneously. Walk the checklist below — each step links to the matching rule.

### 1. Does `name` match the process name (not the path)?

`task.name` is compared against the OS process name as `gopsutil` reports it — typically the binary basename. For a binary at `/usr/local/bin/gokazi`, that's `gokazi`. Verify with:

```shell
ps -eo pid,comm | grep <your-binary>
```

### 2. If `path` is set, does it match the *directory* of the executable?

This is the most common gotcha. `task.path` is compared against `path.Dir(exe)`, **not** the full executable path.

| Scenario | `task.path` | Process `Exe` | Matches? |
|---|---|---|---|
| Binary at `/usr/bin/foo` | `/usr/bin` | `/usr/bin/foo` | ✓ |
| Binary at `/usr/bin/foo` | `/usr/bin/foo` | `/usr/bin/foo` | ✗ |

If you don't care about the path, leave `path` unset.

### 3. If `cwd` is set, does it match exactly?

`task.cwd` is compared against the process working directory after path-cleaning and env-var expansion. Symlinked roots can bite you here — resolve them in advance with `realpath`.

### 4. Do **all** strings in `args` appear in the cmdline?

Args are a **subset match**: every string in `task.args` must appear somewhere in `cmdline[1:]` (i.e. excluding the binary itself). Order doesn't matter, and extra cmdline args are fine.

| `task.args` | Process cmdline | Matches? |
|---|---|---|
| `[list]` | `gokazi list` | ✓ |
| `[list]` | `gokazi list --debug` | ✓ |
| `[list, --debug]` | `gokazi list` | ✗ |
| `[--debug, list]` | `gokazi list --debug` | ✓ |

### 5. Are env-vars expanding the way you expect?

`path`, `cwd`, and `args` all run through `os.ExpandEnv`. Run `gokazi config` to see the post-expansion values. A common mismatch: launching the underlying process from one shell where `$PWD` resolves one way, and running `gokazi list` from another where it resolves differently.

## "I see the right PID but `stop` says not running"

The process likely exited between the two enumerations (`list` and `stop` are separate scans). Re-run `gokazi list`; if it's gone, you're done.

## Still stuck?

Run with `--debug`:

```shell
gokazi --debug list
```

That prints the matcher's intermediate state. If you've narrowed the problem to a specific match field, [open an issue](https://github.com/foomo/gokazi/issues) with the `gokazi.yaml`, the `ps` output, and the debug log.
