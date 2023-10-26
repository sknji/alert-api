.PHONY: clean fmt deps test build run

fmt:
	@go fmt ./...

clean:
	@rm -f alert-api

deps:
	@go mod download

test:
	@go test -v ./...

build: fmt clean deps
	@go build -o alert-api cmd/alert/alert.go

run: build
	@./alert-api