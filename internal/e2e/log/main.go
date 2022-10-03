//go:build tinygo.wasm

package main

import httpwasm "github.com/http-wasm/http-wasm-guest-tinygo/handler"

func logAround() {
	httpwasm.Log("before")
	defer httpwasm.Log("after")

	httpwasm.Next()
}

func main() {
	httpwasm.HandleFn = logAround
}
