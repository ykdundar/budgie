package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func CreateTransactionsTable() {
	createTransactionsTable, queryErr := database.Prepare(
		"CREATE TABLE IF NOT EXISTS transactions (" +
			"ticker TEXT," +
			"transactions_date INTEGER," +
			"price INTEGER"+
			"shares INTEGER" +
			"purchaseValue "+
			"marketValue)",
		)
	defer createTransactionsTable.Close()
	cobra.CheckErr(queryErr)

	_, transactionsErr := createTransactionsTable.Exec()
	cobra.CheckErr(transactionsErr)
}
