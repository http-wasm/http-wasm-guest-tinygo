//go:build tinygo.wasm

package main

import httpwasm "github.com/http-wasm/http-wasm-guest-tinygo"

func main() {
	httpwasm.Log("msg")
	httpwasm.Log("msg1")
	httpwasm.Log("msg")
}
