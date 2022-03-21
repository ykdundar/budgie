package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func CreateStocksTable() {
	createStocksTable, queryErr := database.Prepare(
		"CREATE TABLE IF NOT EXISTS stocks (" +
			"stockId INTEGER PRIMARY KEY," +
			"ticker TEXT," +
			"portfolio_id INTEGER," +
			"UNIQUE(ticker, portfolio_id)," +
			"FOREIGN KEY(portfolio_id)" +
			"REFERENCES portfolios(id))",
	)
	defer createStocksTable.Close()
	cobra.CheckErr(queryErr)

	_, stockErr := createStocksTable.Exec()
	cobra.CheckErr(stockErr)
}
