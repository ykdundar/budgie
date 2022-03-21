package portfolios

import (
	"database/sql"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/internal/objects"
)

func FindPortfolio(name string) objects.Portfolio {
	record := database.Database.QueryRow("SELECT * FROM portfolios WHERE name=?", name)

	portfolio := objects.Portfolio{}

	scanErr := record.Scan(&portfolio.Id, &portfolio.Name, &portfolio.Currency, &portfolio.Active)

	if scanErr == sql.ErrNoRows {
		cobra.CheckErr(scanErr)
	}

	return portfolio
}
