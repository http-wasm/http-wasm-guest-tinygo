// Package imports re-exports imported wasm functions to work around an error
// compiling TinyGo to wasm "cannot use an exported function as value".
package imports

import "github.com/http-wasm/http-wasm-guest-tinygo/handler/api"

// BufLimit is the possibly zero maximum length of a result value to write in
// bytes. If the actual value is larger than this, nothing is written to
// memory.
type BufLimit = uint32

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

// EOFLen is the result of FuncReadBody which allows callers to know if the
// bytes returned are the end of the stream. For compatability with WebAssembly
// Core Specification 1.0, two uint32 values are combined into a single uint64
// in the following order:
//
//   - eof: the body is exhausted.
//   - len: possibly zero length of bytes read from the body.
//
// Here's how to split the results:
//
//   - eof: `uint32(eofLen >> 32)`
//   - len: `uint32(eofLen)`
//
// # Examples
//
//   - 1<<32|0 (4294967296): EOF and no bytes were read
//   - 0<<32|16 (16): 16 bytes were read and there may be more available.
//
// Note: `EOF` is not an error, so process `len` bytes returned regardless.
type EOFLen = uint64

type BodyKind uint32

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

type HeaderKind uint32

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

func EnableFeatures(features api.Features) api.Features {
	return enableFeatures(features)
}

func GetConfig(ptr uint32, limit BufLimit) (len uint32) {
	return getConfig(ptr, limit)
}

func LogEnabled(level api.LogLevel) uint32 {
	return logEnabled(level)
}

func Log(level api.LogLevel, ptr, size uint32) {
	log(level, ptr, size)
}

func GetMethod(ptr uint32, limit BufLimit) (len uint32) {
	return getMethod(ptr, limit)
}

func SetMethod(ptr, size uint32) {
	setMethod(ptr, size)
}

func GetURI(ptr uint32, limit BufLimit) (len uint32) {
	return getURI(ptr, limit)
}

func SetURI(ptr, size uint32) {
	setURI(ptr, size)
}

func GetProtocolVersion(ptr uint32, limit BufLimit) (len uint32) {
	return getProtocolVersion(ptr, limit)
}

func GetHeaderNames(kind HeaderKind, ptr uint32, limit BufLimit) CountLen {
	return getHeaderNames(kind, ptr, limit)
}

func GetHeaderValues(kind HeaderKind, namePtr, nameSize uint32, bufPtr uint32, bufLimit BufLimit) CountLen {
	return getHeaderValues(kind, namePtr, nameSize, bufPtr, bufLimit)
}

func SetHeaderValue(kind HeaderKind, namePtr, nameSize uint32, valuePtr uint32, valueSize uint32) {
	setHeaderValue(kind, namePtr, nameSize, valuePtr, valueSize)
}

func AddHeaderValue(kind HeaderKind, namePtr, nameSize uint32, valuePtr uint32, valueSize uint32) {
	addHeaderValue(kind, namePtr, nameSize, valuePtr, valueSize)
}

func RemoveHeader(kind HeaderKind, namePtr, nameSize uint32) {
	removeHeader(kind, namePtr, nameSize)
}

func ReadBody(kind BodyKind, bufPtr uint32, bufLimit BufLimit) EOFLen {
	return readBody(kind, bufPtr, bufLimit)
}

func WriteBody(kind BodyKind, bufPtr uint32, bufSize uint32) {
	writeBody(kind, bufPtr, bufSize)
}

func GetStatusCode() uint32 {
	return getStatusCode()
}

func SetStatusCode(statusCode uint32) {
	setStatusCode(statusCode)
}

func GetSourceAddr(ptr uint32, limit BufLimit) (len uint32) {
	return getSourceAddr(ptr, limit)
}
