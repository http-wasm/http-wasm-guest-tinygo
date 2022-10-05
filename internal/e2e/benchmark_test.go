package e2e_test

import (
	"context"
	_ "embed"
	"net/http"
	"net/url"
	"testing"

	nethttp "github.com/http-wasm/http-wasm-host-go/handler/nethttp"
)

var (
	readOnlyRequest = &http.Request{
		Method: "GET",
		URL:    &url.URL{Path: "/v1.0/hi"},
		Header: http.Header{},
	}

	readOnlyRequestWithHeader = &http.Request{
		Method: "GET",
		URL:    &url.URL{Path: "/v1.0/hi"},
		Header: http.Header{"Accept": {"text/plain"}},
	}
)

var benches = map[string]struct {
	bins       map[string][]byte
	newRequest func() *http.Request
}{
	"get_path": {
		bins: map[string][]byte{
			"TinyGo": BinBenchGetPathTinyGo,
			"wat":    BinBenchGetPathWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequest
		}},
	"get_request_header exists": {
		bins: map[string][]byte{
			"TinyGo": BinBenchGetRequestHeaderTinyGo,
			"wat":    BinBenchGetRequestHeaderWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequestWithHeader
		}},
	"get_request_header not exists": {
		bins: map[string][]byte{
			"TinyGo": BinBenchGetRequestHeaderTinyGo,
			"wat":    BinBenchGetRequestHeaderWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequest
		}},
	"log": {
		bins: map[string][]byte{
			"TinyGo": BinBenchLogTinyGo,
			"wat":    BinBenchLogWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequest
		}},
	"next": {
		bins: map[string][]byte{
			"TinyGo": BinBenchNextTinyGo,
			"wat":    BinBenchNextWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequest
		}},
	"sendResponse": {
		bins: map[string][]byte{
			"TinyGo": BinBenchSendResponseTinyGo,
			"wat":    BinBenchSendResponseWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequest
		}},
	"setPath": {
		bins: map[string][]byte{
			"TinyGo": BinBenchSetPathTinyGo,
			"wat":    BinBenchSetPathWat,
		},
		newRequest: func() *http.Request {
			return &http.Request{URL: &url.URL{}}
		}},
	"setResponseHeader": {
		bins: map[string][]byte{
			"TinyGo": BinBenchSetResponseHeaderTinyGo,
			"wat":    BinBenchSetResponseHeaderWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequest
		}},
}

func Benchmark(b *testing.B) {
	for n, s := range benches {
		b.Run(n, func(b *testing.B) {
			for n, bin := range s.bins {
				benchmark(b, n, bin, s.newRequest)
			}
		})
	}
}

func benchmark(b *testing.B, name string, bin []byte, newRequest func() *http.Request) {
	ctx := context.Background()

	mw, err := nethttp.NewMiddleware(ctx, bin)
	if err != nil {
		b.Fatal(err)
	}
	defer mw.Close(ctx)

	h := mw.NewHandler(ctx, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))

	b.Run(name, func(b *testing.B) {
		// We don't report allocations because memory allocations for TinyGo are
		// in wasm which isn't visible to the Go benchmark.
		for i := 0; i < b.N; i++ {
			h.ServeHTTP(fakeResponseWriter{}, newRequest())
		}
	})
}

var _ http.ResponseWriter = fakeResponseWriter{}

type fakeResponseWriter struct{}

func (rw fakeResponseWriter) Header() http.Header {
	return http.Header{}
}

func (rw fakeResponseWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

func (rw fakeResponseWriter) WriteHeader(statusCode int) {
	// None of our benchmark tests should send failure status. If there's a
	// failure, it is likely there's a problem in the test data.
	if statusCode != 200 {
		panic(statusCode)
	}
}
