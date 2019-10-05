MAKEFLAGS += --silent

.DEFAULT_GOAL := build
.PHONY: build run test

build:
	go build -o itoa

run: build
	./itoa

test:
	go test -v
