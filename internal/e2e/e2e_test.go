//go:build !tinygo.wasm

package e2e_test

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	httpwasm "github.com/http-wasm/http-wasm-host-go"
	nethttp "github.com/http-wasm/http-wasm-host-go/handler/nethttp"
	"github.com/stretchr/testify/require"
	"github.com/tetratelabs/wazero"

	"github.com/http-wasm/http-wasm-guest-tinygo/internal/test"
)

// testCtx is an arbitrary, non-default context. Non-nil also prevents linter errors.
var testCtx = context.WithValue(context.Background(), struct{}{}, "arbitrary")

func Test_EndToEnd(t *testing.T) {
	type testCase struct {
		name    string
		bin     []byte
		request func(url string) *http.Request
		next    http.Handler
		test    func(t *testing.T, content []byte, logMessages []string, stdout, stderr string)
	}

	tests := []testCase{
		{
			name: "example rewrite",
			bin:  test.BinExampleRewrite,
			request: func(url string) (req *http.Request) {
				url = fmt.Sprintf("%s/v1.0/hi?name=panda", url)
				req, _ = http.NewRequest(http.MethodGet, url, nil)
				return
			},
			next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r.Header.Set("Content-Type", "text/plain")
				w.Write([]byte(r.URL.String())) // nolint
			}),
			test: func(t *testing.T, content []byte, logMessages []string, stdout, stderr string) {
				// Ensure the handler saw the re-written path.
				require.Equal(t, "/v1.0/hello?name=teddy", string(content))
				require.Equal(t, "", stderr)
				require.Equal(t, "", stdout)
				require.Nil(t, logMessages)
			},
		},
		{
			name:    "example wasi",
			bin:     test.BinExampleWASI,
			request: test.RequestExampleWASI,
			next:    test.HandlerExampleWASI,
			test:    testConsole,
		},
		{
			name:    "wasi - wat", // makes sure the implementations match!
			bin:     test.BinE2EWASIWat,
			request: test.RequestExampleWASI,
			next:    test.HandlerExampleWASI,
			test:    testConsole,
		},
		{
			name: "features",
			bin:  test.BinE2EFeaturesTinyGo,
			request: func(url string) (req *http.Request) {
				req, _ = http.NewRequest(http.MethodGet, url, nil)
				return
			},
			next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			}),
			test: func(t *testing.T, content []byte, logMessages []string, stdout, stderr string) {
				require.Equal(t, "buffer-request|buffer-response", string(content))
				require.Equal(t, "", stderr)
				require.Equal(t, "", stdout)
				require.Nil(t, logMessages)
			},
		},
		{
			name: "log",
			bin:  test.BinE2ELog,
			request: func(url string) (req *http.Request) {
				req, _ = http.NewRequest(http.MethodGet, url, nil)
				return
			},
			next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
			}),
			test: func(t *testing.T, content []byte, logMessages []string, stdout, stderr string) {
				require.Equal(t, "", string(content))
				require.Equal(t, "", stderr)
				require.Equal(t, "", stdout)
				require.Equal(t, []string{"before", "after"}, logMessages)
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			var stdoutBuf, stderrBuf bytes.Buffer
			moduleConfig := wazero.NewModuleConfig().WithStdout(&stdoutBuf).WithStderr(&stderrBuf)

			var logMessages []string
			logger := func(_ context.Context, message string) { logMessages = append(logMessages, message) }

			// Configure and compile the WebAssembly guest binary.
			mw, err := nethttp.NewMiddleware(testCtx, tc.bin,
				httpwasm.Logger(logger), httpwasm.ModuleConfig(moduleConfig))
			if err != nil {
				t.Error(err)
			}
			defer mw.Close(testCtx)

			// Wrap the test handler with one implemented in WebAssembly.
			wrapped := mw.NewHandler(testCtx, tc.next)
			require.NoError(t, err)

			// Start the server with the wrapped handler.
			ts := httptest.NewServer(wrapped)
			defer ts.Close()

			// Make a client request and invoke the test.
			resp, err := ts.Client().Do(tc.request(ts.URL))
			require.NoError(t, err)
			defer resp.Body.Close()

			content, _ := io.ReadAll(resp.Body)
			tc.test(t, content, logMessages, stdoutBuf.String(), stderrBuf.String())
		})
	}
}

func testConsole(t *testing.T, content []byte, logMessages []string, stdout, stderr string) {
	// First, verify the content, so we know there are no errors.
	require.Equal(t, `{"hello": "world"}`, string(content))
	require.Equal(t, "", stderr)
	require.Equal(t, `POST /v1.0/hi?name=panda HTTP/1.1
Accept-Encoding: gzip
Content-Length: 18
Content-Type: application/json
Host: localhost
User-Agent: Go-http-client/1.1

{"hello": "panda"}

HTTP/1.1 200
Content-Type: application/json
Set-Cookie: a=b
Set-Cookie: c=d
Trailer: grpc-status
Transfer-Encoding: chunked

{"hello": "world"}
grpc-status: 1
`, stdout)
	require.Nil(t, logMessages)
}
