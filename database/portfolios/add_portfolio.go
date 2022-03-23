package portfolios

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
)

func AddPortfolio(name string, currency string) error{

	addPortfolio, prepErr := database.DBConnection.Prepare(
		"INSERT INTO portfolios (name, currency) VALUES (?, ?)",
	)
	defer addPortfolio.Close()
	cobra.CheckErr(prepErr)

	_, insertErr := addPortfolio.Exec(name, currency)
	cobra.CheckErr(insertErr)

	return insertErr
}
