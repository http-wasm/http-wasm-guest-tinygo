//go:build tinygo.wasm

package imports

//go:wasm-module http-handler
//go:export enable_features
func enableFeatures(features uint64) uint64

//go:wasm-module http-handler
//go:export get_config
func getConfig(ptr uintptr, limit uint32) (len uint32)

//go:wasm-module http-handler
//go:export log
func log(ptr uintptr, size uint32)

//go:wasm-module http-handler
//go:export get_method
func getMethod(ptr uintptr, limit uint32) (len uint32)

//go:wasm-module http-handler
//go:export set_method
func setMethod(ptr uintptr, size uint32)

//go:wasm-module http-handler
//go:export get_uri
func getURI(ptr uintptr, limit uint32) (len uint32)

//go:wasm-module http-handler
//go:export set_uri
func setURI(ptr uintptr, size uint32)

//go:wasm-module http-handler
//go:export get_protocol_version
func getProtocolVersion(ptr uintptr, limit uint32) (len uint32)

//go:wasm-module http-handler
//go:export set_protocol_version
func setProtocolVersion(ptr uintptr, size uint32)

//go:wasm-module http-handler
//go:export get_request_header_names
func getRequestHeaderNames(ptr uintptr, limit uint32) (len uint32)

//go:wasm-module http-handler
//go:export get_request_header
func getRequestHeader(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (okLen uint64)

//go:wasm-module http-handler
//go:export set_request_header
func setRequestHeader(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueLen uint32)

//go:wasm-module http-handler
//go:export read_request_body
func readRequestBody(bufPtr uintptr, bufLimit uint32) (eofLen uint64)

//go:wasm-module http-handler
//go:export write_response_body
func writeRequestBody(bufPtr uintptr, bufLen uint32)

//go:wasm-module http-handler
//go:export get_request_trailer_names
func getRequestTrailerNames(ptr uintptr, limit uint32) (len uint32)

//go:wasm-module http-handler
//go:export get_request_trailer
func getRequestTrailer(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (okLen uint64)

//go:wasm-module http-handler
//go:export set_request_trailer
func setRequestTrailer(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueLen uint32)

//go:wasm-module http-handler
//go:export next
func next()

//go:wasm-module http-handler
//go:export get_status_code
func getStatusCode() uint32

//go:wasm-module http-handler
//go:export set_status_code
func setStatusCode(statusCode uint32)

//go:wasm-module http-handler
//go:export get_response_header_names
func getResponseHeaderNames(ptr uintptr, limit uint32) (len uint32)

//go:wasm-module http-handler
//go:export get_response_header
func getResponseHeader(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (okLen uint64)

//go:wasm-module http-handler
//go:export set_response_header
func setResponseHeader(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueLen uint32)

//go:wasm-module http-handler
//go:export read_response_body
func readResponseBody(bufPtr uintptr, bufLimit uint32) (eofLen uint64)

//go:wasm-module http-handler
//go:export write_response_body
func writeResponseBody(bufPtr uintptr, bufLen uint32)

//go:wasm-module http-handler
//go:export get_response_trailer_names
func getResponseTrailerNames(ptr uintptr, limit uint32) (len uint32)

//go:wasm-module http-handler
//go:export get_response_trailer
func getResponseTrailer(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (okLen uint64)

//go:wasm-module http-handler
//go:export set_response_trailer
func setResponseTrailer(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueLen uint32)
