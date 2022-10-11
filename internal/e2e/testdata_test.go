package e2e_test

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

//go:embed testdata/bench/tinygo/get_uri/main.wasm
var BinBenchGetURITinyGo []byte

//go:embed testdata/bench/wat/get_uri.wasm
var BinBenchGetURIWat []byte

//go:embed testdata/bench/tinygo/get_request_header/main.wasm
var BinBenchGetRequestHeaderTinyGo []byte

//go:embed testdata/bench/wat/get_request_header.wasm
var BinBenchGetRequestHeaderWat []byte

//go:embed testdata/bench/tinygo/log/main.wasm
var BinBenchLogTinyGo []byte

//go:embed testdata/bench/wat/log.wasm
var BinBenchLogWat []byte

//go:embed testdata/bench/tinygo/next/main.wasm
var BinBenchNextTinyGo []byte

//go:embed testdata/bench/wat/next.wasm
var BinBenchNextWat []byte

//go:embed testdata/bench/tinygo/send_response/main.wasm
var BinBenchSendResponseTinyGo []byte

//go:embed testdata/bench/wat/send_response.wasm
var BinBenchSendResponseWat []byte

//go:embed testdata/bench/tinygo/set_uri/main.wasm
var BinBenchSetURITinyGo []byte

//go:embed testdata/bench/wat/set_uri.wasm
var BinBenchSetURIWat []byte

//go:embed testdata/bench/tinygo/set_response_header/main.wasm
var BinBenchSetResponseHeaderTinyGo []byte

//go:embed testdata/bench/wat/set_response_header.wasm
var BinBenchSetResponseHeaderWat []byte
