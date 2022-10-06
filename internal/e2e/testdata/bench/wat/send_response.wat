;; $ wat2wasm --debug-names send_response.wat
(module $send_response
  (import "http-handler" "send_response"
    (func $send_response (param i32 i32 i32)))

  (memory (export "memory") 1 (; 1 page==64KB ;))

  (global $body i32 (i32.const 0))
  (data (i32.const 16) "hello world")
  (global $body_len i32 (i32.const 11))

  (func $handle (export "handle")
    (call $send_response
      (i32.const 200)
      (global.get $body) (global.get $body_len)))
)
