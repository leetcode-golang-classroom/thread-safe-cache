.PHONY=build

build:
	@go build -o bin/main cmd/main.go

run: build
	@./bin/main

test:
	@go test -cover ./... -v --race