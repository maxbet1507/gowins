// Code generated to call native win32. DO NOT EDIT.

package w32api

import (
	"reflect"
	"syscall"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	procGetDC                 = user32.NewProc("GetDC")
	procReleaseDC             = user32.NewProc("ReleaseDC")
	procSystemParametersInfoW = user32.NewProc("SystemParametersInfoW")
	procGetWindowLongPtrW     = user32.NewProc("GetWindowLongPtrW")
	procSetWindowLongPtrW     = user32.NewProc("SetWindowLongPtrW")
	procSetWindowPos          = user32.NewProc("SetWindowPos")
)

func GetDC(hwnd uintptr) (uintptr, error) {
	hdc, _, err := procGetDC.Call(hwnd)
	return hdc, err
}

func ReleaseDC(hwnd uintptr, hdc uintptr) (bool, error) {
	ret, _, err := procReleaseDC.Call(hwnd, hdc)
	return int32(ret) == 1, err
}

func SystemParametersInfo(uiAction, uiParam uint32, pvParam interface{}, fWinIni uint32) (bool, error) {
	var up uintptr

	if pvParam != nil {
		if rv := reflect.ValueOf(pvParam); rv.Kind() == reflect.Ptr {
			up = rv.Pointer()
		} else {
			return false, syscall.Errno(0)
		}
	}

	ret, _, err := procSystemParametersInfoW.Call(uintptr(uiAction), uintptr(uiParam), up, uintptr(fWinIni))
	return ret != 0, err
}

func GetWindowLongPtr(hwnd uintptr, nIndex int32) (uintptr, error) {
	ret, _, err := procGetWindowLongPtrW.Call(hwnd, uintptr(nIndex))
	return ret, err
}

func SetWindowLongPtr(hwnd uintptr, nIndex int32, dwNewLong uintptr) (uintptr, error) {
	ret, _, err := procSetWindowLongPtrW.Call(hwnd, uintptr(nIndex), dwNewLong)
	return ret, err
}

func SetWindowPos(hwnd, hwndInsertAfter uintptr, x, y, cx, cy int32, uFlags uint32) (bool, error) {
	ret, _, err := procSetWindowPos.Call(hwnd, hwndInsertAfter, uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags))
	return ret != 0, err
}
