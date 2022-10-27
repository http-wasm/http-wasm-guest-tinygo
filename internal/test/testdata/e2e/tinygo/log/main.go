package main

import (
	httpwasm "github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	httpwasm.HandleRequestFn = logBefore
	httpwasm.HandleResponseFn = logAfter
}

var log = httpwasm.Host.Log

func logBefore(api.Request, api.Response) (next bool, reqCtx uint32) {
	log(api.LogLevelInfo, "before")
	next = true
	return
}

func logAfter(uint32, api.Request, api.Response, bool) {
	log(api.LogLevelInfo, "after")
}
