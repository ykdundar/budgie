package tableprinters

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/ykdundar/budgie/database/transactions"
	"github.com/ykdundar/budgie/internal/functions"
)

func ReportTransactionPrinter(transactionList []transactions.TransactionSum, head string) {
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
	t.AppendHeader(table.Row{"SYMBOL", "SHARES", "PURCHASE VALUE", "MARKET VALUE", "DIFFERENCE"})

	for _, v := range transactionList {
		lastPrice := functions.GetLatestPrice(v.Ticker)

		var marketValue float64
		var diff float64

		marketValue = lastPrice * float64(v.Shares)
		diff = v.PurchaseValue - marketValue

		t.AppendRow(table.Row{v.Ticker, v.Shares, functions.CropNumbers(v.PurchaseValue), functions.CropNumbers(marketValue), functions.CropNumbers(diff)})
		t.AppendSeparator()
	}

	t.Render()
}
