package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

var body = []byte("hello world")

func main() {
	handler.HandleFn = writeResponseBody
}

func writeResponseBody(req api.Request, resp api.Response, next api.Next) {
	resp.Body().Write(body)
}
