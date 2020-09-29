FROM golang:latest AS build

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

# Doing the go mod download separated so the cache will work to the packages
COPY go.sum go.mod $GOPATH/src/github.com/evandroferreiras/gopher-city-bank/
WORKDIR $GOPATH/src/github.com/evandroferreiras/gopher-city-bank/
RUN go mod download

COPY . $GOPATH/src/github.com/evandroferreiras/gopher-city-bank/

WORKDIR $GOPATH/src/github.com/evandroferreiras/gopher-city-bank/app/

RUN go build -o gopher-city-bank

ENTRYPOINT ./gopher-city-bank

EXPOSE 8585
