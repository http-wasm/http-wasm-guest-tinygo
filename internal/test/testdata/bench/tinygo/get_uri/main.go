package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

func main() {
	pool.SetHandler(getURI{})
}

type getURI struct {
	api.UnimplementedHandler
}

func (getURI) HandleRequest(req api.Request, _ api.Response) (next bool, reqCtx uint32) {
	_ = req.GetURI()
	return // this is a benchmark, so skip the next handler.
}
