.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
build:
	go test -v -race -timeout 30s ./ 

.DEFAULT_GOAL := build