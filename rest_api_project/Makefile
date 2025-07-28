BINARY_NAME=rest-api
.DEFAULT_GOAL := run

build:
	GOARCH=amd64 GOOS=windows go build -o ./target/${BINARY_NAME}-windows main.go

run: build
	./target/${BINARY_NAME}-windows

clean:
	go clean
	rm ./target/${BINARY_NAME}-windows

execute: clean build run

