package database

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/internal"
)

type TransactionSum struct {
	Shares        int
	PurchaseValue float64
}

func ReportRequest() {
	records, queryErr := database.Query("SELECT ticker, shares, purchase_value, transaction_category FROM transactions")
	cobra.CheckErr(queryErr)

	defer records.Close()

	dbRecord := internal.Transaction{}

	transactions := make(map[string]TransactionSum)

	for records.Next() {
		scanErr := records.Scan(&dbRecord.Ticker, &dbRecord.Shares, &dbRecord.PurchaseValue, &dbRecord.TransactionCategory)
		cobra.CheckErr(scanErr)

		if dbRecord.TransactionCategory == 1 {
			transactions[dbRecord.Ticker] = TransactionSum{
				Shares:        transactions[dbRecord.Ticker].Shares + dbRecord.Shares,
				PurchaseValue: transactions[dbRecord.Ticker].PurchaseValue + dbRecord.PurchaseValue,
			}
		} else {
			transactions[dbRecord.Ticker] = TransactionSum{
				Shares:        transactions[dbRecord.Ticker].Shares - dbRecord.Shares,
				PurchaseValue: transactions[dbRecord.Ticker].PurchaseValue - dbRecord.PurchaseValue,
			}
		}

	}
}
