package main

import (
	httpwasm "github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

var enabledFeatures api.Features

func main() {
	requiredFeatures := api.FeatureBufferRequest | api.FeatureBufferResponse
	enabledFeatures = httpwasm.Host.EnableFeatures(requiredFeatures)

	httpwasm.HandleFn = writeFeatures
}

func writeFeatures(req api.Request, resp api.Response, next api.Next) {
	resp.Body().WriteString(enabledFeatures.String())
}
