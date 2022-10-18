package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleFn = log
}

func log(req api.Request, resp api.Response, next api.Next) {
	handler.Host.Log("hello world")
}
