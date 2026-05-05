.DEFAULT_GOAL:=help
-include .makerc

# --- Config -----------------------------------------------------------------

# Newline hack for error output
define br


endef

# --- Targets -----------------------------------------------------------------

# This allows us to accept extra arguments
%: .mise .lefthook
	@:

.PHONY: .mise
# Install dependencies
.mise:
ifeq (, $(shell command -v mise))
	$(error $(br)$(br)Please ensure you have 'mise' installed and activated!$(br)$(br)  $$ brew update$(br)  $$ brew install mise$(br)$(br)See the documentation: https://mise.jdx.dev/getting-started.html)
endif
	@mise install

.PHONY: .lefthook
# Configure git hooks for lefthook
.lefthook:
	@lefthook install --reset-hooks-path

### Tasks

.PHONY: check
## Run lint & tests
check: tidy generate lint test audit

.PHONY: lint
## Run linter
lint:
	@echo "〉golangci-lint run"
	@golangci-lint run

.PHONY: lint.fix
## Fix lint violations
lint.fix:
	@echo "〉golangci-lint run fix"
	@golangci-lint run --fix

.PHONY: generate
## Run go generate
generate:
	@echo "〉go generate"
	@go generate ./...

.PHONY: test
## Run tests
test:
	@echo "〉go test"
	@GO_TEST_TAGS=-skip go test -coverprofile=coverage.out -tags=safe ./...

.PHONY: test.race
## Run tests
test.race:
	@echo "〉go test with -race"
	@GO_TEST_TAGS=-skip go test -coverprofile=coverage.out -tags=safe -race ./...

.PHONY: build
## Build binary
build:
	@echo "〉building bin/gokazi"
	@rm -f bin/gokazi
	@go build -o bin/gokazi ./cmd/gokazi/gokazi.go

.PHONY: build.debug
## Build binary in debug mode
build.debug:
	@echo "〉building debug bin/gokazi"
	@rm -f bin/gokazi
	@go build -gcflags "all=-N -l" -o bin/gokazi ./cmd/gokazi/gokazi.go

.PHONY: install
## Run go install
install:
	@echo "〉installing $$(go env GOPATH)/bin/gokazi"
	@go install -a ./cmd/gokazi

.PHONY: install.debug
## Run go install with debug
install.debug:
	@echo "〉installing debug $$(go env GOPATH)/bin/gokazi"
	@go install -a -gcflags "all=-N -l" ./cmd/gokazi/gokazi.go

### Security

.PHONY: audit
## Run security audit
audit:
	@echo "〉security audit"
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@govulncheck ./...

### Dependencies

.PHONY: tidy
## Run go mod tidy
tidy:
	@echo "〉go mod tidy"
	@go mod tidy

.PHONY: outdated
## Show outdated direct dependencies
outdated:
	@echo "〉go mod outdated"
	@go list -u -m -json all | go-mod-outdated -update -direct

.PHONY: upgrade
## Show outdated direct dependencies
upgrade:
	@echo "〉go mod upgrade"
	@go list -u -m -f '{{if and (not .Indirect) .Update}}{{.Path}}{{end}}' all | xargs -n1 -I{} go get {}@latest
	@$(Make) tidy

### Documentation

.PHONY: docs.cli
## Regenerate CLI reference markdown
docs.cli:
	@echo "〉generating docs/reference/cli"
	@go run ./cmd/docgen

.PHONY: docs
## Open docs (regenerates CLI ref first)
docs: docs.cli
	@echo "〉starting docs"
	@cd docs && bun install && bun run dev

.PHONY: docs.build
## Build docs (regenerates CLI ref first)
docs.build: docs.cli
	@echo "〉building docs"
	@cd docs && bun install && bun run build

.PHONY: godocs
## Open go docs
godocs:
	@echo "〉starting go docs"
	@go doc -http

### Utils

.PHONY: help
# https://patorjk.com/software/taag/#p=display&f=Tmplr&t=gokazi&x=none&v=4&h=4&w=80&we=false
help: g=\033[0;32m
help: b=\033[0;34m
help: w=\033[0;90m
help: e=\033[0m
## Show help text
help:
	@echo "$(g)"
	@echo "    ┓    •"
	@echo "┏┓┏┓┃┏┏┓┓┓"
	@echo "┗┫┗┛┛┗┗┻┗┗"
	@echo " ┛"
	@echo "with ❤ foomo by bestbytes"
	@echo "$(e)"
	@echo "$(b)Usage:$(e)\n  make [task]"
	@awk '{ \
		if($$0 ~ /^### /){ \
			if(help) printf "  %-21s $(w)%s$(e)\n\n", cmd, help; help=""; \
			printf "$(b)\n%s:$(e)\n", substr($$0,5); \
		} else if($$0 ~ /^[a-zA-Z0-9._-]+:/){ \
			cmd = substr($$0, 1, index($$0, ":")-1); \
			if(help) printf "  %-21s $(w)%s$(e)\n", cmd, help; help=""; \
		} else if($$0 ~ /^##/){ \
			help = help ? help "\n                        " substr($$0,3) : substr($$0,3); \
		} else if(help){ \
			print "\n                        $(w)" help "$(e)\n"; help=""; \
		} \
	}' $(MAKEFILE_LIST)
	@echo ""
