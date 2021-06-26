# Asana report

## Initial

```sh
mkdir -p newApp && cd newApp
cobra init --pkg-name github.com/spf13/newApp
```

### Cobra add

```sh
# ./ar serve
cobra add serve
# ./ar config
cobra add config
# add cmd with Subcommand call with `./ar config create`
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

https://sbstjn.com/blog/create-golang-cli-application-with-cobra-and-goxc/
