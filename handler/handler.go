package handler

import (
	"unsafe"

	"github.com/tetratelabs/tinymem"
)

// HandleFn is the entry-point function which defaults to Next.
var HandleFn func() = Next

// handle is only exported to the host.
//
//go:export handle
func handle() { //nolint
	HandleFn()
}

// Log logs a message to the host's logs.
func Log(message string) {
	if len(message) == 0 {
		return // don't incur host call overhead
	}
	ptr, size := tinymem.StringToPtr(message)
	log(ptr, size)
}

// GetURI returns the request path.
func GetURI() string {
	size := getURI(0, 0)
	if size == 0 {
		return ""
	}
	buf := make([]byte, size)
	ptr := uintptr(unsafe.Pointer(&buf[0]))
	_ = getURI(ptr, size)
	return string(buf)
}

// SetURI overwrites the request path with the current value.
func SetURI(path string) {
	ptr, size := tinymem.StringToPtr(path)
	setURI(ptr, size)
}

// GetRequestHeader returns the value of the given request header name or false
// if it doesn't exist.
func GetRequestHeader(name string) (string, bool) {
	namePtr, nameSize := tinymem.StringToPtr(name)
	okSize := getRequestHeader(namePtr, nameSize, 0, 0)
	if okSize == 0 {
		return "", false
	}
	size := uint32(okSize)
	buf := make([]byte, size)
	ptr := uintptr(unsafe.Pointer(&buf[0]))
	_ = getRequestHeader(namePtr, nameSize, ptr, size)
	return string(buf), true
}

// Next is an alternative to SendResponse that dispatches control to the next
// handler on the host.
func Next() {
	next()
}

// SetResponseHeader sets a response header of the given name to the value.
func SetResponseHeader(name, value string) {
	namePtr, nameSize := tinymem.StringToPtr(name)
	valuePtr, valueSize := tinymem.StringToPtr(value)
	setResponseHeader(namePtr, nameSize, valuePtr, valueSize)
}

// SendResponse is an alternative to Next that sends the HTTP response with a
// given status code and optional body.
func SendResponse(statusCode uint32, body []byte) {
	bodyPtr := uintptr(unsafe.Pointer(&body[0])) // TODO: tinymem.SliceToPtr
	bodySize := uint32(len(body))
	setStatusCode(statusCode)
	setResponseBody(bodyPtr, bodySize)
}
