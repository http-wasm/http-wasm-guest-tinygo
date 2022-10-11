//go:build tinygo.wasm

package handler

//go:wasm-module http-handler
//go:export log
func log(ptr uintptr, size uint32)

//go:wasm-module http-handler
//go:export get_uri
func getURI(ptr uintptr, limit uint32) uint32

//go:wasm-module http-handler
//go:export set_uri
func setURI(ptr uintptr, size uint32)

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
//go:export set_status_code
func setStatusCode(statusCode uint32)

//go:wasm-module http-handler
//go:export set_response_body
func setResponseBody(bodyPtr uintptr, bodySize uint32)
