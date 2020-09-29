setup-githooks:
	pip install pre-commit
	pre-commit install
	go get github.com/fzipp/gocyclo
	GO111MODULE=on go get -v -u github.com/go-critic/go-critic/cmd/gocritic

setup-swag:
	go get github.com/swaggo/swag/cmd/swag

generage-swagger:
	swag init -d ./app

run-go-critic:
	gocritic check ./app/...

run:generage-swagger
	LOCAL_ENV=true go run ./app

make test:
	go test ./...

generate-mocks:
	go generate ./...
