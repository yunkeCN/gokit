package util

import "unsafe"

func Byte2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Str2Byte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
