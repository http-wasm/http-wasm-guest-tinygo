package e2e_test

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"net/http"
	"net/url"
	"testing"

	nethttp "github.com/http-wasm/http-wasm-host-go/handler/nethttp"

	"github.com/http-wasm/http-wasm-guest-tinygo/internal/test"
)

var (
	readOnlyRequest                *http.Request
	readOnlyRequestWithHeader      *http.Request
	readOnlyRequestWithLargeHeader *http.Request
	smallBody                      []byte
	largeSize                      int
	largeBody                      []byte
)

func init() {
	smallBody = []byte("hello world")
	largeSize = 4096 // 2x the read buffer size
	largeBody = make([]byte, largeSize)
	for i := 0; i < largeSize/2; i++ {
		largeBody[i] = 'a'
	}
	for i := largeSize / 2; i < largeSize; i++ {
		largeBody[i] = 'b'
	}
	readOnlyRequest = &http.Request{
		Method: "GET",
		URL:    &url.URL{Path: "/v1.0/hi"},
		Header: http.Header{},
	}
	readOnlyRequestWithHeader = &http.Request{
		Method: "GET",
		URL:    &url.URL{Path: "/v1.0/hi"},
		Header: http.Header{
			"Host":   {"localhost"},
			"Accept": {"text/plain"},
		},
	}
	readOnlyRequestWithLargeHeader = &http.Request{
		Method: "GET",
		URL:    &url.URL{Path: "/v1.0/hi"},
		Header: http.Header{
			"Host":   {"localhost"},
			"Accept": {"text/plain"},
			"data":   {string(largeBody)},
		},
	}
}

func postSmall() *http.Request {
	return &http.Request{
		Method:        http.MethodPost,
		ContentLength: int64(len(smallBody)),
		Body:          io.NopCloser(bytes.NewReader(smallBody)),
	}
}

func postLarge() *http.Request {
	return &http.Request{
		Method:        http.MethodPost,
		ContentLength: int64(len(largeBody)),
		Body:          io.NopCloser(bytes.NewReader(largeBody)),
	}
}

var benches = map[string]struct {
	bins       map[string][]byte
	newRequest func() *http.Request
}{
	"log": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchLogTinyGo,
			"wat":    test.BinBenchLogWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequest
		}},
	"get_uri": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchGetURITinyGo,
			"wat":    test.BinBenchGetURIWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequest
		}},
	"set_uri": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchSetURITinyGo,
			"wat":    test.BinBenchSetURIWat,
		},
		newRequest: func() *http.Request {
			return &http.Request{URL: &url.URL{}}
		}},
	"get_request_header_names none": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchGetRequestHeaderNamesTinyGo,
			"wat":    test.BinBenchGetRequestHeaderNamesWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequest
		}},
	"get_request_header_names small": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchGetRequestHeaderNamesTinyGo,
			"wat":    test.BinBenchGetRequestHeaderNamesWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequestWithHeader
		}},
	"get_request_header_names large": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchGetRequestHeaderNamesTinyGo,
			"wat":    test.BinBenchGetRequestHeaderNamesWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequestWithLargeHeader
		}},
	"get_request_header exists": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchGetRequestHeaderTinyGo,
			"wat":    test.BinBenchGetRequestHeaderWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequestWithHeader
		}},
	"get_request_header not exists": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchGetRequestHeaderTinyGo,
			"wat":    test.BinBenchGetRequestHeaderWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequest
		}},
	"read_request_body": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchReadRequestBodyTinyGo,
			"wat":    test.BinBenchReadRequestBodyWat,
		},
		newRequest: postSmall},
	"read_request_body_stream small": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchReadRequestBodyStreamTinyGo,
			"wat":    test.BinBenchReadRequestBodyStreamWat,
		},
		newRequest: postSmall},
	"read_request_body_stream large": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchReadRequestBodyStreamTinyGo,
			"wat":    test.BinBenchReadRequestBodyStreamWat,
		},
		newRequest: postLarge},
	"next": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchNextTinyGo,
			"wat":    test.BinBenchNextWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequest
		}},
	"set_status_code": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchSetStatusCodeTinyGo,
			"wat":    test.BinBenchSetStatusCodeWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequest
		}},
	"set_response_header": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchSetResponseHeaderTinyGo,
			"wat":    test.BinBenchSetResponseHeaderWat,
		},
		newRequest: func() *http.Request {
			return readOnlyRequest
		}},
	"write_response_body": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchWriteResponseBodyTinyGo,
			"wat":    test.BinBenchWriteResponseBodyWat,
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
	if statusCode == 500 {
		panic(statusCode)
	}
}
