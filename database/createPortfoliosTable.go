package database

import (
	"github.com/spf13/cobra"
)

func CreatePortfolioTable() {
	createPortfolios, _ := database.Prepare(
		"CREATE TABLE IF NOT EXISTS portfolios (" +
			"id INTEGER PRIMARY KEY," +
			"name TEXT," +
			"currency TEXT," +
			"active INTEGER)",
	)

	createPortfolios.Close()

	_, portfolioErr := createPortfolios.Exec()

	cobra.CheckErr(portfolioErr)
}
