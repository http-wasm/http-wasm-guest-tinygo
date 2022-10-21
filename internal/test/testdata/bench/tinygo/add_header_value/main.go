package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleFn = addHeader
}

func addHeader(req api.Request, resp api.Response, next api.Next) {
	resp.Headers().Add("Set-Cookie", "a=b")
}
