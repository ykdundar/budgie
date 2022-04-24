package functions

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ykdundar/budgie/api"
)

func SearchStockPrinter(ticker api.Ticker, head string) {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{head, head, head, head, head}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"NAME", "SYMBOL", "EXCHANGE", "COUNTRY", "CITY"})

	for _, v := range ticker.Data {
		t.AppendRow(table.Row{v.Name, v.Symbol, v.StockExchange.Acronym, v.StockExchange.Country, v.StockExchange.City})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true
	t.Render()
}
