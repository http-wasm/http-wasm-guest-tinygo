package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

func main() {
	pool.SetHandler(getHeaderValues{})
}

type getHeaderValues struct {
	api.UnimplementedHandler
}

func (getHeaderValues) HandleRequest(req api.Request, _ api.Response) (next bool, reqCtx uint32) {
	_ = req.Headers().GetAll("Accept")
	return // this is a benchmark, so skip the next handler.
}
