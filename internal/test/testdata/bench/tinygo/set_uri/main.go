package main

import "github.com/http-wasm/http-wasm-guest-tinygo/handler"

func main() {
	handler.HandleFn = setURI
}

func setURI() {
	handler.SetURI("/v1.0/hello")
}
