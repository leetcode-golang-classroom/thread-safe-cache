.PHONY=build

build:
	@go build -o bin/main cmd/cache.go

run: build
	@./bin/main

test:
	@go test -cover ./... -v --race