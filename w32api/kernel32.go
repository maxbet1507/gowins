// Code generated to call native win32. DO NOT EDIT.

package w32api

import (
	"syscall"
	"unsafe"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	procCreateMutexW = kernel32.NewProc("CreateMutexW")
	procCloseHandle  = kernel32.NewProc("CloseHandle")
)

func CreateMutex(lpMutexAttributes *SECURITY_ATTRIBUTES, bInitialOwner bool, lpName string) (uintptr, error) {
	v1 := unsafe.Pointer(lpMutexAttributes)

	v2 := FALSE
	if bInitialOwner {
		v2 = TRUE
	}

	v3 := unsafe.Pointer(syscall.StringToUTF16Ptr(lpName))

	ret, _, err := procCreateMutexW.Call(uintptr(v1), uintptr(v2), uintptr(v3))
	return ret, err
}

func CloseHandle(hObject uintptr) (bool, error) {
	ret, _, err := procCloseHandle.Call(hObject)
	return ret != 0, err
}
