package NonClientMetrics

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/lxn/walk/declarative"
	"github.com/maxbet1507/gowins"
)

var (
	BorderWidth     int32
	ScrollWidth     int32
	ScrollHeight    int32
	CaptionWidth    int32
	CaptionHeight   int32
	CaptionFont     declarative.Font
	SmCaptionWidth  int32
	SmCaptionHeight int32
	SmCaptionFont   declarative.Font
	MenuWidth       int32
	MenuHeight      int32
	MenuFont        declarative.Font
	StatusFont      declarative.Font
	MessageFont     declarative.Font
)

func init() {
	if !Update() {
		panic("failed to call SystemParametersInfo")
	}
}

func Update() bool {
	nonClientMetrics := gowins.NONCLIENTMETRICS{}
	nonClientMetrics.Size = uint32(unsafe.Sizeof(nonClientMetrics))

	if !gowins.SystemParametersInfo(gowins.SPI_GETNONCLIENTMETRICS, nonClientMetrics.Size, &nonClientMetrics, 0) {
		return false
	}

	logPixcelsY := logPixcelsY()

	BorderWidth = nonClientMetrics.BorderWidth
	ScrollWidth = nonClientMetrics.ScrollWidth
	ScrollHeight = nonClientMetrics.ScrollHeight
	CaptionWidth = nonClientMetrics.CaptionWidth
	CaptionHeight = nonClientMetrics.CaptionHeight
	CaptionFont = convertLogFontToWalkDeclarative(nonClientMetrics.CaptionFont, logPixcelsY)
	SmCaptionWidth = nonClientMetrics.SmCaptionWidth
	SmCaptionHeight = nonClientMetrics.SmCaptionHeight
	SmCaptionFont = convertLogFontToWalkDeclarative(nonClientMetrics.SmCaptionFont, logPixcelsY)
	MenuWidth = nonClientMetrics.MenuWidth
	MenuHeight = nonClientMetrics.MenuHeight
	MenuFont = convertLogFontToWalkDeclarative(nonClientMetrics.MenuFont, logPixcelsY)
	StatusFont = convertLogFontToWalkDeclarative(nonClientMetrics.StatusFont, logPixcelsY)
	MessageFont = convertLogFontToWalkDeclarative(nonClientMetrics.MessageFont, logPixcelsY)

	return true
}

func logPixcelsY() int32 {
	hdc := gowins.GetDC(0)
	defer gowins.ReleaseDC(0, hdc)
	return gowins.GetDeviceCaps(hdc, gowins.LOGPIXELSY)
}

func convertLogFontToWalkDeclarative(lf gowins.LOGFONT, logPixcelsY int32) declarative.Font {
	return declarative.Font{
		Family:    syscall.UTF16ToString(lf.FaceName[:]),
		PointSize: int(math.Floor(float64(-lf.Height)*float64(72)/float64(logPixcelsY) + .5)),
		Bold:      lf.Weight > gowins.FW_NORMAL,
		Italic:    lf.Italic != 0,
		Underline: lf.Underline != 0,
		StrikeOut: lf.StrikeOut != 0,
	}
}
