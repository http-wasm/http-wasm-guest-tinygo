package handler

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/imports"
)

// Host is the current host that invokes HandleFn.
var Host api.Host = wasmHost{}

// HandleFn is the entry-point function which defaults to call api.Next.
var HandleFn api.Handler = func(req api.Request, resp api.Response, next api.Next) {
	next()
}

// handle is only exported to the host.
//
//go:export handle
func handle() { //nolint
	HandleFn(wasmRequest{}, wasmResponse{}, imports.Next)
}
