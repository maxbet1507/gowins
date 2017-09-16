// skipping golint for this file
// Code generated  DO NOT EDIT.

package gowins

import (
	"reflect"
	"syscall"
)

var (
	gdi32  = syscall.NewLazyDLL("gdi32.dll")
	user32 = syscall.NewLazyDLL("user32.dll")

	procGetDC     = user32.NewProc("GetDC")
	procReleaseDC = user32.NewProc("ReleaseDC")

	procGetDeviceCaps         = gdi32.NewProc("GetDeviceCaps")
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

const (
	LOGPIXELSY = 90
)

func GetDeviceCaps(hdc uintptr, nIndex int32) int32 {
	ret, _, _ := procGetDeviceCaps.Call(hdc, uintptr(nIndex))
	return int32(ret)
}

const (
	LF_FACESIZE = 32
	FW_NORMAL   = 400
)

type LOGFONT struct {
	Height         int32
	Width          int32
	Escapement     int32
	Orientation    int32
	Weight         int32
	Italic         byte
	Underline      byte
	StrikeOut      byte
	CharSet        byte
	OutPrecision   byte
	ClipPrecision  byte
	Quality        byte
	PitchAndFamily byte
	FaceName       [LF_FACESIZE]uint16
}

const (
	SPI_GETNONCLIENTMETRICS = 41
)

type NONCLIENTMETRICS struct {
	Size            uint32
	BorderWidth     int32
	ScrollWidth     int32
	ScrollHeight    int32
	CaptionWidth    int32
	CaptionHeight   int32
	CaptionFont     LOGFONT
	SmCaptionWidth  int32
	SmCaptionHeight int32
	SmCaptionFont   LOGFONT
	MenuWidth       int32
	MenuHeight      int32
	MenuFont        LOGFONT
	StatusFont      LOGFONT
	MessageFont     LOGFONT
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
