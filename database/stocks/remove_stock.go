package stocks

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/database/portfolios"
	"github.com/ykdundar/budgie/internal/objects"
)

func RemoveStock(portfolio string, ticker string) error{

	portfolioStr := objects.Portfolio{}
	portfolioStr = portfolios.FindPortfolio(portfolio)

	removeStock, queryErr := database.DBConnection.Prepare(fmt.Sprintf("DELETE FROM stocks WHERE ticker='%s' AND portfolio_id=%d", ticker, portfolioStr.Id))
	defer removeStock.Close()
	cobra.CheckErr(queryErr)

	_, removeErr := removeStock.Exec()
	cobra.CheckErr(removeErr)

	return removeErr

}
