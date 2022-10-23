package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleFn = setHeader
}

func setHeader(req api.Request, resp api.Response, next api.Next) {
	resp.Headers().Set("Content-Type", "text/plain")
}