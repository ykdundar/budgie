package portfolios

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
)

func CreatePortfoliosTable() {
	createPortfolios, queryErr := database.DBConnection().Prepare(
		"CREATE TABLE IF NOT EXISTS portfolios (" +
			"id INTEGER PRIMARY KEY," +
			"name TEXT UNIQUE)",
	)
	defer createPortfolios.Close()
	cobra.CheckErr(queryErr)

	_, portfolioErr := createPortfolios.Exec()
	cobra.CheckErr(portfolioErr)
}
