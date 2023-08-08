package api

import "io"

// LogLevel controls the volume of logging. The lower the number the more
// detail is logged.
//
// Note: The most voluminous level, LogLevelDebug is -1 to prevent users from
// accidentally defaulting to it.
type LogLevel int32

const (
	LogLevelDebug LogLevel = -1
	LogLevelInfo  LogLevel = 0
	LogLevelWarn  LogLevel = 1
	LogLevelError LogLevel = 2
	LogLevelNone  LogLevel = 3
)

// Host is the WebAssembly host that accepts server requests. For example,
// if written in Go, the host controls the http.Handler which dispatches to
// the Handler here which would be compiled to wasm.
type Host interface {
	// EnableFeatures tries to enable the given features and returns the
	// Features bitflag supported by the host. This must be called prior to
	// Next to enable Features needed by the guest.
	//
	// This may be called prior to the Handler, for example inside the main
	// function. Doing so reduces overhead per-call and also allows the guest
	// to fail early on unsupported.
	//
	// If called during Handler, any new features are only enabled for the
	// scope of the current request. This allows fine-grained access to
	// expensive features such as FeatureBufferResponse.
	EnableFeatures(Features) Features

	// GetConfig reads any configuration set by the host.
	GetConfig() []byte

	// LogEnabled returns true if the LogLevel is enabled. This value may
	// be cached at request granularity.
	//
	// This function is used to avoid unnecessary overhead generating log
	// messages that the host would discard due to its level being below this.
	LogEnabled(LogLevel) bool

	// Log logs a message to the host's logs at the given LogLevel.
	Log(LogLevel, string)
}

// Handler is what implements the HTTP request/response lifecycle.
type Handler interface {
	// HandleRequest is the entrypoint the Host calls when processing an HTTP
	// request. Default is to proceed to the next handler on the host.
	//
	// Implementations that construct a response locally should set `next=false`.
	// Otherwise, set `next=true` to proceed to the next handler on the host.
	// Handlers that require request correlation can optionally set reqCtx.
	// The host will propagate this value as the first parameter to HandleResponse.
	//
	// Here are some examples:
	//
	// Modify the incoming request:
	//
	//	type router struct {
	//		api.UnimplementedHandler
	//	}
	//
	//	func (router) HandleRequest(req api.Request, _ api.Response) (next bool, reqCtx uint32) {
	//		if req.GetURI() == "/v1.0/hi?name=panda" {
	//			req.SetURI("/v1.0/hello?name=teddy")
	//		}
	//		next = true
	//		return
	//	}
	//
	// Ex. To serve a response locally:
	//
	//	type hello struct {
	//		api.UnimplementedHandler
	//	}
	//
	//	func (hello) HandleRequest(_ api.Request, resp api.Response) (next bool, reqCtx uint32) {
	//		resp.Body().WriteString("hello world")
	//	}
	//
	// Return an empty 200 response
	//
	//	type noop struct {
	//		api.UnimplementedHandler
	//	}
	//
	//	func (noop) HandleRequest(api.Request, api.Response) (next bool, reqCtx uint32) { }
	//
	// # Request Context
	//
	// Implementations who return `next=true` may also set request correlation
	// data as reqCtx. The host will propagate this to HandleResponse.
	HandleRequest(Request, Response) (next bool, reqCtx uint32)

	// HandleResponse is invoked when HandleRequest returned `next=true`. Its
	// possibly zero `reqCtx` result is propagated here.
	//
	// `isError=true` when the host erred since returning from HandleRequest.
	HandleResponse(reqCtx uint32, req Request, resp Response, isError bool)
}

// UnimplementedHandler calls the next handler.
type UnimplementedHandler struct{}

func (UnimplementedHandler) HandleRequest(Request, Response) (next bool, reqCtx uint32) {
	next = true
	return
}

func (UnimplementedHandler) HandleResponse(uint32, Request, Response, bool) {}

// Request is the incoming HTTP request sent by the client or an upstream
// handler.
type Request interface {
	// GetMethod returns the method. Ex. "GET"
	GetMethod() string

	// SetMethod overwrites the method.
	SetMethod(method string)

	// GetURI returns the request URI. Ex. "/v1.0/hi?name=panda"
	//
	// This is analogous to the URL property of http.Request when used server-side.
	// Those who wish to use url.URL can parse this with url.ParseRequestURI.
	//
	// Note: The URI may include query parameters.
	GetURI() string

	// SetURI overwrites the URI.
	//
	// This is analogous to setting the URL property of http.Request when used server-side.
	// Those who wish to use url.URL can generate it with url.URL.String.
	//
	// Note: The URI may include query parameters.
	SetURI(uri string)

	// GetProtocolVersion returns the HTTP protocol version. Ex. "HTTP/1.1"
	GetProtocolVersion() string

	// Headers allows access to any incoming request headers.
	Headers() Header

	// Body allows access to any incoming request body. To read this without
	// preventing the Next from reading it, enable FeatureBufferRequest.
	Body() Body

	// Trailers allows access to any incoming request trailing headers
	// (trailers).
	Trailers() Header
}

// Response is any outgoing HTTP response. If not generated by Next, this
// defaults to an empty HTTP 200 response.
type Response interface {
	// GetStatusCode returns the HTTP status code. Ex. 200
	GetStatusCode() uint32

	// SetStatusCode overwrites the HTTP status code.
	SetStatusCode(statusCode uint32)

	// Headers allows access to any outgoing request headers.
	Headers() Header

	// Body allows access to any outgoing response body. Access to this after
	// Next requires FeatureBufferResponse.
	Body() Body

	// Trailers allows access to any outgoing request trailing headers
	// (trailers).
	Trailers() Header
}

// Header is a key-value field in an HTTP message. This type is used regardless
// of whether the headers are trailing (trailers) or not.
type Header interface {
	// Names returns all field names, in a possibly arbitrary order determined
	// by the host.
	Names() []string

	// Get returns one value for the given name, or false if there are none.
	Get(name string) (value string, ok bool)

	// GetAll returns all values for the given name, or nil if there are none.
	GetAll(name string) []string

	// Set overwrites any header values for the given name.
	Set(name, value string)

	// Add adds a header value for the given name.
	Add(name, value string)

	// Remove removes all values for the given name.
	Remove(name string)
}

// Body is the HTTP message body.
type Body interface {
	// WriteTo writes all data in the body to the writer and returns the length
	// in bytes or an error if the writer raises one.
	WriteTo(io.Writer) (size uint64, err error)

	// Read reads the body into the buffer and returns the length in bytes read
	// and true if the stream is empty as a result.
	Read([]byte) (size uint32, eof bool)

	// Write adds data to the current response body.
	Write([]byte)

	// WriteString is similar to Write, except for strings.
	WriteString(string)
}
