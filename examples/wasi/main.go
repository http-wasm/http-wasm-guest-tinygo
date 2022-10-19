package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	httpwasm "github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

// main ensures buffering is available on the host.
//
// Note: required features does not include api.FeatureTrailers because some
// hosts don't support them, and the impact is minimal for logging.
func main() {
	requiredFeatures := api.FeatureBufferRequest | api.FeatureBufferResponse
	if want, have := requiredFeatures, httpwasm.Host.EnableFeatures(requiredFeatures); !have.IsEnabled(want) {
		panic(fmt.Sprint("unexpected features, want: ", want.String(), ", have: ", have.String()))
	}
	httpwasm.HandleFn = handle
}

// handle prints HTTP requests and responses to the console using os.Stdout.
//
// Note: Internally, TinyGo uses WASI to implement os.Stdout. For example,
// writing is a call to the imported function `fd_write`.
func handle(req api.Request, resp api.Response, next api.Next) {
	// Print the incoming request to the console.
	printRequestLine(req)
	printHeaders(req.Headers())
	printBody(req.Body())
	printHeaders(req.Trailers())

	// Handle the request, in whichever way defined by the host.
	next()

	fmt.Println()

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
	fmt.Println(req.GetMethod(), req.GetURI(), req.GetProtocolVersion())
}

// printHeaders prints each header field to the wasi. Ex "a: b"
func printHeaders(h api.Header) {
	for _, n := range h.Names() {
		for _, v := range h.GetAll(n) {
			fmt.Println(n + ": " + v)
		}
	}
}

// compile-time check to ensure bodyWriter implements io.Writer.
var _ io.Writer = (*bodyWriter)(nil)

type bodyWriter bool

// Write adds an extra newline prior to printing it to the wasi.
func (b *bodyWriter) Write(p []byte) (n int, err error) {
	if !*b {
		fmt.Println()
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
		fmt.Println() // add another newline for visibility
	}
}

// printResponseLine prints the response line to the wasi, without the
// status reason. Ex "HTTP/1.1 200"
func printResponseLine(req api.Request, resp api.Response) {
	fmt.Println(req.GetProtocolVersion(), strconv.Itoa(int(resp.GetStatusCode())))
}
