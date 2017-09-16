package NonClientMetrics

import (
	"fmt"
	"math"
	"syscall"
	"unsafe"

	"github.com/lxn/walk/declarative"
	"github.com/maxbet1507/gowins/w32api"
)

type Border struct {
	Width int32
}

type Scroll struct {
	Width  int32
	Height int32
}

type Caption struct {
	Width  int32
	Height int32
	Font   declarative.Font
}

type SmCaption struct {
	Width  int32
	Height int32
	Font   declarative.Font
}

type Menu struct {
	Width  int32
	Height int32
	Font   declarative.Font
}

type Status struct {
	Font declarative.Font
}

type Message struct {
	Font declarative.Font
}

type NonClientMetrics struct {
	Border    Border
	Scroll    Scroll
	Caption   Caption
	SmCaption SmCaption
	Menu      Menu
	Status    Status
	Message   Message
}

func Get() (*NonClientMetrics, error) {
	nonClientMetrics := w32api.NONCLIENTMETRICS{}
	nonClientMetrics.Size = uint32(unsafe.Sizeof(nonClientMetrics))

	if !w32api.SystemParametersInfo(w32api.SPI_GETNONCLIENTMETRICS, nonClientMetrics.Size, &nonClientMetrics, 0) {
		return nil, fmt.Errorf("failed to call SystemParametersInfo")
	}

	logPixcelsY := logPixcelsY()

	return &NonClientMetrics{
		Border: Border{
			Width: nonClientMetrics.BorderWidth,
		},
		Scroll: Scroll{
			Width:  nonClientMetrics.ScrollWidth,
			Height: nonClientMetrics.ScrollHeight,
		},
		Caption: Caption{
			Width:  nonClientMetrics.CaptionWidth,
			Height: nonClientMetrics.CaptionHeight,
			Font:   convertLogFontToWalkDeclarative(nonClientMetrics.CaptionFont, logPixcelsY),
		},
		SmCaption: SmCaption{
			Width:  nonClientMetrics.SmCaptionWidth,
			Height: nonClientMetrics.SmCaptionHeight,
			Font:   convertLogFontToWalkDeclarative(nonClientMetrics.SmCaptionFont, logPixcelsY),
		},
		Menu: Menu{
			Width:  nonClientMetrics.MenuWidth,
			Height: nonClientMetrics.MenuHeight,
			Font:   convertLogFontToWalkDeclarative(nonClientMetrics.MenuFont, logPixcelsY),
		},
		Status: Status{
			Font: convertLogFontToWalkDeclarative(nonClientMetrics.StatusFont, logPixcelsY),
		},
		Message: Message{
			Font: convertLogFontToWalkDeclarative(nonClientMetrics.MessageFont, logPixcelsY),
		},
	}, nil
}

func logPixcelsY() int32 {
	hdc := w32api.GetDC(0)
	defer w32api.ReleaseDC(0, hdc)
	return w32api.GetDeviceCaps(hdc, w32api.LOGPIXELSY)
}

func convertLogFontToWalkDeclarative(lf w32api.LOGFONT, logPixcelsY int32) declarative.Font {
	return declarative.Font{
		Family:    syscall.UTF16ToString(lf.FaceName[:]),
		PointSize: int(math.Floor(float64(-lf.Height)*float64(72)/float64(logPixcelsY) + .5)),
		Bold:      lf.Weight > w32api.FW_NORMAL,
		Italic:    lf.Italic != 0,
		Underline: lf.Underline != 0,
		StrikeOut: lf.StrikeOut != 0,
	}
}
