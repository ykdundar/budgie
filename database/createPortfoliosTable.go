package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func CreatePortfolioTable() {
	createPortfolios, queryErr := database.Prepare(
		"CREATE TABLE IF NOT EXISTS portfolios (" +
			"id INTEGER PRIMARY KEY," +
			"name TEXT UNIQUE," +
			"currency TEXT," +
			"active INTEGER)",
	)

	defer createPortfolios.Close()
	cobra.CheckErr(queryErr)

	_, portfolioErr := createPortfolios.Exec()

	cobra.CheckErr(portfolioErr)
}
