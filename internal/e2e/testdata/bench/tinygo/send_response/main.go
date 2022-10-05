package main

import "github.com/http-wasm/http-wasm-guest-tinygo/handler" //nolint

func main() {
	handler.HandleFn = sendResponse
}

var body = []byte("hello world")

func sendResponse() {
	handler.SendResponse(200, body)
}
