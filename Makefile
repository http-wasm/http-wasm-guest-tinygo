goimports := golang.org/x/tools/cmd/goimports@v0.1.12
golangci_lint := github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0

.PHONY: test
test:
	@go test -v ./...
	@tinygo test -target=wasi -v ./...

.PHONY: test.e2e
test.e2e:
	@cd internal && go test ./... -v -timeout 120s

tinygo_sources := example/main.go $(wildcard internal/e2e/*/*.go)
PHONY: build.e2e
build.e2e: $(tinygo_sources)
	@for f in $^; do \
	    tinygo build -o $$(echo $$f | sed -e 's/\.go/\.wasm/') -scheduler=none --no-debug -target=wasi $$f; \
	done

golangci_lint_path := $(shell go env GOPATH)/bin/golangci-lint

$(golangci_lint_path):
	@go install $(golangci_lint)

.PHONY: lint
lint: $(golangci_lint_path)
	@CGO_ENABLED=0 $(golangci_lint_path) run --timeout 5m
	@CGO_ENABLED=0 $(golangci_lint_path) run --build-tags tinygo.wasm --timeout 5m

.PHONY: format
format:
	@find . -type f -name '*.go' | xargs gofmt -s -w
	@for f in `find . -name '*.go'`; do \
	    awk '/^import \($$/,/^\)$$/{if($$0=="")next}{print}' $$f > /tmp/fmt; \
	    mv /tmp/fmt $$f; \
	done
	@go run $(goimports) -w -local github.com/http-wasm/http-wasm-guest-tinygo `find . -name '*.go'`

.PHONY: check
check:
	@$(MAKE) lint
	@$(MAKE) format
	@go mod tidy
	@(cd internal; go mod tidy)
	@(cd internal/e2e; go mod tidy)
	@if [ ! -z "`git status -s`" ]; then \
		echo "The following differences will fail CI until committed:"; \
		git diff --exit-code; \
	fi

.PHONY: clean
clean: ## Ensure a clean build
	@go clean -testcache
