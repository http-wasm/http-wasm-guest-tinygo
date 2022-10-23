//go:build tinygo.wasm

package imports

import "github.com/http-wasm/http-wasm-guest-tinygo/handler/api"

//go:wasm-module http-handler
//go:export enable_features
func enableFeatures(features api.Features) api.Features

//go:wasm-module http-handler
//go:export get_config
func getConfig(ptr uintptr, limit BufLimit) (len uint32)

//go:wasm-module http-handler
//go:export log
func log(level api.LogLevel, ptr uintptr, size uint32)

//go:wasm-module http-handler
//go:export log_enabled
func logEnabled(level api.LogLevel) uint32

//go:wasm-module http-handler
//go:export get_method
func getMethod(ptr uintptr, limit BufLimit) (len uint32)

//go:wasm-module http-handler
//go:export set_method
func setMethod(ptr uintptr, size uint32)

//go:wasm-module http-handler
//go:export get_uri
func getURI(ptr uintptr, limit BufLimit) (len uint32)

//go:wasm-module http-handler
//go:export set_uri
func setURI(ptr uintptr, size uint32)

//go:wasm-module http-handler
//go:export get_protocol_version
func getProtocolVersion(ptr uintptr, limit BufLimit) (len uint32)

//go:wasm-module http-handler
//go:export get_header_names
func getHeaderNames(kind HeaderKind, ptr uintptr, limit BufLimit) (countLen CountLen)

//go:wasm-module http-handler
//go:export get_header_values
func getHeaderValues(kind HeaderKind, namePtr uintptr, nameSize uint32, bufPtr uintptr, buflimit BufLimit) (countLen CountLen)

//go:wasm-module http-handler
//go:export set_header_value
func setHeaderValue(kind HeaderKind, namePtr uintptr, nameSize uint32, valuePtr uintptr, valueLen uint32)

//go:wasm-module http-handler
//go:export add_header_value
func addHeaderValue(kind HeaderKind, namePtr uintptr, nameSize uint32, valuePtr uintptr, valueLen uint32)

//go:wasm-module http-handler
//go:export remove_header
func removeHeader(kind HeaderKind, namePtr uintptr, nameSize uint32)

//go:wasm-module http-handler
//go:export read_body
func readBody(kind BodyKind, bufPtr uintptr, buflimit BufLimit) (eofLen EOFLen)

//go:wasm-module http-handler
//go:export write_body
func writeBody(kind BodyKind, bufPtr uintptr, bufLen uint32)

//go:wasm-module http-handler
//go:export next
func next()

//go:wasm-module http-handler
//go:export get_status_code
func getStatusCode() uint32

//go:wasm-module http-handler
//go:export set_status_code
func setStatusCode(statusCode uint32)
