package main

import "github.com/http-wasm/http-wasm-guest-tinygo/handler" //nolint

func main() {
	handler.HandleFn = setPath
}

func setPath() {
	handler.SetPath("/v1.0/hello")
}
