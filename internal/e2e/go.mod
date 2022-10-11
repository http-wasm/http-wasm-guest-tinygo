module github.com/http-wasm/http-wasm-guest-tinygo/internal

go 1.18

require (
	github.com/http-wasm/http-wasm-host-go v0.0.0-20221011003337-81d8f3916f05
	github.com/stretchr/testify v1.8.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/tetratelabs/wazero v1.0.0-pre.2.0.20221003082636-0b4dbfd8d6ca // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/http-wasm/http-wasm-guest-tinygo => ../../
