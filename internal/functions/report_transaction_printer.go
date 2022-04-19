package functions

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ykdundar/budgie/api"
	"github.com/ykdundar/budgie/database/transactions"
	"os"
)

func ReportTransactionPrinter(transactionList map[string]transactions.TransactionSum, head string) {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{head, head, head, head, head}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"SYMBOL", "SHARES", "PURCHASE VALUE", "MARKET VALUE", "DIFFERENCE"})

	for k, v := range transactionList {

		currentPrice, _ := api.IntradayRequest([]string{k})
		var marketValue float64
		var diff float64

		if currentPrice.Data[0].Last == 0 {
			curPrice, _ := api.EndOfDayRequest([]string{k}, "latest")
			marketValue = curPrice.Data[0].Close * float64(v.Shares)
			diff = v.PurchaseValue - marketValue
			t.AppendRow(table.Row{k, v.Shares, v.PurchaseValue, fmt.Sprintf("%.2f \n", marketValue), diff})
			t.AppendSeparator()
		}
		marketValue = currentPrice.Data[0].Last * float64(v.Shares)
		diff = v.PurchaseValue - marketValue

		t.AppendRow(table.Row{k, v.Shares, v.PurchaseValue, marketValue, diff})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true
	t.Render()
}
