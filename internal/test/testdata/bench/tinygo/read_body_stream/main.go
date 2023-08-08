package main

import (
	"io"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

func main() {
	pool.SetHandler(readBodyStream{})
}

type readBodyStream struct {
	api.UnimplementedHandler
}

func (readBodyStream) HandleRequest(req api.Request, _ api.Response) (next bool, reqCtx uint32) {
	if size, err := req.Body().WriteTo(io.Discard); err != nil {
		panic(err)
	} else if size == 0 {
		panic("size == 0")
	}
	return // this is a benchmark, so skip the next handler.
}
