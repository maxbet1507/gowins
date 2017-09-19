package w32api

import (
	"syscall"
)

// ErrorCode returns https://msdn.microsoft.com/en-us/library/cc231199.aspx
func ErrorCode(err error) int {
	if code, ok := err.(syscall.Errno); ok {
		return int(code)
	}
	return -1
}
