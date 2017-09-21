package main

import (
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"github.com/maxbet1507/gowins/WindowExStyle"
)

func main() {
	var le *walk.LineEdit

	d := declarative.Dialog{
		FixedSize: true,
		Layout:    declarative.VBox{},
		Children: []declarative.Widget{
			declarative.Label{Text: "File"},
			declarative.Label{Text: "ファイル"},
			declarative.LineEdit{AssignTo: &le},
			declarative.PushButton{Text: "OK", OnClicked: func() {
				WindowExStyle.ClientEdge(le.Handle(), false)
			}},
		},
		Title: "サンプル",
	}
	d.Run(nil)
}
