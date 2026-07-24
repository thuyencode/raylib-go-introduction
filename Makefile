.PHONY: dev build start lint format clean

BINARY=raylib-go-introduction

ifneq (,$(wildcard .env))
include .env
endif

dev:
	air

build:
	go build -tags $(BUILD_TAGS) -o $(BINARY) .

start:
	./$(BINARY)

lint:
	go vet ./...

format:
	go fmt ./...

clean:
	rm -f $(BINARY)
