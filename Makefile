.PHONY: build

build:
	go build -v ./cmd/apiserver
run:
	go run ./cmd/apiserver

.DEFAULT_GOAL := build