//go:build tinygo.wasm

package http_wasm

//go:wasm-module http_wasm
//go:export log
func log(ptr uintptr, size uint32)
