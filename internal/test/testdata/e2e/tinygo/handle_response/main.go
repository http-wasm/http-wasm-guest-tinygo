package main

import (
	"strconv"

	httpwasm "github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

func main() {
	httpwasm.Host.EnableFeatures(api.FeatureBufferResponse)
	pool.SetHandler(handler{})
}

const magic = uint32(43)

type handler struct{}

func (handler) HandleRequest(api.Request, api.Response) (next bool, reqCtx uint32) {
	return true, magic
}

func (handler) HandleResponse(reqCtx uint32, _ api.Request, resp api.Response, _ bool) {
	resp.Body().WriteString(strconv.Itoa(int(reqCtx)))
}
