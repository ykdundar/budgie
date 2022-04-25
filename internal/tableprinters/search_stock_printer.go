package tableprinters

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/ykdundar/budgie/api"
)

func SearchStockPrinter(ticker api.Ticker, head string) {
	t := table.NewWriter()
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Number:      1,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Number:      2,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Number:      3,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Number:      4,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Number:      5,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
	})
	t.SetStyle(table.StyleRounded)
	t.SetTitle(head)
	t.Style().Title.Align = text.AlignCenter
	t.Style().Title.Format = text.FormatUpper
	t.SetOutputMirror(os.Stdout)
	t.Style().Options.SeparateRows = true
	t.AppendHeader(table.Row{"NAME", "SYMBOL", "EXCHANGE", "COUNTRY", "CITY"})

	for _, v := range ticker.Data {
		t.AppendRow(table.Row{v.Name, v.Symbol, v.StockExchange.Acronym, v.StockExchange.Country, v.StockExchange.City})
		t.AppendSeparator()
	}

	t.Render()
}
