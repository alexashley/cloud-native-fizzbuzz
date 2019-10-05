MAKEFLAGS += --silent

.DEFAULT_GOAL := build
.PHONY: build run test

build:
	go build -o mod-3

run: build
	./mod-3

test:
	go test -v
