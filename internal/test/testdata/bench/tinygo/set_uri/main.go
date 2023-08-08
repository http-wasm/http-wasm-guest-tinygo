package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

func main() {
	pool.SetHandler(setURI{})
}

type setURI struct {
	api.UnimplementedHandler
}

func (setURI) HandleRequest(req api.Request, _ api.Response) (next bool, reqCtx uint32) {
	req.SetURI("/v1.0/hello")
	return // this is a benchmark, so skip the next handler.
}
