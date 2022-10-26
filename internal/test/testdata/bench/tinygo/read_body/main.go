package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleRequestFn = readBody
}

var buf = make([]byte, 2048)

func readBody(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	if size, eof := req.Body().Read(buf); size == 0 {
		panic("size == 0")
	} else if !eof {
		panic("!eof")
	}
	return // this is a benchmark, so skip the next handler.
}
