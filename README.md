# Covid Summary

## Project structure

- `apperror/` a custom error
- `cmd/` a go main package
- `config/` a config loader
- `internal/` all service handling and core logic packages

## Files overview

### `.env`

the app config file
|Env|Description|
|-|-|
|`APP_ENV`|`development` or `production`<br>**\*`production` if not specified or not match**|
|`APP_PORT`|listening port<br>**\*defaults on `8000`**|
|`COVID_STAT_SERVER`|covid stat hostname to be fetched from|

### `export-env.sh`

exports all env variables in `.env` file

### `Makefile`

- `dev` auto reload go app with nodemon
- `run` run go app
- `test` run go test with exported coverage profile
- `cover` open coverage profile in html
