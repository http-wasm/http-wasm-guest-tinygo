package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/examples/router"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/once"
)

func main() {
	once.Handle(router.Handler{})
}
