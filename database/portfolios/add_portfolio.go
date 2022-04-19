package portfolios

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
)

func AddPortfolio(name string) {
	addPortfolio, prepErr := database.DBConnection().Prepare(
		"INSERT INTO portfolios (name) VALUES (?)",
	)
	defer addPortfolio.Close()
	cobra.CheckErr(prepErr)

	_, insertErr := addPortfolio.Exec(name)
	cobra.CheckErr(insertErr)
}
