# Asana report

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

Run

```sh
go run main.go
```

### Publish way

Publish

```sh
publish.sh
```

Run

```sh
./asana-report serve
```
