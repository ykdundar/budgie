package database

import (
	"github.com/spf13/cobra"
)

func CreateStocksTable() {
	createStocksTable, _ := database.Prepare(
		"CREATE TABLE IF NOT EXISTS stocks (" +
			"stockId INTEGER PRIMARY KEY," +
			"name TEXT," +
			"ticker TEXT," +
			"buy_date INTEGER," +
			"sell_date INTEGER," +
			"buy_price INTEGER ," +
			"sell_price INTEGER ," +
			"shares INTEGER," +
			"portfolio_id INTEGER," +
			"FOREIGN KEY(portfolio_id)" +
			"REFERENCES portfolio(id))",
	)

	defer createStocksTable.Close()

	_, stockErr := createStocksTable.Exec()

	cobra.CheckErr(stockErr)
}
