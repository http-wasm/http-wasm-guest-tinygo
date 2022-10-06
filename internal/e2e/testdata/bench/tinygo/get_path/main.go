package main

import "github.com/http-wasm/http-wasm-guest-tinygo/handler" //nolint

func main() {
	handler.HandleFn = getPath
}

func getPath() {
	_ = handler.GetPath()
}
