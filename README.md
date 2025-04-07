[![Build Status](https://github.com/foomo/gokazi/actions/workflows/pr.yml/badge.svg?branch=main&event=push)](https://github.com/foomo/gokazi/actions/workflows/pr.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/foomo/gokazi)](https://goreportcard.com/report/github.com/foomo/gokazi)
[![Coverage Status](https://coveralls.io/repos/github/foomo/gokazi/badge.svg?branch=main&)](https://coveralls.io/github/foomo/gokazi?branch=main)
[![GoDoc](https://godoc.org/github.com/foomo/gokazi?status.svg)](https://godoc.org/github.com/foomo/gokazi)

# gokazi

> Simple daemonless process (swaheli: task <> kazi) manager.

Works by comparing processes by arguments, cwd and executable path.

## Install

Install the latest release of the cli:

```shell
$ brew update
$ brew install foomo/tap/gokazi
```

## Usage

```shell
$ gokazi help
CLI process manager

Usage:
  gogazi [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Print config
  help        Help about any command
  list        List tasks
  stop        Stop task by id
  version     Print version

Flags:
      --debug   enable debug logs
  -h, --help    help for gogazi

Use "gogazi [command] --help" for more information about a command.
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

Please refer to the [CONTRIBUTING](.github/CONTRIBUTING.md) details and follow the [CODE_OF_CONDUCT](.github/CODE_OF_CONDUCT.md) and [SECURITY](.github/SECURITY.md) guidelines.

## License

Distributed under MIT License, please see license file within the code for more details.

_Made with ♥ [foomo](https://www.foomo.org) by [bestbytes](https://www.bestbytes.com)_
