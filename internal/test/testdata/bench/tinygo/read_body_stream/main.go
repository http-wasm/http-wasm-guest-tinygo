package main

import (
	"io"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleFn = readBodyStream
}

func readBodyStream(req api.Request, resp api.Response, next api.Next) {
	if size, err := req.Body().WriteTo(io.Discard); err != nil {
		panic(err)
	} else if size == 0 {
		panic("size == 0")
	}
}
