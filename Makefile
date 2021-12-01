GOCMD = $(shell which go)
GOFMT = $(shell which gofmt)
GOFILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: help
help: ## - Show help message
	@printf "\033[32m\xE2\x9c\x93 usage: make [target]\n\n\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## - Build a service locally
	docker-compose build

.PHONY: run
run: ## - Run a service locally
	docker-compose up

.PHONY: code_format
code_format: ## - Format code
	$(GOFMT) -d -w $(GOFILES)

.PHONY: code_lint
code_lint: ## - Run a linter
	golangci-lint run	
