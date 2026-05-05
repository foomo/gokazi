# CLI

The `gokazi` CLI has these user-facing commands. All of them share the same config-loading flag (`-c`, repeatable, accepts `-` for stdin).

| Command | Purpose |
|---|---|
| [`gokazi`](/reference/cli/gokazi) | Root command. Without arguments prints help. |
| [`list`](/reference/cli/gokazi_list) | Match every configured task against running processes and print status as JSON. |
| [`stop`](/reference/cli/gokazi_stop) | Kill the process matched by a single task ID. |
| [`config`](/reference/cli/gokazi_config) | Print the merged, expanded configuration. |
| [`version`](/reference/cli/gokazi_version) | Print the binary version. |
| [`completion`](/reference/cli/gokazi_completion) | Generate a shell completion script. |

::: info Generated pages
The per-command pages below are generated from cobra command metadata by `cmd/docgen`. Run `make docs.cli` to regenerate them after changing flags or short/long descriptions.
:::
