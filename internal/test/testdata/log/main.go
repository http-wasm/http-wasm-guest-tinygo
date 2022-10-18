package main

import (
	httpwasm "github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	httpwasm.HandleFn = logAround
}

var log = httpwasm.Host.Log

func logAround(req api.Request, resp api.Response, next api.Next) {
	log("before")
	defer log("after")

	next()
}
