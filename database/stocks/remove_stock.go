package stocks

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/database/portfolios"
)

func RemoveStock(portfolioName string, ticker string) {
	portfolio := portfolios.FindPortfolio(portfolioName)

	removeStock, queryErr := database.DBConnection().Prepare(fmt.Sprintf("DELETE FROM stocks WHERE ticker='%s' AND portfolio_id=%d", ticker, portfolio.ID))
	cobra.CheckErr(queryErr)
	defer removeStock.Close()

	_, removeErr := removeStock.Exec()
	cobra.CheckErr(removeErr)
}
