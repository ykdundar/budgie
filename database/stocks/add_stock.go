package stocks

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/database/portfolios"
)

func AddStock(portfolioName string, ticker string) {
	portfolio := portfolios.FindPortfolio(portfolioName)

	addStock, queryErr := database.DBConnection().Prepare("INSERT INTO stocks (portfolio_id, ticker) VALUES (?, ?)")
	cobra.CheckErr(queryErr)
	defer addStock.Close()

	_, insertErr := addStock.Exec(portfolio.ID, ticker)
	cobra.CheckErr(insertErr)
}
