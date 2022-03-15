package database

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/internal"
)

func ListAllTransactions() []internal.Transaction {
	records, queryErr := database.Query("SELECT id, ticker, price,  shares, transaction_category FROM transactions ")
	defer records.Close()
	cobra.CheckErr(queryErr)

	transactions := internal.Transaction{}
	var transaction []internal.Transaction

	for records.Next() {
		scanErr := records.Scan(&transactions.Id, &transactions.Ticker, &transactions.Price, &transactions.Shares, &transactions.TransactionCategory)
		cobra.CheckErr(scanErr)
		transaction = append(transaction, transactions)
	}
	return transaction

}
