// Package imports re-exports imported wasm functions to work around an error
// compiling TinyGo to wasm "cannot use an exported function as value".
package imports

import "github.com/http-wasm/http-wasm-guest-tinygo/handler/api"

// CountLen describes a possible empty sequence of NUL-terminated strings. For
// compatability with WebAssembly Core Specification 1.0, two uint32 values are
// combined into a single uint64 in the following order:
//
//   - count: zero if the sequence is empty, or the count of strings.
//   - len: possibly zero length of the sequence, including NUL-terminators.
//
// If the uint64 result is zero, the sequence is empty. Otherwise, you need to
// split the results like so.
//
//   - count: `uint32(countLen >> 32)`
//   - len: `uint32(countLen)`
//
// # Examples
//
//   - "": 0<<32|0 or simply zero.
//   - "Accept\0": 1<<32|7
//   - "Content-Type\0Content-Length\0": 2<<32|28
type CountLen = uint64

type BodyKind = uint32

const (
	// BodyKindRequest represents an operation on an HTTP request body.
	//
	// # Notes on FuncReadBody
	//
	// FeatureBufferResponse is required to read the request body without
	// consuming it. To enable it, call FuncEnableFeatures before FuncNext.
	// Otherwise, a downstream handler may panic attempting to read a request
	// body already read upstream.
	//
	// # Notes on FuncWriteBody
	//
	// The first call to FuncWriteBody in FuncHandle overwrites any request
	// body.
	BodyKindRequest BodyKind = 0

	// BodyKindResponse represents an operation on an HTTP request body.
	//
	// # Notes on FuncReadBody
	//
	// FeatureBufferResponse is required to read the response body produced by
	// FuncNext. To enable it, call FuncEnableFeatures beforehand. Otherwise,
	// a handler may panic calling FuncReadBody with BodyKindResponse.
	//
	// # Notes on FuncWriteBody
	//
	// The first call to FuncWriteBody in FuncHandle or after FuncNext
	// overwrites any response body.
	BodyKindResponse BodyKind = 1
)

type HeaderKind = uint32

const (
	// HeaderKindRequest represents an operation on HTTP request headers.
	HeaderKindRequest HeaderKind = 0

	// HeaderKindResponse represents an operation on HTTP response headers.
	HeaderKindResponse HeaderKind = 1

	// HeaderKindRequestTrailers represents an operation on HTTP request
	// trailers (trailing headers). This requires FeatureTrailers.
	//
	// To enable FeatureTrailers, call FuncEnableFeatures prior to FuncNext.
	// Doing otherwise, may result in a panic.
	HeaderKindRequestTrailers HeaderKind = 2

	// HeaderKindResponseTrailers represents an operation on HTTP response
	// trailers (trailing headers). This requires FeatureTrailers.
	//
	// To enable FeatureTrailers, call FuncEnableFeatures prior to FuncNext.
	// Doing otherwise, may result in a panic.
	HeaderKindResponseTrailers HeaderKind = 3
)

func EnableFeatures(features uint64) uint64 {
	return enableFeatures(features)
}

func GetConfig(ptr uintptr, limit uint32) (len uint32) {
	return getConfig(ptr, limit)
}

func Log(level api.LogLevel, ptr uintptr, size uint32) {
	log(level, ptr, size)
}

func GetMethod(ptr uintptr, limit uint32) (len uint32) {
	return getMethod(ptr, limit)
}

func SetMethod(ptr uintptr, size uint32) {
	setMethod(ptr, size)
}

func GetURI(ptr uintptr, limit uint32) (len uint32) {
	return getURI(ptr, limit)
}

func SetURI(ptr uintptr, size uint32) {
	setURI(ptr, size)
}

func GetProtocolVersion(ptr uintptr, limit uint32) (len uint32) {
	return getProtocolVersion(ptr, limit)
}

func GetHeaderNames(kind HeaderKind, ptr uintptr, limit uint32) (countLen CountLen) {
	return getHeaderNames(kind, ptr, limit)
}

func GetHeaderValues(kind HeaderKind, namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (countLen CountLen) {
	return getHeaderValues(kind, namePtr, nameSize, bufPtr, bufLimit)
}

func SetHeaderValue(kind HeaderKind, namePtr uintptr, nameSize uint32, valuePtr uintptr, valueSize uint32) {
	setHeaderValue(kind, namePtr, nameSize, valuePtr, valueSize)
}

func ReadBody(kind BodyKind, bufPtr uintptr, bufLimit uint32) (eofLen uint64) {
	return readBody(kind, bufPtr, bufLimit)
}

func WriteBody(kind BodyKind, bufPtr uintptr, bufSize uint32) {
	writeBody(kind, bufPtr, bufSize)
}

func Next() {
	next()
}

func GetStatusCode() uint32 {
	return getStatusCode()
}

func SetStatusCode(statusCode uint32) {
	setStatusCode(statusCode)
}
