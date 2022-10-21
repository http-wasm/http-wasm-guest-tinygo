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
	log(api.LogLevelInfo, "before")
	defer log(api.LogLevelInfo, "after")

	next()
}
