package functions

import (
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ykdundar/budgie/internal/objects"
)

func ListTransactionPrinter(transactions []objects.Transaction, head string) {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{head, head, head, head, head, head, head, head}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"ID", "SYMBOL", "CATEGORY", "DATE", "SHARES",
		"PRICE", "MARKET VALUE", "PURCHASE VALUE"})

	for _, v := range transactions {
		var marketValue float64
		lastPrice := GetLatestPrice(v.Ticker)
		marketValue = lastPrice * float64(v.Shares)

		t.AppendRow(
			table.Row{v.ID, v.Ticker, IntToString(v.TransactionCategory),
				time.Unix(int64(v.TransactionDate), 0).Format("2006-1-2"),
				v.Shares, v.Price,
				fmt.Sprintf("%.2f \n", marketValue), v.PurchaseValue})
	}

	t.AppendSeparator()
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true
	t.Render()
}
