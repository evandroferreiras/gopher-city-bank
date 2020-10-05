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

## Environment Variables

| Variable | values           | Description |
| -------- | ---------------- | ----------- |
| LOCAL_ENV |  `boolean` | Identify if your running localhost, so the logs will be more friendly |
| JWT_SIGNING_KEY | `string`  | Secret key used to encrypt Jwt |
| USE_MEMORY_DB | `boolean` | Identify your application will run using a memory based repository (collection of `maps`) or a real database |
| EXECUTE_AUTOMIGRATE | `boolean` | When true, the application will run as a Auto migration mode and will create the needed tables |
## Usage
To run automigrate and created the needed tables:
```shell
make run-automigrate
```

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

> When running (docker or local) you can access http://localhost:8585/api/healthcheck

## API Documentation

There is a swagger documentation, to check it out, just run and access http://localhost:8585/swagger/index.html

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
