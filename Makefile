gofumpt       := mvdan.cc/gofumpt@v0.5.0
gosimports    := github.com/rinchsan/gosimports/cmd/gosimports@v0.3.8
golangci_lint := github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2

.PHONY: testdata
testdata:
	@$(MAKE) build.wat
	@$(MAKE) build.tinygo

tinygo_sources := $(wildcard examples/*/*.go) $(wildcard internal/test/testdata/*/*.go) $(wildcard internal/test/testdata/*/*/*.go) $(wildcard internal/test/testdata/*/*/*/*.go)
build.tinygo: $(tinygo_sources)
	@for f in $^; do \
	    tinygo build -o $$(echo $$f | sed -e 's/\.go/\.wasm/') -scheduler=none --no-debug -target=wasi $$f; \
	done

wat_sources := $(wildcard internal/test/testdata/*/*/*.wat)
build.wat: $(wat_sources)
	@for f in $^; do \
	    wat2wasm -o $$(echo $$f | sed -e 's/\.wat/\.wasm/') --debug-names $$f; \
	done

.PHONY: test
test:
	@go test -v ./...
	@tinygo test -target=wasi -v ./...

.PHONY: test.e2e
test.e2e:
	@cd internal/e2e && go test ./... -v -timeout 120s

.PHONY: bench
bench:
	@(cd internal/e2e; go test -run=NONE -bench=. .)

golangci_lint_path := $(shell go env GOPATH)/bin/golangci-lint

$(golangci_lint_path):
	@go install $(golangci_lint)

.PHONY: lint
lint: $(golangci_lint_path)
	@CGO_ENABLED=0 $(golangci_lint_path) run --timeout 5m
	@# not --build-tags tinygo.wasm as it triggers "could not load export data"

.PHONY: format
format:
	@go run $(gofumpt) -l -w .
	@go run $(gosimports) -local github.com/http-wasm/ -w $(shell find . -name '*.go' -type f)

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
