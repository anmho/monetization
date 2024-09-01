


all: build

.PHONY: run
run: build
	@./bin/api

.PHONY: build
build:
	@go build -o ./bin/api ./cmd/api/main.go

.PHONY: watch
watch: build
	@air --build.cmd "make build" --build.bin "./bin/api"