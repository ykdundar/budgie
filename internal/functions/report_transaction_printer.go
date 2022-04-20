package functions

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ykdundar/budgie/database/transactions"
	"os"
)

func ReportTransactionPrinter(transactionList []transactions.TransactionSum, head string) {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{head, head, head, head, head}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"SYMBOL", "SHARES", "PURCHASE VALUE", "MARKET VALUE", "DIFFERENCE"})

	for _, v := range transactionList {
		lastPrice := GetLatestPrice(v.Ticker)

		var marketValue float64
		var diff float64

		marketValue = lastPrice * float64(v.Shares)
		diff = v.PurchaseValue - marketValue

		t.AppendRow(table.Row{v.Ticker, v.Shares, v.PurchaseValue, fmt.Sprintf("%.2f \n", marketValue), diff})
	}

	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true
	t.Render()
}
