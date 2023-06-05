//go:build !tinygo.wasm

package e2e_test

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tetratelabs/wazero"

	"github.com/http-wasm/http-wasm-guest-tinygo/internal/test"
	"github.com/http-wasm/http-wasm-host-go/api"
	"github.com/http-wasm/http-wasm-host-go/handler"
	nethttp "github.com/http-wasm/http-wasm-host-go/handler/nethttp"
)

// testCtx is an arbitrary, non-default context. Non-nil also prevents linter errors.
var testCtx = context.WithValue(context.Background(), struct{}{}, "arbitrary")

// compile-time check to ensure recordingLogger implements api.Logger.
var _ api.Logger = &recordingLogger{}

type recordingLogger struct {
	logMessages []string
}

func (r *recordingLogger) IsEnabled(level api.LogLevel) bool {
	return level == api.LogLevelInfo
}

func (r *recordingLogger) Log(ctx context.Context, level api.LogLevel, message string) {
	if level == api.LogLevelInfo {
		r.logMessages = append(r.logMessages, message)
	}
}

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
			name:    "example router guest response",
			bin:     test.BinExampleRouter,
			request: get,
			next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				t.Fatal("host should not see this request")
			}),
			test: func(t *testing.T, content []byte, logMessages []string, stdout, stderr string) {
				// should see the response written by the guest.
				require.Equal(t, "hello", string(content))
				require.Empty(t, stderr)
				require.Empty(t, stdout)
				require.Empty(t, logMessages)
			},
		},
		{
			name: "example router host response",
			bin:  test.BinExampleRouter,
			request: func(url string) (req *http.Request) {
				req, _ = http.NewRequest(http.MethodGet, url+"/host", nil)
				return
			},
			next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Ensure the handler saw the re-written path.
				require.Equal(t, "/", r.URL.Path)
			}),
			test: func(t *testing.T, content []byte, logMessages []string, stdout, stderr string) {
				require.Empty(t, content)
				require.Empty(t, stderr)
				require.Empty(t, stdout)
				require.Empty(t, logMessages)
			},
		},
		{
			name:    "example wasi tinygo",
			bin:     test.BinExampleWASI,
			request: test.RequestExampleWASI,
			next:    test.HandlerExampleWASI,
			test:    testConsole,
		},
		{
			name:    "example wasi wat", // makes sure the implementations match!
			bin:     test.BinExampleWASIWat,
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
				require.Equal(t, "buffer_request|buffer_response", string(content))
				require.Empty(t, stderr)
				require.Empty(t, stdout)
				require.Empty(t, logMessages)
			},
		},
		{
			name: "handle_response",
			bin:  test.BinHandleResponse,
			request: func(url string) (req *http.Request) {
				req, _ = http.NewRequest(http.MethodGet, url, nil)
				return
			},
			next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
			}),
			test: func(t *testing.T, content []byte, logMessages []string, stdout, stderr string) {
				require.Equal(t, "43", string(content))
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
				require.Empty(t, content)
				require.Empty(t, stderr)
				require.Empty(t, stdout)
				require.Equal(t, []string{"before", "after"}, logMessages)
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			var stdoutBuf, stderrBuf bytes.Buffer
			moduleConfig := wazero.NewModuleConfig().WithStdout(&stdoutBuf).WithStderr(&stderrBuf)

			logger := recordingLogger{}

			// Configure and compile the WebAssembly guest binary.
			mw, err := nethttp.NewMiddleware(testCtx, tc.bin,
				handler.Logger(&logger), handler.ModuleConfig(moduleConfig))
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
			tc.test(t, content, logger.logMessages, stdoutBuf.String(), stderrBuf.String())
		})
	}
}

func testConsole(t *testing.T, content []byte, logMessages []string, stdout, stderr string) {
	// First, verify the content, so we know there are no errors.
	require.Equal(t, `{"hello": "world"}`, string(content))
	require.Empty(t, stderr)
	require.Equal(t, `POST /v1.0/hi?name=panda HTTP/1.1
accept-encoding: gzip
content-length: 18
content-type: application/json
host: localhost
user-agent: Go-http-client/1.1

{"hello": "panda"}

HTTP/1.1 200
content-type: application/json
set-cookie: a=b
set-cookie: c=d
trailer: grpc-status
transfer-encoding: chunked

{"hello": "world"}
grpc-status: 1
`, stdout)
	require.Empty(t, logMessages)
}
