package main

import (
	httpwasm "github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

var enabledFeatures api.Features

func main() {
	requiredFeatures := api.FeatureBufferRequest | api.FeatureBufferResponse
	enabledFeatures = httpwasm.Host.EnableFeatures(requiredFeatures)

	httpwasm.HandleRequestFn = writeFeatures
}

func writeFeatures(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	resp.Body().WriteString(enabledFeatures.String())
	return
}
