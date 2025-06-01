BINARY_NAME=sersi
PKG_NAME=github.com/sersi-project/core

.PHONY: all build run lint clean test deps build-linux build-windows build-mac build-all

all: build

build:
	go build -o bin/$(BINARY_NAME) .

run:
	go run main.go version

lint:
	golangci-lint run

clean:
	rm -rf bin

test:
	go test ./...

deps:
	go mod tidy

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/$(BINARY_NAME) .

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/windows/$(BINARY_NAME).exe .

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/$(BINARY_NAME) .
	GOOS=darwin GOARCH=arm64 go build -o bin/mac-arm64/$(BINARY_NAME) .

build-all:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/$(BINARY_NAME) .
	GOOS=windows GOARCH=amd64 go build -o bin/windows/$(BINARY_NAME).exe .
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/$(BINARY_NAME) .
	GOOS=darwin GOARCH=arm64 go build -o bin/mac-arm64/$(BINARY_NAME) .

help:
	@echo "Makefile commands:"
	@echo "  make build        - Compile the project"
	@echo "  make run          - Run the project"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Remove build files"
	@echo "  make deps         - Clean up modules"
	@echo "  make build-linux  - Build for Linux"
	@echo "  make build-mac    - Build for MacOS"
	@echo "  make build-all    - Build for all platforms"
	@echo "  make lint         - Lint the code (requires golangci-lint)"
