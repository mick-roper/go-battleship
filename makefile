.PHONY: all

all: clean install test build

clean:
	rm -rf bin/

install:
	go get ./...

test:
	go test ./...

build:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/game entrypoint/main.go