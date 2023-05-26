.PHONY: build
build:
	go build -o ./server.exe ./cmd/app

.PHONY: test
test:
	go test -v -timeout 30s ./...

.DEFAULT_GOAL := build