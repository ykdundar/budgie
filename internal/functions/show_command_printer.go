package functions

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ykdundar/budgie/api"
	"os"
)

func ShowCmdPrinter(intraday api.Intraday, header string) {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{header, header, header, header, header, header, header, header, header}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"SYMBOL", "OPEN", "HIGH", "LOW", "LAST", "CLOSE", "VOLUME", "DATE", "EXCHANGE"})

	for _, v := range intraday.Data {
		t.AppendRow(table.Row{v.Symbol, v.Open, v.High, v.Low, v.Last,
			v.Close, v.Volume, v.Date, v.Exchange})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true
	t.Render()
}
