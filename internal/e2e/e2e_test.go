//go:build !tinygo.wasm

package e2e_test

import (
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
)

// testCtx is an arbitrary, non-default context. Non-nil also prevents linter errors.
var testCtx = context.WithValue(context.Background(), struct{}{}, "arbitrary")

func Test_EndToEnd(t *testing.T) {
	type testCase struct {
		name    string
		bin     []byte
		request func(url string) (*http.Request, error)
		next    http.Handler
		test    func(t *testing.T, content []byte, logMessages []string)
	}

	tests := []testCase{
		{
			name: "example",
			bin:  BinExample,
			request: func(url string) (*http.Request, error) {
				url = fmt.Sprintf("%s/v1.0/hi?name=panda", url)
				return http.NewRequest(http.MethodGet, url, nil)
			},
			next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r.Header.Set("Content-Type", "text/plain")
				w.Write([]byte(r.URL.String())) // nolint
			}),
			test: func(t *testing.T, content []byte, logMessages []string) {
				// Ensure the handler saw the re-written path.
				require.Equal(t, "/v1.0/hello?name=teddy", string(content))
			},
		},
		{
			name: "log",
			bin:  BinLog,
			request: func(url string) (*http.Request, error) {
				return http.NewRequest(http.MethodGet, url, nil)
			},
			next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
			}),
			test: func(t *testing.T, content []byte, logMessages []string) {
				require.Equal(t, []string{"before", "after"}, logMessages)
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			var logMessages []string
			logger := func(_ context.Context, message string) { logMessages = append(logMessages, message) }

			// Configure and compile the WebAssembly guest binary. In this case, it is
			// a logging interceptor.
			mw, err := nethttp.NewMiddleware(testCtx, tc.bin, httpwasm.Logger(logger))
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
			req, err := tc.request(ts.URL)
			require.NoError(t, err)

			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			content, _ := io.ReadAll(resp.Body)
			tc.test(t, content, logMessages)
		})
	}
}
