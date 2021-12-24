# Go Util

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gookit/goutil?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/gookit/goutil)](https://github.com/snail-plus/goutil)
[![GoDoc](https://godoc.org/github.com/snail-plus/goutil?status.svg)](https://pkg.go.dev/github.com/snail-plus/goutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/snail-plus/goutil)](https://goreportcard.com/report/github.com/snail-plus/goutil)
[![Unit-Tests](https://github.com/snail-plus/goutil/workflows/Unit-Tests/badge.svg)](https://github.com/snail-plus/goutil/actions)
[![Coverage Status](https://coveralls.io/repos/github/gookit/goutil/badge.svg?branch=master)](https://coveralls.io/github/gookit/goutil?branch=master)

💪 Useful utils for the Go: string, array/slice, map, format, CLI, ENV, filesystem, testing and more.

- `arrutil` array/slice util functions
- `dump`  Simple variable printing tool, printing slice, map will automatically wrap each element and display the call location
- `cliutil` CLI util functions
- `envutil` ENV util for check current runtime env information
- `fmtutil` format data util functions
- `fsutil` filesystem util functions
- `jsonutil` JSON util functions
- `maputil` map util functions
- `mathutil` math util functions
- `netutil` network util functions
- `strutil` string util functions
- `testutil` test help util functions

> **[中文说明](README.zh-CN.md)**

## GoDoc

- [Godoc for github](https://pkg.go.dev/github.com/snail-plus/goutil)

## Packages
{{pgkFuncs}}
## Code Check & Testing

```bash
gofmt -w -l ./
golint ./...
go test ./...
```

## Gookit packages

  - [gookit/ini](https://github.com/snail-plus/ini) Go config management, use INI files
  - [gookit/rux](https://github.com/snail-plus/rux) Simple and fast request router for golang HTTP
  - [gookit/gcli](https://github.com/snail-plus/gcli) Build CLI application, tool library, running CLI commands
  - [gookit/slog](https://github.com/snail-plus/slog) Lightweight, easy to extend, configurable logging library written in Go
  - [gookit/color](https://github.com/gookit/color) A command-line color library with true color support, universal API methods and Windows support
  - [gookit/event](https://github.com/snail-plus/event) Lightweight event manager and dispatcher implements by Go
  - [gookit/cache](https://github.com/snail-plus/cache) Generic cache use and cache manager for golang. support File, Memory, Redis, Memcached.
  - [gookit/config](https://github.com/snail-plus/config) Go config management. support JSON, YAML, TOML, INI, HCL, ENV and Flags
  - [gookit/filter](https://github.com/snail-plus/filter) Provide filtering, sanitizing, and conversion of golang data
  - [gookit/validate](https://github.com/snail-plus/validate) Use for data validation and filtering. support Map, Struct, Form data
  - [gookit/goutil](https://github.com/snail-plus/goutil) Some utils for the Go: string, array/slice, map, format, cli, env, filesystem, test and more
  - More, please see https://github.com/snail-plus

## License

[MIT](LICENSE)
