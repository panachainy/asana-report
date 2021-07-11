# Setup project

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
