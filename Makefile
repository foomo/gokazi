.DEFAULT_GOAL	:= help
-include .makerc

# --- Validation --------------------------------------------------------------

husky=$(shell command -v husky 2> /dev/null)
ifndef husky
  $(error "missing executable 'husky', please install")
endif

# --- Targets -----------------------------------------------------------------

# This allows us to accept extra arguments
%:
	@:

## === Tasks ===

.PHONY: doc
## Run tests
doc:
	@open "http://localhost:6060/pkg/github.com/foomo/gokazi/"
	@godoc -http=localhost:6060 -play

.PHONY: test
## Run tests
test:
	@GO_TEST_TAGS=-skip go test -tags=safe -coverprofile=coverage.out -race ./...
	#@GO_TEST_TAGS=-skip go test -tags=safe -coverprofile=coverage.out -race -json ./... 2>&1 | tee /tmp/gotest.log | gotestfmt

.PHONY: lint
## Run linter
lint:
	@golangci-lint run

.PHONY: lint.fix
## Fix lint violations
lint.fix:
	@golangci-lint run --fix

.PHONY: tidy
## Run go mod tidy
tidy:
	@go mod tidy

.PHONY: outdated
## Show outdated direct dependencies
outdated:
	@go list -u -m -json all | go-mod-outdated -update -direct

## Install binary
install:
	@go build -o ${GOPATH}/bin/gokazi main.go

## Build binary
build:
	@mkdir -p bin
	@go build -o bin/gokazi main.go

## === Utils ===

.PHONY: help
## Show help text
help:
	@awk '{ \
			if ($$0 ~ /^.PHONY: [a-zA-Z\-\_0-9]+$$/) { \
				helpCommand = substr($$0, index($$0, ":") + 2); \
				if (helpMessage) { \
					printf "\033[36m%-23s\033[0m %s\n", \
						helpCommand, helpMessage; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^[a-zA-Z\-\_0-9.]+:/) { \
				helpCommand = substr($$0, 0, index($$0, ":")); \
				if (helpMessage) { \
					printf "\033[36m%-23s\033[0m %s\n", \
						helpCommand, helpMessage"\n"; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^##/) { \
				if (helpMessage) { \
					helpMessage = helpMessage"\n                        "substr($$0, 3); \
				} else { \
					helpMessage = substr($$0, 3); \
				} \
			} else { \
				if (helpMessage) { \
					print "\n                        "helpMessage"\n" \
				} \
				helpMessage = ""; \
			} \
		}' \
		$(MAKEFILE_LIST)
