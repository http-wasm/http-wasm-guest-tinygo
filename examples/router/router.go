package router

import (
	"strings"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

// Handler implements a simple HTTP router.
type Handler struct {
	api.UnimplementedHandler
}

// HandleRequest implements the same method as documented on api.Handler
func (Handler) HandleRequest(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	// If the URI starts with /host, trim it and dispatch to the next handler.
	if uri := req.GetURI(); strings.HasPrefix(uri, "/host") {
		req.SetURI(uri[5:])
		next = true // proceed to the next handler on the host.
		return
	}

	// Serve a static response
	resp.Headers().Set("Content-Type", "text/plain")
	resp.Body().WriteString("hello")
	return // skip the next handler, as we wrote a response.
}
