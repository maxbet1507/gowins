package main

import (
	"fmt"

	"github.com/lxn/walk/declarative"
	"github.com/maxbet1507/gowins/NonClientMetrics"
)

func main() {
	fmt.Println(NonClientMetrics.MessageFont)

	d := declarative.Dialog{
		Font:   NonClientMetrics.MessageFont,
		Layout: declarative.VBox{},
		Children: []declarative.Widget{
			declarative.Label{Text: "File"},
			declarative.Label{Text: "ファイル"},
		},
	}
	d.Run(nil)
}
