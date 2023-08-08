package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/examples/router"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

func main() {
	pool.SetHandler(router.Handler{})
}
