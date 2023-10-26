.PHONY: clean fmt deps test build run

fmt:
	@go fmt ./...

clean:
	@rm -f alert-api *.db

deps:
	@go mod download

test:
	@go test -v ./...

build: fmt clean deps
	@go build -o alert-api cmd/alert/main.go

run: build
	@./alert-api