//go:build tinygo.wasm

package main

import http_wasm "github.com/http-wasm/http-wasm-guest-tinygo"

func main() {
	http_wasm.Log("msg")
	http_wasm.Log("msg1")
	http_wasm.Log("msg")
}
