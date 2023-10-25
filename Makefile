.PHONY: clean fmt deps build run

fmt:
	@go fmt ./...

clean:
	@rm -f alert-api

deps:
	@go mod download

build: fmt clean deps
	@go build -o alert-api cmd/alert/alert.go

run: build
	@./alert-api