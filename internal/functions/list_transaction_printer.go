package functions

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ykdundar/budgie/internal/objects"
	"os"
)

func ListTransactionPrinter(transactions []objects.Transaction, head string) {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{head, head, head, head, head, head, head, head}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"ID", "SYMBOL", "CATEGORY","DATE", "SHARES",
	"PRICE", "MARKET VALUE", "PURCHASE VALUE"})

	for _, v := range transactions {
		t.AppendRow(table.Row{v.Id, v.Ticker, v.TransactionCategory,
			v.TransactionDate, v.Shares, v.Price, v.MarketValue, v.PurchaseValue})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true
	t.Render()
}