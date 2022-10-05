package main

import "github.com/http-wasm/http-wasm-guest-tinygo/handler" //nolint

func main() {
	handler.HandleFn = getRequestHeader
}

func getRequestHeader() {
	_, _ = handler.GetRequestHeader("Accept")
}
