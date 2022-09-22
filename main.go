package httpwasm

import "github.com/tetratelabs/tinymem"

func Log(message string) {
	ptr, size := tinymem.StringToPtr(message)
	if size == 0 {
		return // don't incur host call overhead
	}
	log(ptr, size)
}
