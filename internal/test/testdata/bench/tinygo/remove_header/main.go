package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

func main() {
	pool.SetHandler(removeHeader{})
}

type removeHeader struct {
	api.UnimplementedHandler
}

func (removeHeader) HandleRequest(_ api.Request, resp api.Response) (next bool, reqCtx uint32) {
	resp.Headers().Remove("Set-Cookie")
	return // this is a benchmark, so skip the next handler.
}
