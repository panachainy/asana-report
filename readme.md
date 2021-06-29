# Asana report

[![Coverage Status](https://coveralls.io/repos/github/panachainy/asana-report/badge.svg?branch=main)](https://coveralls.io/github/panachainy/asana-report?branch=main)

## Initial

```sh
mkdir -p newApp && cd newApp
cobra init --pkg-name github.com/spf13/newApp
```

### Cobra add

```sh
# ./asar serve
cobra add serve
# ./asar config
cobra add config
# add cmd with Subcommand call with `./asar config create`
cobra add create -p 'configCmd'
```

## Getting stared

### Development

Run with development

```sh
go run main.go
```

Build

```sh
make build
```

Install

```sh
make install
```

Try Run

```sh
make try
```

### ENV

Must set prefix when you set config with env.

Example

```env
export ASAR_PROJECT_BASE=te222234123123
```

## REF

* [Blog create golang cli](https://sbstjn.com/blog/create-golang-cli-application-with-cobra-and-goxc/)
* [Hello-Cobra](https://github.com/KEINOS/Hello-Cobra)
* [Issue-188](https://github.com/spf13/viper/issues/188#issuecomment-399884438)
* [Http to go](https://mholt.github.io/curl-to-go/)
* [gjson](https://github.com/tidwall/gjson)
