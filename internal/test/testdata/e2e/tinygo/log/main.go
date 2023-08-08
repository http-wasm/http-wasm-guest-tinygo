package main

import (
	httpwasm "github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

func main() {
	pool.SetHandler(log{})
}

type log struct{}

func (log) HandleRequest(api.Request, api.Response) (next bool, reqCtx uint32) {
	httpwasm.Log(api.LogLevelInfo, "before")
	next = true
	return
}

func (log) HandleResponse(uint32, api.Request, api.Response, bool) {
	httpwasm.Log(api.LogLevelInfo, "after")
}
