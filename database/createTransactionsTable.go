package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func CreateTransactionsTable() {
	createTransactionsTable, queryErr := database.Prepare(
		"CREATE TABLE IF NOT EXISTS transactions (" +
			"id INTEGER PRIMARY KEY,"+
			"ticker TEXT," +
			"transactions_date INTEGER," +
			"price INTEGER," +
			"shares INTEGER," +
			"transaction_category INTEGER," +
			"purchase_value INTEGER," +
			"market_value INTEGER)",
		)
	defer createTransactionsTable.Close()
	cobra.CheckErr(queryErr)

	_, transactionsErr := createTransactionsTable.Exec()
	cobra.CheckErr(transactionsErr)
}
