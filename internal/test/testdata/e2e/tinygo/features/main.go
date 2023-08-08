package main

import (
	httpwasm "github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/pool"
)

var enabledFeatures api.Features

func main() {
	requiredFeatures := api.FeatureBufferRequest | api.FeatureBufferResponse
	enabledFeatures = httpwasm.Host.EnableFeatures(requiredFeatures)

	pool.SetHandler(writeFeatures{})
}

type writeFeatures struct {
	api.UnimplementedHandler
}

func (writeFeatures) HandleRequest(_ api.Request, resp api.Response) (next bool, reqCtx uint32) {
	resp.Body().WriteString(enabledFeatures.String())
	return
}
