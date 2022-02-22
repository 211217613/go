.DEFAULT_GOAL := build
.PHONY:build fmt lint vet

fmt:
		go fmt ./...

lint: fmt
		golint ./...

vet: fmt
		go vet ./...

build: vet
		go build main.go
