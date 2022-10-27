package main

import (
	"io"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleRequestFn = readBodyStream
}

func readBodyStream(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	if size, err := req.Body().WriteTo(io.Discard); err != nil {
		panic(err)
	} else if size == 0 {
		panic("size == 0")
	}
	return // this is a benchmark, so skip the next handler.
}
