setup-githooks:
	pip install pre-commit
	pre-commit install
	go get github.com/fzipp/gocyclo
	GO111MODULE=on go get -v -u github.com/go-critic/go-critic/cmd/gocritic

run:
	go run ./app

make test:
	go test ./...

generate-mocks:
	go generate ./...
