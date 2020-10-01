
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
	JWT_SIGNING_KEY=batman LOCAL_ENV=true go run ./app

make test: __run-go-critic
	go test ./...

generate-mocks:
	go generate ./...
# ###########

# Docker commands
docker-compose-build:
	docker-compose --file ./docker/docker-compose.yaml build

docker-compose-up:
	docker-compose --file ./docker/docker-compose.yaml up

# ###########

