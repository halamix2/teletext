
.PHONY: all build test lint run check

all: test lint build


check: lint test

run:
	go run  ./cmd/updater

build:
	go build  -o ./updater ./cmd/updater 


test:
	go test -v ./...

lint:
	golangci-lint run ./...    
