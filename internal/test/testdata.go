package test

import (
	_ "embed"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var BinExampleWASI = func() []byte {
	return binExample("wasi")
}()

var BinExampleRewrite = func() []byte {
	return binExample("rewrite")
}()

func binExample(name string) []byte {
	p := path.Join("..", "..", "examples", name, "main.wasm")
	if wasm, err := os.ReadFile(p); err != nil {
		log.Panicln(err)
		return nil
	} else {
		return wasm
	}
}

var requestExampleConsoleBody = `{"hello": "panda"}`

func RequestExampleWASI(url string) (req *http.Request) {
	body := strings.NewReader(requestExampleConsoleBody)
	req, _ = http.NewRequest(http.MethodPost, url+"/v1.0/hi?name=panda", body)
	req.Header.Set("Content-Type", "application/json")
	req.Host = "localhost"
	return
}

var HandlerExampleWASI = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Set-Cookie", "a=b") // rewrite of multiple headers
	w.Header().Add("Set-Cookie", "c=d")

	// Use chunked encoding so we can set a test trailer
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Trailer", "grpc-status")
	w.Header().Set(http.TrailerPrefix+"grpc-status", "1")
	w.Write([]byte(`{"hello": "world"}`)) // nolint
})

//go:embed testdata/e2e/wat/wasi.wasm
var BinE2EWASIWat []byte

//go:embed testdata/e2e/tinygo/features/main.wasm
var BinE2EFeaturesTinyGo []byte

//go:embed testdata/e2e/tinygo/log/main.wasm
var BinE2ELog []byte

//go:embed testdata/bench/tinygo/log/main.wasm
var BinBenchLogTinyGo []byte

//go:embed testdata/bench/wat/log.wasm
var BinBenchLogWat []byte

//go:embed testdata/bench/tinygo/get_uri/main.wasm
var BinBenchGetURITinyGo []byte

//go:embed testdata/bench/wat/get_uri.wasm
var BinBenchGetURIWat []byte

//go:embed testdata/bench/tinygo/set_uri/main.wasm
var BinBenchSetURITinyGo []byte

//go:embed testdata/bench/wat/set_uri.wasm
var BinBenchSetURIWat []byte

//go:embed testdata/bench/tinygo/get_request_header_names/main.wasm
var BinBenchGetRequestHeaderNamesTinyGo []byte

//go:embed testdata/bench/wat/get_request_header_names.wasm
var BinBenchGetRequestHeaderNamesWat []byte

//go:embed testdata/bench/tinygo/get_request_header/main.wasm
var BinBenchGetRequestHeaderTinyGo []byte

//go:embed testdata/bench/wat/get_request_header.wasm
var BinBenchGetRequestHeaderWat []byte

//go:embed testdata/bench/tinygo/read_request_body/main.wasm
var BinBenchReadRequestBodyTinyGo []byte

//go:embed testdata/bench/wat/read_request_body.wasm
var BinBenchReadRequestBodyWat []byte

//go:embed testdata/bench/tinygo/read_request_body_stream/main.wasm
var BinBenchReadRequestBodyStreamTinyGo []byte

//go:embed testdata/bench/wat/read_request_body_stream.wasm
var BinBenchReadRequestBodyStreamWat []byte

//go:embed testdata/bench/tinygo/next/main.wasm
var BinBenchNextTinyGo []byte

//go:embed testdata/bench/wat/next.wasm
var BinBenchNextWat []byte

//go:embed testdata/bench/tinygo/set_status_code/main.wasm
var BinBenchSetStatusCodeTinyGo []byte

//go:embed testdata/bench/wat/set_status_code.wasm
var BinBenchSetStatusCodeWat []byte

//go:embed testdata/bench/tinygo/set_response_header/main.wasm
var BinBenchSetResponseHeaderTinyGo []byte

//go:embed testdata/bench/wat/set_response_header.wasm
var BinBenchSetResponseHeaderWat []byte

//go:embed testdata/bench/tinygo/write_response_body/main.wasm
var BinBenchWriteResponseBodyTinyGo []byte

//go:embed testdata/bench/wat/write_response_body.wasm
var BinBenchWriteResponseBodyWat []byte
