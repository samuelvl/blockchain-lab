# Variables shared between targets
APP_NAME        := Blockchain Lab
APP_DESCRIPTION := Learn about blockchain technology
APP_VERSION     := 1.0.0
GO_PROJECT      := github.com/samuelvl/blockchain-lab
BIN_PATH        := bin
BIN_NAME        := blockchain-lab


##@ General
.PHONY: help

help: ## Display this help.
	@awk 'BEGIN { \
			FS = ":.*##"; \
			printf "\nUsage:\n  make \033[36m<target>\033[0m\n" \
		} \
		/^[a-zA-Z_0-9-]+:.*?##/ { \
			printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 \
		} \
		/^##@/ { \
			printf "\n\033[1m%s\033[0m\n", substr($$0, 5) \
		} ' $(MAKEFILE_LIST)


##@ Golang dependency management
.PHONY: update-deps download-deps deps

update-deps: ## Update go.mod with last dependencies versions.
	$(info • Updating dependencies...)
	@go get -v -u all

download-deps: ## Download dependencies in the vendor folder.
	$(info • Downloading vendor dependencies...)
	@go mod tidy -v
	@go mod vendor -v

deps: update-deps download-deps


##@ Golang testing
.PHONY: test

test: ## Run unit tests.
	$(info • Running unit tests...)
	@go test -v -race ./pkg/...


##@ Golang building
.PHONY: build

build: ## Generate binaries.
	$(info • Generating $(BIN_NAME) binary...)
	@go build \
		-work \
		-race \
		-o $(BIN_PATH)/$(BIN_NAME)


##@ Golang execution
.PHONY: run

run: ## Run the app locally.
	@sh -c '$(BIN_PATH)/$(BIN_NAME)'
