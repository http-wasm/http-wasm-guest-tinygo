package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

func main() {
	pool.SetHandler(addHeader{})
}

type addHeader struct {
	api.UnimplementedHandler
}

func (addHeader) HandleRequest(_ api.Request, resp api.Response) (next bool, reqCtx uint32) {
	resp.Headers().Add("Set-Cookie", "a=b")
	return // this is a benchmark, so skip the next handler.
}
