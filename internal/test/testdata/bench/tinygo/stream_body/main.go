package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleFn = readBody
}

var empty = make([]byte, 0)

func readBody(req api.Request, resp api.Response, next api.Next) {
	size, _ := req.Body().Read(empty)
	_, _ = req.Body().Read(make([]byte, size))
}
