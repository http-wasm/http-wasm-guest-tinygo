package main

import (
	"strconv"

	httpwasm "github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	httpwasm.Host.EnableFeatures(api.FeatureBufferResponse)
	httpwasm.HandleRequestFn = handleRequest
	httpwasm.HandleResponseFn = handleResponse
}

const magic = uint32(43)

func handleRequest(api.Request, api.Response) (next bool, reqCtx uint32) {
	return true, magic
}

func handleResponse(reqCtx uint32, _ api.Request, resp api.Response, _ bool) {
	resp.Body().WriteString(strconv.Itoa(int(reqCtx)))
}
