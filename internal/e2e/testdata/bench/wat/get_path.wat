;; $ wat2wasm --debug-names get_path.wat
(module $get_path
  (import "http-handler" "get_path"
    (func $get_path (param i32 i32) (result i32)))

  (memory (export "memory") 1 (; 1 page==64KB ;))
  (global $buf i32 (i32.const 0))
  (global $buf_limit i32 (i32.const 64))

  (func $handle (export "handle")
    (call $get_path
      (global.get $buf)
      (global.get $buf_limit))
    (drop))
)
