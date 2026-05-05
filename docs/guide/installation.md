# Installation

## Homebrew (macOS / Linux)

```shell
brew install foomo/tap/gokazi
```

The tap lives at [foomo/homebrew-tap](https://github.com/foomo/homebrew-tap).

## Docker

```shell
docker run --rm foomo/gokazi:latest --help
```

Multi-arch images (`amd64`, `arm64`) are published to [Docker Hub](https://hub.docker.com/r/foomo/gokazi).

## mise

```shell
mise use github:foomo/gokazi
```

Or run a one-off:

```shell
mise x github:foomo/gokazi -- --help
```

See [mise.jdx.dev](https://mise.jdx.dev) for project-level pinning.

## Binary release

Download the archive for your OS/arch from the [releases page](https://github.com/foomo/gokazi/releases) and extract `gokazi` into your `$PATH`.

## go install

Requires Go 1.26+.

```shell
go install github.com/foomo/gokazi/cmd/gokazi@latest
```

::: warning Path change
The install path moved from `github.com/foomo/gokazi` to `github.com/foomo/gokazi/cmd/gokazi` when the binary was relocated into a `cmd/` subdirectory. Update any scripts that pinned the old path.
:::

## Verify

```shell
gokazi version
```

Next: [Your first task](/guide/your-first-task).
