package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

var body = []byte("hello world")

func main() {
	pool.SetHandler(writeBody{})
}

type writeBody struct {
	api.UnimplementedHandler
}

func (writeBody) HandleRequest(_ api.Request, resp api.Response) (next bool, reqCtx uint32) {
	resp.Body().Write(body)
	return // this is a benchmark, so skip the next handler.
}
