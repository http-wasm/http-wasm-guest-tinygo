package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleFn = removeHeader
}

func removeHeader(req api.Request, resp api.Response, next api.Next) {
	resp.Headers().Remove("Set-Cookie")
}
