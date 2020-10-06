
LOCALHOST_MYSQL_DSN := "api-user:api-password@tcp(127.0.0.1:3306)/gophercitybank?charset=utf8mb4&parseTime=True&loc=Local"

# Developer setup commands
setup-githooks:
	pip install pre-commit
	pre-commit install
	go get github.com/fzipp/gocyclo
	GO111MODULE=on go get -v -u github.com/go-critic/go-critic/cmd/gocritic

setup-swag:
	go get github.com/swaggo/swag/cmd/swag
# ###########

# Localhost developing commands
__generage-swagger:
	swag init -d ./app

__run-go-critic:
	gocritic check ./app/...

run:__generage-swagger __run-go-critic
	JWT_SIGNING_KEY=batman \
	LOCAL_ENV=true \
	USE_MEMORY_DB=$(USE_MEMORY_DB) \
	MYSQL_DSN=$(LOCALHOST_MYSQL_DSN) \
	go run -race ./app

make test: __run-go-critic
	go test -race ./...

generate-mocks:
	go generate ./...
# ###########

# Database commands
run-automigrate:
		MYSQL_DSN=$(LOCALHOST_MYSQL_DSN) \
		EXECUTE_AUTOMIGRATE=true go run -race ./app
# ############

# Docker commands
docker-compose-build:
	docker-compose --file ./docker/docker-compose.yaml build

docker-compose-up:
	docker-compose --file ./docker/docker-compose.yaml up

docker-compose-up-mysql:
	docker-compose --file ./docker/docker-compose.yaml up mysql_server
# ###########

