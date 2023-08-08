package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

func main() {
	pool.SetHandler(log{})
}

type log struct {
	api.UnimplementedHandler
}

func (log) HandleRequest(api.Request, api.Response) (next bool, reqCtx uint32) {
	handler.Log(api.LogLevelInfo, "hello world")
	return // this is a benchmark, so skip the next handler.
}
