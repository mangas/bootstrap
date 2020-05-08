# go utils

# variables
GO_PKG_LIST = ???



## Print
go-print: 
	@echo
	@echo -- GO --
	@echo GO_PKG_LIST: $(GO_PKG_LIST)
	@echo


## boilerplate-update
go-boilerplate-update: ## go-boilerplate-update
	# See: https://github.com/lyft/boilerplate
	# Example: See: https://github.com/lyft/flytepropeller/tree/master/boilerplate
	# TODO: This will be redundant once we have BS releases working.
	@boilerplate/update.sh


## Build the code
go-build: ## go-build
	@echo Building
	@go build -v -o bs .

## Build the code to gobin
go-build-global: ## go-build-global
	@echo Building
	@go build -v -o $(GOPATH)/bin/bs .

## Run the code
go-run: ## go-run
	@echo Running
	@go run -v .

## Format with go-fmt
go-fmt: ## go-fmt
	@echo Formatting
	@go fmt .

## Lint with golangci-lint
go-lint: ## go-lint
	@echo Linting
	@golangci-lint run --no-config --issues-exit-code=0 --timeout=5m

## Run the tests
go-test: ## go-test
	@echo Running tests
	@go test -race -v .

## Run the tests with coverage
go-test-coverage: ## go-test-coverage
	@echo Running tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${GO_PKG_LIST}

## Display test coverage
go-display-coverage: ## go-display-coverage
	@echo Displaying test coverage
	@go tool cover -html=cover.out




