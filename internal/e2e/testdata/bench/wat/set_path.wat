;; $ wat2wasm --debug-names set_path.wat
(module $set_path
  (import "http-handler" "set_path"
    (func $set_path (param i32 i32)))

  (memory (export "memory") 1 (; 1 page==64KB ;))
  (global $path i32 (i32.const 0))
  (data (i32.const 0) "/v1.0/hello")
  (global $path_len i32 (i32.const 11))

  (func $handle (export "handle")
    (call $set_path
      (global.get $path)
      (global.get $path_len)))
)
