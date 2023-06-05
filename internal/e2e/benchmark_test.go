package e2e_test

import (
	"bytes"
	"context"
	_ "embed"
	"net/http"
	"testing"

	"github.com/http-wasm/http-wasm-guest-tinygo/internal/test"
	nethttp "github.com/http-wasm/http-wasm-host-go/handler/nethttp"
)

var (
	noopHandler http.Handler
	smallBody   []byte
	largeSize   int
	largeBody   []byte
)

func init() {
	noopHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	smallBody = []byte("hello world")
	largeSize = 4096 // 2x the read buffer size
	largeBody = make([]byte, largeSize)
	for i := 0; i < largeSize/2; i++ {
		largeBody[i] = 'a'
	}
	for i := largeSize / 2; i < largeSize; i++ {
		largeBody[i] = 'b'
	}
}

func get(url string) (req *http.Request) {
	req, _ = http.NewRequest(http.MethodGet, url+"/v1.0/hi", nil)
	return
}

func getWithLargeHeader(url string) (req *http.Request) {
	req, _ = http.NewRequest(http.MethodGet, url+"/v1.0/hi", nil)
	req.Header.Add("data", string(largeBody))
	return
}

func getWithQuery(url string) (req *http.Request) {
	req, _ = http.NewRequest(http.MethodGet, url+"/v1.0/hi?name=panda", nil)
	return
}

func getWithoutHeaders(url string) (req *http.Request) {
	req, _ = http.NewRequest(http.MethodGet, url+"/v1.0/hi", nil)
	req.Header = http.Header{}
	return
}

func post(url string) (req *http.Request) {
	body := bytes.NewReader(smallBody)
	req, _ = http.NewRequest(http.MethodPost, url, body)
	return
}

func postLarge(url string) (req *http.Request) {
	body := bytes.NewReader(largeBody)
	req, _ = http.NewRequest(http.MethodPost, url, body)
	return
}

var benches = map[string]struct {
	bins    map[string][]byte
	next    http.Handler
	request func(url string) *http.Request
}{
	"example wasi": {
		bins: map[string][]byte{
			"TinyGo": test.BinExampleWASI,
			"wat":    test.BinExampleWASIWat,
		},
		next:    test.HandlerExampleWASI,
		request: test.RequestExampleWASI,
	},
	"example router host response": {
		bins: map[string][]byte{
			"TinyGo": test.BinExampleRouter,
			"wat":    test.BinExampleRouterWat,
		},
		request: func(url string) (req *http.Request) {
			req, _ = http.NewRequest(http.MethodGet, url+"/host", nil)
			return
		},
	},
	"example router wasm response": {
		bins: map[string][]byte{
			"TinyGo": test.BinExampleRouter,
			"wat":    test.BinExampleRouterWat,
		},
		request: func(url string) (req *http.Request) {
			req, _ = http.NewRequest(http.MethodGet, url, nil)
			return
		},
	},
	"log": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchLogTinyGo,
			"wat":    test.BinBenchLogWat,
		},
		request: get,
	},
	"get_uri": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchGetURITinyGo,
			"wat":    test.BinBenchGetURIWat,
		},
		request: getWithQuery,
	},
	"set_uri": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchSetURITinyGo,
			"wat":    test.BinBenchSetURIWat,
		},
		request: getWithQuery,
	},
	"get_header_names none": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchGetHeaderValuesNamesTinyGo,
			"wat":    test.BinBenchGetHeaderValuesNamesWat,
		},
		request: getWithoutHeaders,
	},
	"get_header_names": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchGetHeaderValuesNamesTinyGo,
			"wat":    test.BinBenchGetHeaderValuesNamesWat,
		},
		request: get,
	},
	"get_header_names large": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchGetHeaderValuesNamesTinyGo,
			"wat":    test.BinBenchGetHeaderValuesNamesWat,
		},
		request: getWithLargeHeader,
	},
	"get_header_values exists": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchGetHeaderValuesTinyGo,
			"wat":    test.BinBenchGetHeaderValuesWat,
		},
		request: get,
	},
	"get_header_values not exists": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchGetHeaderValuesTinyGo,
			"wat":    test.BinBenchGetHeaderValuesWat,
		},
		request: getWithoutHeaders,
	},
	"set_header_value exists": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchSetHeaderValueTinyGo,
			"wat":    test.BinBenchSetHeaderValueWat,
		},
		request: get,
	},
	"set_header_value not exists": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchSetHeaderValueTinyGo,
			"wat":    test.BinBenchSetHeaderValueWat,
		},
		request: getWithoutHeaders,
	},
	"add_header_value exists": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchAddHeaderValueTinyGo,
			"wat":    test.BinBenchAddHeaderValueWat,
		},
		request: get,
	},
	"add_header_value not exists": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchAddHeaderValueTinyGo,
			"wat":    test.BinBenchAddHeaderValueWat,
		},
		request: getWithoutHeaders,
	},
	"remove_header exists": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchRemoveHeaderTinyGo,
			"wat":    test.BinBenchRemoveHeaderWat,
		},
		request: get,
	},
	"remove_header not exists": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchRemoveHeaderTinyGo,
			"wat":    test.BinBenchRemoveHeaderWat,
		},
		request: getWithoutHeaders,
	},
	"read_body": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchReadBodyTinyGo,
			"wat":    test.BinBenchReadBodyWat,
		},
		request: post,
	},
	"read_body_stream": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchReadBodyStreamTinyGo,
			"wat":    test.BinBenchReadBodyStreamWat,
		},
		request: post,
	},
	"read_body_stream large": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchReadBodyStreamTinyGo,
			"wat":    test.BinBenchReadBodyStreamWat,
		},
		request: postLarge,
	},
	"write_body": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchWriteBodyTinyGo,
			"wat":    test.BinBenchWriteBodyWat,
		},
		request: get,
	},
	"set_status_code": {
		bins: map[string][]byte{
			"TinyGo": test.BinBenchSetStatusCodeTinyGo,
			"wat":    test.BinBenchSetStatusCodeWat,
		},
		request: get,
	},
}

func Benchmark(b *testing.B) {
	for n, s := range benches {
		s := s
		b.Run(n, func(b *testing.B) {
			for n, bin := range s.bins {
				benchmark(b, n, bin, s.next, s.request)
			}
		})
	}
}

func benchmark(b *testing.B, name string, bin []byte, handler http.Handler, newRequest func(string) *http.Request) {
	ctx := context.Background()

	mw, err := nethttp.NewMiddleware(ctx, bin)
	if err != nil {
		b.Fatal(err)
	}
	defer mw.Close(ctx)

	if handler == nil {
		handler = noopHandler
	}
	h := mw.NewHandler(ctx, handler)

	b.Run(name, func(b *testing.B) {
		// We don't report allocations because memory allocations for TinyGo are
		// in wasm which isn't visible to the Go benchmark.
		for i := 0; i < b.N; i++ {
			h.ServeHTTP(fakeResponseWriter{}, newRequest("http://localhost"))
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
