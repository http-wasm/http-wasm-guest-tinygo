//go:build tinygo.wasm

package handler

//go:wasm-module http-handler
//go:export log
func log(ptr uintptr, size uint32)

//go:wasm-module http-handler
//go:export get_path
func getPath(ptr uintptr, limit uint32) uint32

//go:wasm-module http-handler
//go:export set_path
func setPath(ptr uintptr, size uint32)

//go:wasm-module http-handler
//go:export get_request_header
func getRequestHeader(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueLimit uint32) uint64

//go:wasm-module http-handler
//go:export next
func next()

//go:wasm-module http-handler
//go:export set_response_header
func setResponseHeader(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueSize uint32)

//go:wasm-module http-handler
//go:export send_response
func sendResponse(statusCode uint32, bodyPtr uintptr, bodySize uint32)
