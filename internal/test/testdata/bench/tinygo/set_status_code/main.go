package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleFn = setStatusCode
}

func setStatusCode(req api.Request, resp api.Response, next api.Next) {
	resp.SetStatusCode(404)
}
