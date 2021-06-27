# Asana report

[![Coverage Status](https://coveralls.io/repos/github/panachainy/asana-report/badge.svg?branch=feature/prepair-config)](https://coveralls.io/github/panachainy/asana-report?branch=feature/prepair-config)

## Initial

```sh
mkdir -p newApp && cd newApp
cobra init --pkg-name github.com/spf13/newApp
```

### Cobra add

```sh
# ./asana-report serve
cobra add serve
# ./asana-report config
cobra add config
# add cmd with Subcommand call with `./asana-report config create`
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

## REF

* https://sbstjn.com/blog/create-golang-cli-application-with-cobra-and-goxc/

* https://github.com/KEINOS/Hello-Cobra
