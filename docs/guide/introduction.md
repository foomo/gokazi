# Introduction

`gokazi` (Swahili for *task*) is a small CLI process manager. Unlike `systemd`, `pm2`, or `foreman`, it has **no supervising daemon and no PID file**. Every invocation of `gokazi list` or `gokazi stop` enumerates the live OS process table and matches each configured task against the running processes.

## Mental model

Think of gokazi as a **declarative `ps` filter with a kill button**. You describe each task in YAML — what binary it runs, where, and with which args — and gokazi tells you which of those tasks are currently alive on your machine, by PID. You can then stop them by ID.

::: tip Matching rules
A configured task matches a process when **all** of these line up:

- `name` equals the process name (typically the binary basename)
- `path` (if set) equals the directory containing the executable
- `cwd` (if set) equals the process working directory
- every string in `args` (if set) appears in the process command-line

See [Troubleshooting](/guide/troubleshooting) for the full checklist.
:::

## When to reach for gokazi

- You start dev servers from a Makefile or justfile and want a clean way to find and stop them later.
- You need to detect whether a long-running command is already running before starting a new one.
- You want a single config file shared across the team that maps memorable IDs to "the process I'm running right now".

## When **not** to reach for gokazi

- You need lifecycle management with restarts, log rotation, or supervision — use `systemd`, `pm2`, or `supervisord`.
- You need to *start* the processes you manage — gokazi observes, it does not spawn.

Next: [Installation](/guide/installation).
