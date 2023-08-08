package main

import (
	log "github.com/http-wasm/http-wasm-guest-tinygo/examples/wasi"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/once"
)

func main() {
	once.Handle(log.Handler{})
}
