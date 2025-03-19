ifneq ("$(wildcard .env)","")
$(info using .env)
include .env
export $(shell sed 's/=.*//' .env)
endif

CONTAINER_NAME=task-system
COMPOSEV2 := $(shell docker compose version)

ifdef COMPOSEV2
    COMMAND=docker compose
else
    COMMAND=docker-compose
endif

COMPOSEV2 := $(shell docker compose version)

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	@go run ./cmd/...

current_time = $(shell date "+%Y-%m-%dT%H:%M:%S%z")
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

ISALPINE := $(shell grep 'Alpine' /etc/os-release  -c)
musl=
ifeq ($(ISALPINE), 2)
        musl=-tags musl
endif

## build/api: build the cmd/api application
.PHONY: build
build:
	@echo 'Building...'
	go build ${musl} -ldflags=${linker_flags} -o=./bin/application ./cmd

## clean/apps: clear generated bin files
.PHONY: clean/apps
clean/apps:
	@echo 'Remove builded apps'
	@rm -rf ./bin

## docker/build: build the local environment for development
.PHONY: docker/build
docker/build:
	docker build -t task-system:latest -f .setup/build/Dockerfile .

## docker/up: start the local http server
.PHONY: docker/up
docker/up:
	$(COMMAND) up -d

## docker/down: shutdown the running containers
.PHONY: docker/down
docker/down:
	$(COMMAND) down

## docker/logs: start the local log
.PHONY: docker/logs
docker/logs:
	docker logs --tail 1000 -f task-system

## test: test all code
.PHONY: test
test:
	go test -race -vet=off -coverpkg ./internal/... -v -coverprofile=cover.out ./...
	go tool cover -func=cover.out

## generate: generate mocks
.PHONY: generate
generate:
	ROOT_DIR=$(shell pwd) go generate ./...

## tidy: tidy dependencies
.PHONY: tidy
tidy:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	golangci-lint run --fix