package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

func main() {
	pool.SetHandler(readBody{})
}

type readBody struct {
	api.UnimplementedHandler
}

var empty = make([]byte, 0)

func (readBody) HandleRequest(req api.Request, _ api.Response) (next bool, reqCtx uint32) {
	size, _ := req.Body().Read(empty)
	_, _ = req.Body().Read(make([]byte, size))
	return // this is a benchmark, so skip the next handler.
}
