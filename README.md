# Covid Summary

This project utilizes Clean Code Architecture containing the following layers.

- `*.model.go` holds app entities
- `*.handler.go` handles http request from gin framework
- `*.service.go` holds core logic
- `*.repository.go` fetches data from external api

## Project structure

- **`apperror/`** a custom error
- **`cmd/`** a go main package
- **`config/`** a config loader
- **`internal/`** all services handling and core logics
- **`.env`** the app config file
  |Env variables|Description|
  |-|-|
  |`APP_ENV`|`development` or `production`<br>**\*`production` if not specified nor matched**|
  |`APP_PORT`|listening port<br>**\*defaults on `8000`**|
  |`COVID_STAT_SERVER`|covid stat hostname to be fetched from|

- **`export-env.sh`** exports all env variables in `.env` file
- **`Makefile`** shortcut scripts
  - `dev` auto reload go app with nodemon
  - `run` run go app
  - `test` run go test with exported coverage profile
  - `cover` open coverage profile in html
- **`Dockerfile` `.dockerignore` `docker-compose.yaml`** containerized app
