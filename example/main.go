package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleFn = rewrite
}

func rewrite(req api.Request, _ api.Response, next api.Next) {
	if req.GetURI() == "/v1.0/hi?name=panda" {
		req.SetURI("/v1.0/hello?name=teddy")
	}
	next()
}
