package main

import (
	"net/url"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleRequestFn = setURI
}

func setURI(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	u, _ := url.ParseRequestURI("/v1.0/hello")
	req.SetURI(*u)
	return // this is a benchmark, so skip the next handler.
}
