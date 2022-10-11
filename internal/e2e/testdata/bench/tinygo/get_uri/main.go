package main

import "github.com/http-wasm/http-wasm-guest-tinygo/handler" //nolint

func main() {
	handler.HandleFn = getURI
}

func getURI() {
	_ = handler.GetURI()
}
