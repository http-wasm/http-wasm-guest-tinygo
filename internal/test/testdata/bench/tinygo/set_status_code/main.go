package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleRequestFn = setStatusCode
}

func setStatusCode(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	resp.SetStatusCode(404)
	return // this is a benchmark, so skip the next handler.
}
