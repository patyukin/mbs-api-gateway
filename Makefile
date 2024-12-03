.PHONY: gen lint lint-fix test

LOCAL_BIN:=$(CURDIR)/bin

gen:
	go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/api_gateway/main.go -o docs/

test:
	go test ./...

lint:
	golangci-lint cache clean
	golangci-lint run ./...

lint-fix:
	golangci-lint run --fix ./...
