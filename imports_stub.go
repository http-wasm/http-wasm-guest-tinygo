//go:build !tinygo.wasm

package httpwasm

// log is stubbed for compilation outside TinyGo
func log(ptr uintptr, size uint32) {}
