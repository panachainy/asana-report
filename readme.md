# Asana report

[![Coverage Status](https://coveralls.io/repos/github/panachainy/asar/badge.svg?branch=feature/prepair-config)](https://coveralls.io/github/panachainy/asar?branch=feature/prepair-config)

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

## REF

* https://sbstjn.com/blog/create-golang-cli-application-with-cobra-and-goxc/

* https://github.com/KEINOS/Hello-Cobra
