package main

import (
	"github.com/lxn/walk/declarative"
	"github.com/maxbet1507/gowins/NonClientMetrics"
)

func main() {
	nonClientMetrics, _ := NonClientMetrics.Get()

	d := declarative.Dialog{
		Font:      nonClientMetrics.Message.Font,
		FixedSize: true,
		Layout:    declarative.VBox{},
		Children: []declarative.Widget{
			declarative.Label{Text: "File"},
			declarative.Label{Text: "ファイル"},
			declarative.LineEdit{},
			declarative.PushButton{Text: "OK"},
		},
		Title: "サンプル",
	}
	d.Run(nil)
}
