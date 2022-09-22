//go:build tinygo.wasm

package httpwasm

//go:wasm-module httpwasm
//go:export log
func log(ptr uintptr, size uint32)
