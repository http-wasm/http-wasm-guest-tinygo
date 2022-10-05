package main

import "github.com/http-wasm/http-wasm-guest-tinygo/handler" //nolint

func main() {
	handler.HandleFn = log
}

func log() {
	handler.Log("hello world")
}
