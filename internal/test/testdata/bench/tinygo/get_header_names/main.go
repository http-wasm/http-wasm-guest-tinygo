package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

func main() {
	pool.SetHandler(getHeaderNames{})
}

type getHeaderNames struct {
	api.UnimplementedHandler
}

func (getHeaderNames) HandleRequest(req api.Request, _ api.Response) (next bool, reqCtx uint32) {
	_ = req.Headers().Names()
	return // this is a benchmark, so skip the next handler.
}
