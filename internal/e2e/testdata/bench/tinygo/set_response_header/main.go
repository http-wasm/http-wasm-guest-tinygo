package main

import "github.com/http-wasm/http-wasm-guest-tinygo/handler" //nolint

func main() {
	handler.HandleFn = setResponseHeader
}

func setResponseHeader() {
	handler.SetResponseHeader("Content-Type", "text/plain")
}
