package WindowExStyle

import (
	"github.com/lxn/win"
)

func ClientEdge(hwnd win.HWND, enable bool) {
	exstyle := win.GetWindowLong(hwnd, win.GWL_EXSTYLE)
	if enable {
		exstyle |= win.WS_EX_CLIENTEDGE
	} else {
		exstyle &= ^win.WS_EX_CLIENTEDGE
	}
	win.SetWindowLong(hwnd, win.GWL_EXSTYLE, exstyle)
	win.SetWindowPos(hwnd, 0, 0, 0, 0, 0, win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_NOZORDER|win.SWP_FRAMECHANGED)
}
