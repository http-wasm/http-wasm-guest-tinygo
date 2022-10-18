package handler

import "unsafe"

func getString(fn func(ptr uintptr, limit uint32) (len uint32)) string {
	if b := getBytes(fn); len(b) == 0 {
		return ""
	} else {
		return string(b)
	}
}

func getBytes(fn func(ptr uintptr, limit uint32) (len uint32)) []byte {
	size := fn(0, 0)
	if size == 0 {
		return nil
	}
	buf := make([]byte, size)
	ptr := uintptr(unsafe.Pointer(&buf[0]))
	_ = fn(ptr, size)
	return buf
}
