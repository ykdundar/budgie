package transactions

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
)

func CreateTransactionsTable() {
	createTransactionsTable, queryErr := database.DBConnection().Prepare(
		"CREATE TABLE IF NOT EXISTS transactions (" +
			"id INTEGER PRIMARY KEY," +
			"ticker TEXT," +
			"transactions_date INTEGER," +
			"price INTEGER," +
			"shares INTEGER," +
			"transaction_category INTEGER," +
			"purchase_value INTEGER)",
	)
	cobra.CheckErr(queryErr)
	defer createTransactionsTable.Close()

	_, transactionsErr := createTransactionsTable.Exec()
	cobra.CheckErr(transactionsErr)
}
