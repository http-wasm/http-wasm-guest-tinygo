//go:build !tinygo.wasm

package internal_test

import (
	"context"
	_ "embed"
	"log"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

// testCtx is an arbitrary, non-default context. Non-nil also prevents linter errors.
var testCtx = context.WithValue(context.Background(), struct{}{}, "arbitrary")

var guestWasm map[string][]byte

const (
	guestWasmLog = "log"
)

// TestMain ensures we can read the test wasm prior to running e2e tests.
func TestMain(m *testing.M) {
	wasms := []string{guestWasmLog}
	guestWasm = make(map[string][]byte, len(wasms))
	for _, name := range wasms {
		if wasm, err := os.ReadFile(path.Join("e2e", name, "main.wasm")); err != nil {
			log.Panicln(err)
		} else {
			guestWasm[name] = wasm
		}
	}
	os.Exit(m.Run())
}

func testLog(t *testing.T, guest api.Module, host *host) {
	// main invokes Log
	require.Equal(t, []string{"msg", "msg1", "msg"}, host.consoleLogMessages)
}

func Test_EndToEnd(t *testing.T) {
	type testCase struct {
		name  string
		guest []byte
		test  func(t *testing.T, guest api.Module, host *host)
	}

	tests := []testCase{
		{
			name:  "Log",
			guest: guestWasm[guestWasmLog],
			test:  testLog,
		},
	}

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntimeWithConfig(testCtx, wazero.NewRuntimeConfig().
		// WebAssembly 2.0 allows use of any version of TinyGo, including 0.24+.
		WithWasmCore2())
	defer r.Close(testCtx) // This closes everything this Runtime created.

	// Instantiate WASI, which implements system I/O such as console output and
	// is required for `tinygo build -target=wasi`
	if _, err := wasi_snapshot_preview1.Instantiate(testCtx, r); err != nil {
		t.Errorf("Error instantiating WASI - %v", err)
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			h, cleanup := instantiateHost(t, r)
			defer cleanup.Close(testCtx)

			g, err := r.InstantiateModuleFromBinary(testCtx, tc.guest)
			if err != nil {
				t.Errorf("Error instantiating waPC guest - %v", err)
			}
			defer g.Close(testCtx)

			tc.test(t, g, h)
		})
	}
}
