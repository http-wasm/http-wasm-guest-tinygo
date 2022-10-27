package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleRequestFn = readBody
}

var empty = make([]byte, 0)

func readBody(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	size, _ := req.Body().Read(empty)
	_, _ = req.Body().Read(make([]byte, size))
	return // this is a benchmark, so skip the next handler.
}
