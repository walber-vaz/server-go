.PHONY: default run build test clean

include .env

default: run

run:
	@go run ./cmd/http/main.go

build:
	@CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o ./bin/$(APP_NAME) ./cmd/http/main.go

test:
	@go test -v ./ ...

clean:
	@rm -rf bin
	@rm -rf docs
