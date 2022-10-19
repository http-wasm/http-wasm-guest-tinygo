[![Build](https://github.com/http-wasm/http-wasm-guest-tinygo/workflows/build/badge.svg)](https://github.com/http-wasm/http-wasm-guest-tinygo)
[![Go Report Card](https://goreportcard.com/badge/github.com/http-wasm/http-wasm-guest-tinygo)](https://goreportcard.com/report/github.com/http-wasm/http-wasm-guest-tinygo)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

# http-wasm Guest Library for TinyGo

[http-wasm][1] is HTTP client middleware implemented in [WebAssembly][2].
This is a [TinyGo WASI][3] library that implements the [Guest ABI][4].

## Example
The following is an [example](examples/rewrite) of rewriting the request URI.

```go
package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleFn = handle
}

// handle rewrites the request before dispatching to the next handler.
//
// Note: This is not a redirect, rather in-process routing.
func handle(req api.Request, _ api.Response, next api.Next) {
	if req.GetURI() == "/v1.0/hi?name=panda" {
		req.SetURI("/v1.0/hello?name=teddy")
	}
	next()
}
```

If you make changes, you can rebuild it like so:
```sh
tinygo build -o examples/rewrite/main.wasm -scheduler=none --no-debug -target=wasi examples/rewrite/main.go
```

There are also more [examples](examples) you may wish to try out!

# WARNING: This is a proof of concept!

The current maturity phase is proof of concept. Once this has a working example
in [dapr][6], we will go back and revisit things intentionally deferred.

Meanwhile, minor details and test coverage will fall short of production
standards. This helps us deliver the proof-of-concept faster and prevents
wasted energy in the case that the concept isn't acceptable at all.

[1]: https://github.com/http-wasm
[2]: https://webassembly.org/
[3]: https://wazero.io/languages/tinygo/
[4]: https://github.com/http-wasm/http-wasm-abi
[5]: https://github.com/http-wasm/components-contrib/
