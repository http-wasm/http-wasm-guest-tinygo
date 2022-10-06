package main

import "github.com/http-wasm/http-wasm-guest-tinygo/handler" //nolint

func main() {
	handler.HandleFn = next
}

// next is the default, but doing this is more explicit.
func next() {
	handler.Next()
}
