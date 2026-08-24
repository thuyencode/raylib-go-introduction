.PHONY: dev build build-windows start lint format clean

BINARY=main
ENV_FILE=.env

ifneq (,$(wildcard $(ENV_FILE)))
include $(ENV_FILE)
endif

dev:
	air

build:
	go build -tags $(BUILD_TAGS) -o $(BINARY) .

build-windows:
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -tags $(BUILD_TAGS) -ldflags "-s -w" -o $(BINARY).exe .

start:
	./$(BINARY)

lint:
	go vet ./...

format:
	go fmt ./...

clean:
	rm -f $(BINARY) $(BINARY).exe
