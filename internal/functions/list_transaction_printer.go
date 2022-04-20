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
		var marketValue float64

		if currentPrice.Data[0].Last == 0{
			curPrice, _ := api.EndOfDayRequest([]string{v.Ticker}, "latest")
			marketValue = curPrice.Data[0].Close * float64(v.Shares)
		} else {
			marketValue = currentPrice.Data[0].Last * float64(v.Shares)
		}

		t.AppendRow(
			table.Row{v.Id, v.Ticker, IntToString(v.TransactionCategory) ,
				time.Unix(int64(v.TransactionDate), 0).Format("2006-1-2"),
				v.Shares, v.Price,
				fmt.Sprintf("%.2f \n", marketValue), v.PurchaseValue})
	}

	t.AppendSeparator()
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true
	t.Render()
}
