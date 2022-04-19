package functions

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ykdundar/budgie/api"
	"os"
	"time"
)

func ShowCmdPrinter(intraday api.Intraday, header string) {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{header, header, header, header, header, header, header, header}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"SYMBOL", "OPEN", "HIGH", "LOW", "LAST", "CLOSE", "DATE", "EXCHANGE"})

	for _, v := range intraday.Data {
		formattedDate, _ := time.Parse("2006-01-02T15:04:05-0700", v.Date)
		t.AppendRow(table.Row{v.Symbol, v.Open, v.High, v.Low, v.Last,
			v.Close, formattedDate.Format("2006-01-02"), v.Exchange})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true
	t.Render()
}
