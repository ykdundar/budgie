package functions

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ykdundar/budgie/api"
	"os"
)

func PrintOnTable(intraday api.Intraday, name string, headers table.Row) {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{name, name, name, name, name, name, name, name, name}, rowConfigAutoMerge)
	t.AppendHeader(headers)

	for _, v := range intraday.Data {
		t.AppendRow(table.Row{v.Open, v.High, v.Low, v.Last,
			v.Close, v.Volume, v.Date, v.Symbol, v.Exchange})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true
	t.Render()
}
