package database

import (
	"database/sql"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/internal"
)

func FindPortfolio(name string) internal.Portfolio {
	record := database.QueryRow("SELECT * FROM portfolios WHERE name=?", name)

	portfolio := internal.Portfolio{}

	scanErr := record.Scan(&portfolio.Id, &portfolio.Name, &portfolio.Currency, &portfolio.Active)

	if scanErr == sql.ErrNoRows {
		cobra.CheckErr(scanErr)
	}

	return portfolio
}
