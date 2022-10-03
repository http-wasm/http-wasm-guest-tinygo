module github.com/http-wasm/http-wasm-guest-tinygo/internal/e2e

// Match min version in commit.yaml
go 1.18

require github.com/http-wasm/http-wasm-guest-tinygo v0.0.0

require github.com/tetratelabs/tinymem v0.1.0 // indirect

replace github.com/http-wasm/http-wasm-guest-tinygo => ../../
