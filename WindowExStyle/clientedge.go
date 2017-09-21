package WindowExStyle

import (
	"github.com/maxbet1507/gowins/w32api"
)

func ClientEdge(hwnd uintptr, enable bool) {
	exstyle, _ := w32api.GetWindowLongPtr(hwnd, w32api.GWL_EXSTYLE)
	if enable {
		w32api.SetWindowLongPtr(hwnd, w32api.GWL_EXSTYLE, uintptr(int32(exstyle)|w32api.WS_EX_CLIENTEDGE))
	} else {
		w32api.SetWindowLongPtr(hwnd, w32api.GWL_EXSTYLE, uintptr(int32(exstyle) & ^w32api.WS_EX_CLIENTEDGE))
	}
	w32api.SetWindowPos(hwnd, 0, 0, 0, 0, 0, w32api.SWP_NOMOVE|w32api.SWP_NOSIZE|w32api.SWP_NOZORDER|w32api.SWP_FRAMECHANGED)
}
