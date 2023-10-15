.POSIX:

all:
	mkdir -p dist
	go build -o dist/ ./cmd/...

clean:
	rm -rf ./dist

build:
	mkdir -p ./dist
	go build -o ./dist/ ./cmd/...

run:
	go run ./cmd/

test:
	go test ./... -p 1

test-ci: test-short

test-short:
	go test ./... --short

swagger:
	swag init -g ./cmd/main.go

migrate:
	go run ./cmd/migrate/main.go

.PHONY: run test test-ci test-short
