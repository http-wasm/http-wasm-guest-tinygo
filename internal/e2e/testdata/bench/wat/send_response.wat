;; $ wat2wasm --debug-names send_response.wat
(module $send_response
  (import "http-handler" "set_status_code"
    (func $set_status_code (param i32)))
  (import "http-handler" "set_response_body"
    (func $set_response_body (param i32)))

  (memory (export "memory") 1 (; 1 page==64KB ;))

  (global $body i32 (i32.const 0))
  (data (i32.const 16) "hello world")
  (global $body_len i32 (i32.const 11))

  (func $handle (export "handle")
    (call $set_status_code
      (i32.const 200))
    (call $set_response_body
      (global.get $body) (global.get $body_len))
  )
)
