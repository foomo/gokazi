# Your first task

In this walkthrough you will:

1. Write a `gokazi.yaml` describing one task.
2. Start the underlying process by hand.
3. Use `gokazi list` to detect it.
4. Use `gokazi stop` to kill it.

## 1. Create the config

Create `gokazi.yaml` next to your project. We will manage `gokazi` itself — running `gokazi list` against another `gokazi list` is a tidy self-test.

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/foomo/gokazi/main/gokazi.schema.json
version: '1.0'

tasks:
  list:
    name: gokazi
    description: Long-running listing process
    path: $PWD/bin
    cwd: $PWD
    args: [list]
```

Field-by-field:

- `name: gokazi` — match any process whose OS name is `gokazi`.
- `path: $PWD/bin` — narrow to executables under `./bin/` (env-vars expand at match time).
- `cwd: $PWD` — narrow to the project directory.
- `args: [list]` — the cmdline must contain `list` (extras allowed).

## 2. Start the process

You need a `gokazi` binary in `./bin/`. From the gokazi repo:

```shell
make build
```

Then start one in the background:

```shell
./bin/gokazi list &
```

## 3. Detect it

```shell
./bin/gokazi list
```

You should see your `list` task with `running: true` and a PID.

## 4. Stop it

```shell
./bin/gokazi stop list
```

::: info What just happened
`gokazi stop` walked the OS process table, matched the task by `name`+`path`+`cwd`+`args`, and sent it a kill signal with a 5-second context timeout. There is no daemon — every command is a fresh enumeration.
:::

Next: [Managing tasks](/guide/managing-tasks).
