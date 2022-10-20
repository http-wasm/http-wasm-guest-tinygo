[![Build](https://github.com/http-wasm/http-wasm-guest-tinygo/workflows/build/badge.svg)](https://github.com/http-wasm/http-wasm-guest-tinygo)
[![Go Report Card](https://goreportcard.com/badge/github.com/http-wasm/http-wasm-guest-tinygo)](https://goreportcard.com/report/github.com/http-wasm/http-wasm-guest-tinygo)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

# http-wasm Guest Library for TinyGo

[http-wasm][1] is HTTP client middleware implemented in [WebAssembly][2].
This is a [TinyGo WASI][3] library that implements the [Guest ABI][4].

## Example
The following is an [example](examples/router) of routing middleware:

```go
package main

import (
	"strings"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleFn = handle
}

// handle implements a simple HTTP router.
func handle(req api.Request, resp api.Response, next api.Next) {
	// If the URI starts with /host, trim it and dispatch to the next handler.
	if uri := req.GetURI(); strings.HasPrefix(uri, "/host") {
		req.SetURI(uri[5:])
		next()
	} else { // Serve a static response
		resp.Headers().Set("Content-Type", "text/plain")
		resp.Body().WriteString("hello")
	}
}
```

If you make changes, you can rebuild it like so:
```sh
tinygo build -o examples/router/main.wasm -scheduler=none --no-debug -target=wasi examples/router/main.go
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
