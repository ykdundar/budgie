package transactions

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"time"
)

func AddTransaction(ticker string, price float64, shares int, category string, date string, lastPrice float64) {
	val, timeErr := time.Parse("02.01.2006", date)
	cobra.CheckErr(timeErr)

	addTransaction, queryErr := database.DBConnection.Prepare(
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
		unixTime      = int(val.Unix())
		purchaseValue = price * float64(shares)
		marketValue   = lastPrice * float64(shares)
	)

	_, insertErr := addTransaction.Exec(ticker, price, shares, transactionCategory, unixTime, purchaseValue, marketValue)
	cobra.CheckErr(insertErr)
}
