module github.com/http-wasm/http-wasm-guest-tinygo/e2e

go 1.18

require (
	github.com/http-wasm/http-wasm-guest-tinygo v0.0.0
	github.com/http-wasm/http-wasm-host-go v0.5.1
	github.com/stretchr/testify v1.8.4
	github.com/tetratelabs/wazero v1.2.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/http-wasm/http-wasm-guest-tinygo => ../../
