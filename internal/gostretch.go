package internal

import (
	"crypto/sha512"
)

const (
	LOW = uintptr(512 * 1024 * 1024)
	MID = uintptr(1024 * 1024 * 1024)
	HIGH = uintptr(2048 * 1024 * 1024)
	EXTREME = uintptr(4096 * 1024 * 1024)
)

func Stretch(data string, size uintptr, buf []byte) []byte {
	buf = buf[:0]
	inBuf := [64]byte{}
	inBuf = sha512.Sum512([]byte(data))
	buf = append(buf, inBuf[:]...)
	for uintptr(len(buf)) < size {
		inBuf = sha512.Sum512(buf)
		buf = append(buf, inBuf[:]...)
		if uintptr(len(buf)) * 2 <= size {
			buf = append(buf, buf...)
		} else if uintptr(len(buf)) < size {
			buf = append(buf, buf[:size - uintptr(len(buf))]...)
		} else {
			break
		}
	}
	inBuf = sha512.Sum512(buf[:])
	for i := 0; i < 100; i++ {
		inBuf = sha512.Sum512(inBuf[:])
	}
	copy(buf[:64], inBuf[:])
	return buf[:64]
}
