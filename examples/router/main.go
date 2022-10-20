package main

import (
	"strings"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleFn = handle
}

// handle implements a simple HTTP router.
func handle(req api.Request, resp api.Response, next api.Next) {
	// If the URI starts with /host, trim it and dispatch to the next handler.
	if uri := req.GetURI(); strings.HasPrefix(uri, "/host") {
		req.SetURI(uri[5:])
		next()
	} else { // Serve a static response
		resp.Headers().Set("Content-Type", "text/plain")
		resp.Body().WriteString("hello")
	}
}
