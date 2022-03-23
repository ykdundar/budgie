package portfolios

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/internal/functions"
)

func AddPortfolio(name string, currency string, active bool) error{
	activeValue := functions.BoolConverter(active)

	addPortfolio, prepErr := database.DBConnection.Prepare(
		"INSERT INTO portfolios (name, currency, active) VALUES (?, ?, ?)",
	)
	defer addPortfolio.Close()
	cobra.CheckErr(prepErr)

	_, insertErr := addPortfolio.Exec(name, currency, activeValue)
	cobra.CheckErr(insertErr)

	return insertErr
}
