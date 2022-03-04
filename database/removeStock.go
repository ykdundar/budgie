package database

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/internal"
)

func RemoveStock(portfolio string, ticker string) {
	portfolioStr := internal.Portfolio{}
	portfolioStr = SelectPortfolio(portfolio)

	removeStock, queryErr := database.Prepare(fmt.Sprintf("DELETE FROM stocks WHERE ticker='%s' AND portfolio_id=%d", ticker, portfolioStr.Id))
	defer removeStock.Close()
	cobra.CheckErr(queryErr)

	_, removeErr := removeStock.Exec()
	cobra.CheckErr(removeErr)
	fmt.Printf("'%s' is removed from '%s' succesfully", ticker, portfolio)
}
