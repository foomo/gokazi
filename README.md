[![Go Report Card](https://goreportcard.com/badge/github.com/foomo/gokazi)](https://goreportcard.com/report/github.com/foomo/gokazi)
[![GoDoc](https://godoc.org/github.com/foomo/gokazi?status.svg)](https://godoc.org/github.com/foomo/gokazi)
[![Github All Releases](https://img.shields.io/github/downloads/foomo/gokazi/total.svg?style=flat-square)](https://github.com/foomo/gokazi/releases)
[![Docker Pulls](https://img.shields.io/docker/pulls/foomo/gokazi?style=flat-square)](https://hub.docker.com/r/foomo/gokazi)
[![GitHub stars](https://img.shields.io/github/stars/foomo/gokazi.svg?style=flat-square)](https://github.com/foomo/gokazi)

<p align="center">
  <img alt="gokazi-providers" src="docs/public/logo.png" width="400" height="400"/>
</p>

# gokazi

> Simple daemonless process (swaheli: task <> kazi) manager.

Works by comparing processes by arguments, cwd and executable path.

## Installation

## Installation

<details>
<summary><b>Homebrew</b> (macOS / Linux)</summary>

```shell
brew install foomo/tap/gokazi
```

See the [foomo/homebrew-tap](https://github.com/foomo/homebrew-tap) repository.

</details>

<details>
<summary><b>Docker</b></summary>

```shell
docker run --rm foomo/gokazi:latest --help
```

Multi-arch images (`amd64`, `arm64`) are published to [Docker Hub](https://hub.docker.com/r/foomo/gokazi).

</details>

<details>
<summary><b>mise</b></summary>

```shell
mise use github:foomo/gokazi
```

or run directly:

```shell
mise x github:foomo/gokazi -- --help
```

See [mise.jdx.dev](https://mise.jdx.dev).

</details>

<details>
<summary><b>Binary release</b></summary>

Download the archive for your OS/arch from the [releases page](https://github.com/foomo/gokazi/releases) and extract `gokazi` into your `$PATH`.

</details>

<details>
<summary><b>go install</b></summary>

```shell
go install github.com/foomo/gokazi/cmd/gokazi@latest
```

Requires Go 1.26+.

</details>


## Usage

```shell
$ gokazi help
CLI process manager

Usage:
  gokazi [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Print config
  help        Help about any command
  list        List tasks
  stop        Stop task by id
  version     Print version

Flags:
      --debug   enable debug logs
  -h, --help    help for gokazi

Use "gokazi [command] --help" for more information about a command.
```

## Configuration

```yaml
version: '1.0'

tasks:
  list:
    name: gokazi
    description: Own listing process
    path: $PWD/bin
    cwd: $PWD
    args: [list]
  list-debug:
    name: gokazi
    path: $PWD/bin
    cwd: $PWD
    args: [list, --debug]
```

## How to Contribute

Contributions are welcome! Please read the [contributing guide](docs/CONTRIBUTING.md).

![Contributors](https://contributors-table.vercel.app/image?repo=foomo/gokazi&width=50&columns=15)

## License

Distributed under MIT License, please see the [license](./LICENSE) file within the code for more details.

_Made with ♥ [foomo](https://www.foomo.org) by [bestbytes](https://www.bestbytes.com)_
