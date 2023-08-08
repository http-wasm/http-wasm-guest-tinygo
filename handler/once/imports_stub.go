//go:build !(tinygo.wasm || wasip1)

package once

// awaitResponse is stubbed for compilation outside TinyGo.
func awaitResponse(ctxNext uint64) (isError uint32) { return }
