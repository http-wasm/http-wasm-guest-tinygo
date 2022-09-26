module github.com/http-wasm/http-wasm-guest-tinygo/internal

go 1.18

require (
	github.com/stretchr/testify v1.8.0
	github.com/tetratelabs/wazero v1.0.0-pre.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/http-wasm/http-wasm-guest-tinygo => ../
