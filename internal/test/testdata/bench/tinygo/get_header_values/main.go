package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleFn = getHeaderValues
}

func getHeaderValues(req api.Request, resp api.Response, next api.Next) {
	_ = req.Headers().GetAll("Accept")
}
