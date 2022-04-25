package tableprinters

import (
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/ykdundar/budgie/internal/functions"
	"github.com/ykdundar/budgie/internal/objects"
)

var tableConfig = []table.ColumnConfig{
	{
		Number:      1,
		Align:       text.AlignCenter,
		AlignHeader: text.AlignCenter,
		WidthMin:    15,
		WidthMax:    15,
	},
	{
		Number:      2,
		Align:       text.AlignCenter,
		AlignHeader: text.AlignCenter,
		WidthMin:    15,
		WidthMax:    15,
	},
	{
		Number:      3,
		Align:       text.AlignCenter,
		AlignHeader: text.AlignCenter,
		WidthMin:    15,
		WidthMax:    15,
	},
	{
		Number:      4,
		Align:       text.AlignCenter,
		AlignHeader: text.AlignCenter,
		WidthMin:    15,
		WidthMax:    15,
	},
	{
		Number:      5,
		Align:       text.AlignCenter,
		AlignHeader: text.AlignCenter,
		WidthMin:    15,
		WidthMax:    15,
	},
	{
		Number:      6,
		Align:       text.AlignCenter,
		AlignHeader: text.AlignCenter,
		WidthMin:    15,
		WidthMax:    15,
	},
	{
		Number:      7,
		Align:       text.AlignCenter,
		AlignHeader: text.AlignCenter,
		WidthMin:    15,
		WidthMax:    15,
	},
	{
		Number:      8,
		Align:       text.AlignCenter,
		AlignHeader: text.AlignCenter,
		WidthMin:    15,
		WidthMax:    15,
	},
}

func ListTransactionPrinter(transactions []objects.Transaction, head string) {
	t := table.NewWriter()
	t.SetColumnConfigs(tableConfig)
	t.SetStyle(table.StyleRounded)
	t.SetTitle(head)
	t.Style().Title.Align = text.AlignCenter
	t.Style().Title.Format = text.FormatUpper
	t.SetOutputMirror(os.Stdout)
	t.Style().Options.SeparateRows = true
	t.AppendHeader(table.Row{"ID", "SYMBOL", "CATEGORY", "DATE", "SHARES",
		"PRICE", "MARKET VALUE", "PURCHASE VALUE"})

	for _, v := range transactions {
		var marketValue float64
		lastPrice := functions.GetLatestPrice(v.Ticker)
		marketValue = lastPrice * float64(v.Shares)

		t.AppendRow(
			table.Row{v.ID, v.Ticker, functions.IntToString(v.TransactionCategory),
				time.Unix(int64(v.TransactionDate), 0).Format("2006-1-2"),
				v.Shares, functions.CropNumbers(v.Price),
				functions.CropNumbers(marketValue), functions.CropNumbers(v.PurchaseValue)})
		t.AppendSeparator()
	}
	t.Render()
}
