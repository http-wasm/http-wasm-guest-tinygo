package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleFn = readBody
}

var buf = make([]byte, 2048)

func readBody(req api.Request, resp api.Response, next api.Next) {
	if size, eof := req.Body().Read(buf); size == 0 {
		panic("size == 0")
	} else if !eof {
		panic("!eof")
	}
}
