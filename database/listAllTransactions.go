package database

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/internal"
)

func ListAllTransactions() {
	records, queryErr := database.Query("SELECT id, ticker, price,  shares, transaction_category FROM transactions ")
	defer records.Close()
	cobra.CheckErr(queryErr)

	transaction := internal.Transaction{}
	var transactions []internal.Transaction

	for records.Next() {
		scanErr := records.Scan(&transaction.Id, &transaction.Ticker, &transaction.Price, &transaction.Shares, &transaction.TransactionCategory)
		cobra.CheckErr(scanErr)
		transactions = append(transactions, transaction)
	}
	for _, v := range transactions{
		fmt.Println("Id: ", v.Id, "Ticker: ", v.Ticker, "Price: ", v.Price, "Shares: ", v.Shares, "Category: ", v.TransactionCategory)
	}
}
