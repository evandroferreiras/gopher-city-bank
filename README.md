# Gopher Bank City

![image](https://user-images.githubusercontent.com/3903012/94623063-c3948e00-0289-11eb-98df-9fb832c92aab.png)

This is the Gopher ~~Gotham~~ Bank City app.

## Technologies and frameworks

- Go 1.14+
- Echo Framework
- Swagger
- Logrus
- Testify
- Mockery
- Docker

## Environment Variable

| Variable | values           | Description |
| -------- | ---------------- | ----------- |
| LOCAL_ENV |  `true`/`false` | Identify if your running localhost, so the logs will be more friendly |


## Usage

To run localhost:
```shell
make run
```

To run using docker:
```shell
make docker-compose-up
```

To run unit tests:
```shell
make test
```

## Contributing

To contribute with this project, please prepare your setup. This project uses some tools to help and improve the developments:

-  [swag](https://github.com/swaggo/swag): This is used to generate the boilerplate code necessary for the swagger API documentation.
```
make setup-swag
```

- git hooks: We use pre-commit with some lints e go tools to assure a concise codebase.
```
make setup-githooks
```
### Directories

- [app](app/): Application code
- [docker](docker/): Dockerfiles
- [docs](docs/): Swagger files
