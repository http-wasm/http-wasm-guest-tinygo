//go:build tinygo.wasm || wasip1

package once

//go:wasmimport http_handler await_response
func awaitResponse(ctxNext uint64) (isError uint32)
