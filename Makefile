.PHONY: build
build:
	go build -v ./cmd/jetcracks

.DEFAULT_GOAL := build
