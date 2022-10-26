package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleRequestFn = log
}

func log(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	handler.Host.Log(api.LogLevelInfo, "hello world")
	return // this is a benchmark, so skip the next handler.
}
