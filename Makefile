# Load the .env file if it exists
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

build:
	@go build -o bin/app ./cmd/

run: build
	@./bin/app

dev: 
	@air

# Run tests with verbose output and coverage
test:
	@go test -v ./... -cover

# Run tests with coverage output and preview in a browser
test-preview:
	@go test ./filename/ -coverprofile=coverage.out 
	@go tool cover -html=coverage.out

.PHONY: run test build dev test-preview