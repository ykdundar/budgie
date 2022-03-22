package portfolios

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/internal/functions"
)

func AddPortfolio(name string, currency string, active bool) {
	activeValue := functions.ConvertBoolToInt(active)

	createPortfolio, prepErr := database.DBConnection.Prepare(
		"INSERT INTO portfolios (name, currency, active) VALUES (?, ?, ?)",
	)
	defer createPortfolio.Close()
	cobra.CheckErr(prepErr)

	_, insertErr := createPortfolio.Exec(name, currency, activeValue)
	cobra.CheckErr(insertErr)

	fmt.Printf("'%s' is created succesfully\n", name)
}
