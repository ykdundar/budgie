package stocks

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/database/portfolios"
)

func AddStock(portfolioName string, ticker string) error{
	portfolio := portfolios.FindPortfolio(portfolioName)

	addStock, queryErr := database.DBConnection.Prepare("INSERT INTO stocks (portfolio_id, ticker) VALUES (?, ?)")
	defer addStock.Close()
	cobra.CheckErr(queryErr)

	_, insertErr := addStock.Exec(portfolio.Id, ticker)
	cobra.CheckErr(insertErr)

	return insertErr
}
