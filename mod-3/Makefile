MAKEFLAGS += --silent

.DEFAULT_GOAL := build
.PHONY: build docker run test

VERSION := $$(git rev-parse --short HEAD)

build:
	go build -o mod-3

docker:
	docker build .. \
		-f $$(pwd)/Dockerfile \
		-t cloud-native-fizzbuzz/mod-3:latest \
		-t cloud-native-fizzbuzz/mod-3:$(VERSION)

run: build
	./mod-3

test:
	go test -v
