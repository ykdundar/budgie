package database

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func AddTransaction(ticker string, price int,  shares int, date string, category string) {
	val, timeErr := time.Parse("02.01.2006", date)
	cobra.CheckErr(timeErr)

	addTransaction, queryErr := database.Prepare(
		"INSERT INTO transactions" +
			"(ticker, price, shares, transaction_category, transactions_date, purchase_value, market_value)" +
			"VALUES (?,?,?,?,?,?,?)")
	defer addTransaction.Close()
	cobra.CheckErr(queryErr)

	transactionCategory := 1

	if category == "sell" {
		transactionCategory = 0
	}

	var (
		unixTime = int(val.Unix())
		purchaseValue = price * shares
	)
	_, insertErr := addTransaction.Exec(ticker, price, shares, transactionCategory, unixTime, purchaseValue, 0)
	cobra.CheckErr(insertErr)

	fmt.Printf("'%s' is added succesfully\n", ticker)
}
