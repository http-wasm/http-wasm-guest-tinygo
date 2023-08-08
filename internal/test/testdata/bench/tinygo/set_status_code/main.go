package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

func main() {
	pool.SetHandler(setStatusCode{})
}

type setStatusCode struct {
	api.UnimplementedHandler
}

func (setStatusCode) HandleRequest(_ api.Request, resp api.Response) (next bool, reqCtx uint32) {
	resp.SetStatusCode(404)
	return // this is a benchmark, so skip the next handler.
}
