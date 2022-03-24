package portfolios

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
)

func DeletePortfolio(name string) {
	deletePortfolio, queryErr := database.DBConnection.Prepare(
		fmt.Sprintf("DELETE FROM portfolios WHERE name= '%s'", name),
	)
	defer deletePortfolio.Close()
	cobra.CheckErr(queryErr)

	_, deleteErr := deletePortfolio.Exec()
	cobra.CheckErr(deleteErr)
}
