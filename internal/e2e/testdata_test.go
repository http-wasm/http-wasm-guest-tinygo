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

//go:embed testdata/bench/tinygo/get_path/main.wasm
var BinBenchGetPathTinyGo []byte

//go:embed testdata/bench/wat/get_path.wasm
var BinBenchGetPathWat []byte

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

//go:embed testdata/bench/tinygo/set_path/main.wasm
var BinBenchSetPathTinyGo []byte

//go:embed testdata/bench/wat/set_path.wasm
var BinBenchSetPathWat []byte

//go:embed testdata/bench/tinygo/set_response_header/main.wasm
var BinBenchSetResponseHeaderTinyGo []byte

//go:embed testdata/bench/wat/set_response_header.wasm
var BinBenchSetResponseHeaderWat []byte
