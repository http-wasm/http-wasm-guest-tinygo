package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

func main() {
	pool.SetHandler(setHeader{})
}

type setHeader struct {
	api.UnimplementedHandler
}

func (setHeader) HandleRequest(_ api.Request, resp api.Response) (next bool, reqCtx uint32) {
	resp.Headers().Set("Content-Type", "text/plain")
	return // this is a benchmark, so skip the next handler.
}
