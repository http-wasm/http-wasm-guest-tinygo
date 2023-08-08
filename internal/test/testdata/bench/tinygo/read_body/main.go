package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

func main() {
	pool.SetHandler(readBody{})
}

type readBody struct {
	api.UnimplementedHandler
}

var buf = make([]byte, 2048)

func (readBody) HandleRequest(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	if size, eof := req.Body().Read(buf); size == 0 {
		panic("size == 0")
	} else if !eof {
		panic("!eof")
	}
	return // this is a benchmark, so skip the next handler.
}
