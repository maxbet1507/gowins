// Code generated to call native win32. DO NOT EDIT.

package w32api

import (
	"syscall"
)

var (
	gdi32 = syscall.NewLazyDLL("gdi32.dll")

	procGetDeviceCaps = gdi32.NewProc("GetDeviceCaps")
)

func GetDeviceCaps(hdc uintptr, nIndex int32) (int32, error) {
	ret, _, err := procGetDeviceCaps.Call(hdc, uintptr(nIndex))
	return int32(ret), err
}
