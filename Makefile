include .env
export

run: swag-v1 ### run
	@go generate ./... && cp ./configs/default.config.yml ./configs/config.yml && clear && \
	go build -o toc-pd-captain ./cmd/app && ./toc-pd-captain
.PHONY: run

help: ## display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

swag-v1: ### swag init
	@echo 'package main' > ./swagger.go && \
	swag init -g internal/controller/http/router/router.go && \
	rm -rf ./swagger.go && \
	echo "" >> ./docs/swagger.json && git add ./docs
.PHONY: swag-v1

lint: ### check by golangci linter
	@golangci-lint run
.PHONY: lint

test: ### run test
	@go test -v -cover -race ./internal/...
.PHONY: test

migrate-up-all: ### migration up to latest
	@migrate -path migrations -database '$(PG_URL)$(DB_NAME)?sslmode=disable' up
.PHONY: migrate-up-all

migrate-down-last: ### migration down one step
	@migrate -path migrations -database '$(PG_URL)$(DB_NAME)?sslmode=disable' down 1
.PHONY: migrate-down-last

migrate-create:  ### create new migration
	@migrate create -ext sql -dir migrations -tz "Asia/Taipei" -format "2006010215" 'migration'
.PHONY: migrate-create
