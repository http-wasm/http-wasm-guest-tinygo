//go:build !tinygo.wasm

package internal_test

import (
	"context"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

// instantiateHost instantiates a test waPC host and returns it and a cleanup function.
func instantiateHost(t *testing.T, r wazero.Runtime) (*host, api.Closer) {
	h := &host{t: t}
	// Export host functions (in the order defined in https://github.com/http-wasm/http-wasm-abi)
	if host, err := r.NewModuleBuilder("httpwasm").
		ExportFunction("log", h.log,
			"log", "ptr", "size").
		Instantiate(testCtx, r); err != nil {
		t.Errorf("Error instantiating waPC host - %v", err)
		return h, nil
	} else {
		return h, host
	}
}

type host struct {
	t                  *testing.T
	consoleLogMessages []string
}

// log is the WebAssembly function export "log", which logs a string to the console.
func (w *host) log(ctx context.Context, m api.Module, ptr, size uint32) {
	msg := w.requireReadString(ctx, m.Memory(), "msg", ptr, size)
	w.consoleLogMessages = append(w.consoleLogMessages, msg)
}

// requireReadString is a convenience function that casts requireRead
func (w *host) requireReadString(ctx context.Context, mem api.Memory, fieldName string, offset, byteCount uint32) string {
	return string(w.requireRead(ctx, mem, fieldName, offset, byteCount))
}

// requireRead is like api.Memory except that it panics if the offset and byteCount are out of range.
func (w *host) requireRead(ctx context.Context, mem api.Memory, fieldName string, offset, byteCount uint32) []byte {
	buf, ok := mem.Read(ctx, offset, byteCount)
	if !ok {
		w.t.Fatalf("out of memory reading %s", fieldName)
	}
	return buf
}
