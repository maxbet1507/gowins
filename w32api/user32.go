// skipping golint for this file
// Code generated  DO NOT EDIT.

package w32api

import (
	"reflect"
	"syscall"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	procGetDC     = user32.NewProc("GetDC")
	procReleaseDC = user32.NewProc("ReleaseDC")

	procSystemParametersInfoW = user32.NewProc("SystemParametersInfoW")
)

func GetDC(hwnd uintptr) uintptr {
	hdc, _, _ := procGetDC.Call(hwnd)
	return hdc
}

func ReleaseDC(hwnd uintptr, hdc uintptr) bool {
	ret, _, _ := procReleaseDC.Call(hwnd, hdc)
	return int32(ret) == 1
}

func SystemParametersInfo(uiAction, uiParam uint32, pvParam interface{}, fWinIni uint32) bool {
	var up uintptr

	if pvParam != nil {
		if rv := reflect.ValueOf(pvParam); rv.Kind() == reflect.Ptr {
			up = rv.Pointer()
		} else {
			return false
		}
	}

	ret, _, _ := procSystemParametersInfoW.Call(uintptr(uiAction), uintptr(uiParam), up, uintptr(fWinIni))
	return ret != 0
}
