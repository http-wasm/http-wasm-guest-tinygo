module github.com/http-wasm/http-wasm-guest-tinygo/e2e

go 1.18

require (
	github.com/http-wasm/http-wasm-guest-tinygo v0.0.0
	github.com/http-wasm/http-wasm-host-go v0.0.0-20221028145646-90b87cdcf397
	github.com/stretchr/testify v1.8.0
	github.com/tetratelabs/wazero v1.0.0-pre.2.0.20221028145108-be33572289ac
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/http-wasm/http-wasm-guest-tinygo => ../../
