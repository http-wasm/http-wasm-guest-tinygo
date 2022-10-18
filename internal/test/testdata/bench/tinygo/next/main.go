package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleFn = next
}

// next is the default, but doing this is more explicit.
func next(req api.Request, resp api.Response, next api.Next) {
	next()
}
