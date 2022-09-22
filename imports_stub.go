//go:build !tinygo.wasm

package http_wasm

// log is stubbed for compilation outside TinyGo
func log(ptr uintptr, size uint32) {}
