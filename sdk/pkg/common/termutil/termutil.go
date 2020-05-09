package termutil

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"os"
)

var (
	bsStyle = table.Style{
		Name: "bsStyle",
		Box:  table.StyleBoxLight,
		Color: table.ColorOptions{
			Header:      text.Colors{text.BgBlack, text.FgBlue},
			IndexColumn: text.Colors{text.BgBlack, text.FgHiBlue},
		},
		Format: table.FormatOptionsDefault,
		Options: table.Options{
			DrawBorder:      true,
			SeparateColumns: false,
			SeparateHeader:  true,
			SeparateRows:    true,
		},
		Title: table.TitleOptionsDefault,
	}
)

func CreateTable(kv map[string]string, title string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(bsStyle)
	var rows []table.Row
	for k, v := range kv {
		rows = append(rows, table.Row{fmt.Sprintf("%s:",k), v})
	}
	t.AppendHeader(table.Row{title})
	t.AppendRows(rows)
	t.Render()
}
