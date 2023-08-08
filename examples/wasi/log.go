package log

import (
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

// init ensures buffering is available on the host.
//
// Note: required features does not include api.FeatureTrailers because some
// hosts don't support them, and the impact is minimal for logging.
func init() {
	requiredFeatures := api.FeatureBufferRequest | api.FeatureBufferResponse
	if want, have := requiredFeatures, handler.Host.EnableFeatures(requiredFeatures); !have.IsEnabled(want) {
		panic("unexpected features, want: " + want.String() + ", have: " + have.String())
	}
}

// Handler prints HTTP requests and responses to the console using os.Stdout.
//
// Note: Internally, TinyGo uses WASI to implement os.Stdout. For example,
// writing is a call to the imported function `fd_write`.
type Handler struct{}

// HandleRequest implements the same method as documented on api.Handler
func (Handler) HandleRequest(req api.Request, _ api.Response) (next bool, reqCtx uint32) {
	// Print the incoming request to the console.
	printRequestLine(req)
	printHeaders(req.Headers())
	printBody(req.Body())
	printHeaders(req.Trailers())

	next = true // proceed to the next handler on the host.
	return
}

// HandleResponse implements the same method as documented on api.Handler
func (Handler) HandleResponse(_ uint32, req api.Request, resp api.Response, isError bool) {
	println()

	if isError {
		println("host error")
		return
	}

	// Because we enabled buffering, we can read the response.
	// Print it to the console.
	printResponseLine(req, resp)
	printHeaders(resp.Headers())
	printBody(resp.Body())
	printHeaders(resp.Trailers())
}

// printRequestLine prints the request line to the wasi.
// Ex "GET /a HTTP/1.1"
func printRequestLine(req api.Request) {
	println(req.GetMethod(), req.GetURI(), req.GetProtocolVersion())
}

// printHeaders prints each header field to the wasi. Ex "a: b"
func printHeaders(h api.Header) {
	for _, n := range h.Names() {
		for _, v := range h.GetAll(n) {
			println(n + ": " + v)
		}
	}
}

// compile-time check to ensure bodyWriter implements io.Writer.
var _ io.Writer = (*bodyWriter)(nil)

type bodyWriter bool

// Write adds an extra newline prior to printing it to the wasi.
func (b *bodyWriter) Write(p []byte) (n int, err error) {
	if !*b {
		println()
		*b = true
	}
	return os.Stdout.Write(p)
}

// printBody prints the body to the wasi.
func printBody(b api.Body) {
	var w bodyWriter
	if size, err := b.WriteTo(&w); err != nil {
		panic(err)
	} else if size > 0 {
		println() // add another newline for visibility
	}
}

// printResponseLine prints the response line to the wasi, without the
// status reason. Ex "HTTP/1.1 200"
func printResponseLine(req api.Request, resp api.Response) {
	println(req.GetProtocolVersion(), strconv.Itoa(int(resp.GetStatusCode())))
}

// println is like fmt.Println, but faster and smaller with TinyGo.
func println(s ...string) {
	if len(s) == 0 {
		os.Stdout.WriteString("\n") // nolint
		return
	}
	var b strings.Builder
	b.WriteString(s[0])
	for _, s := range s[1:] {
		b.WriteByte(' ')
		b.WriteString(s)
	}
	b.WriteByte('\n')
	os.Stdout.WriteString(b.String()) // nolint
}
