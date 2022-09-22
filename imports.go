//go:build tinygo.wasm

package httpwasm

//go:wasm-module http
//go:export log
func log(ptr uintptr, size uint32)
