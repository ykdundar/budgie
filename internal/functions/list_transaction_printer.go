package functions

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ykdundar/budgie/api"
	"github.com/ykdundar/budgie/internal/objects"
	"os"
	"time"
)

func ListTransactionPrinter(transactions []objects.Transaction, head string) {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{head, head, head, head, head, head, head, head}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"ID", "SYMBOL", "CATEGORY", "DATE", "SHARES",
		"PRICE", "MARKET VALUE", "PURCHASE VALUE"})

	for _, v := range transactions {
		currentPrice, _ := api.IntradayRequest([]string{v.Ticker})
		marketValue := float64(v.Shares) * currentPrice.Data[0].Last
		t.AppendRow(
			table.Row{v.Id, v.Ticker, v.TransactionCategory,
			time.Unix(int64(v.TransactionDate), 0).Format("2006-1-2"),
			v.Shares, v.Price,
			fmt.Sprintf("%.2f \n", marketValue), v.PurchaseValue})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true
	t.Render()
}
