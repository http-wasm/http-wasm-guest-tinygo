package test

import (
	_ "embed"
	"log"
	"os"
	"path"
)

var BinExample = func() []byte {
	p := path.Join("..", "..", "example", "main.wasm")
	if wasm, err := os.ReadFile(p); err != nil {
		log.Panicln(err)
		return nil
	} else {
		return wasm
	}
}()

//go:embed testdata/log/main.wasm
var BinLog []byte

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
