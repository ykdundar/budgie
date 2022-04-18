package functions

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ykdundar/budgie/api"
	"github.com/ykdundar/budgie/database/transactions"
	"os"
)

func ReportTransactionPrinter(transactions map[string]transactions.TransactionSum, head string) {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{head, head, head, head, head}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"SYMBOL", "SHARES", "PURCHASE VALUE", "MARKET VALUE", "DIFFERENCE"})

	for k, v := range transactions {
		currentPrice, _ := api.IntradayRequest([]string{k})
		marketValue := currentPrice.Data[0].Last * float64(v.Shares)
		diff := marketValue - v.PurchaseValue

		t.AppendRow(table.Row{k, v.Shares, v.PurchaseValue, marketValue, diff})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true
	t.Render()
}
