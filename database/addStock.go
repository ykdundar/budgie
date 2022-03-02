package database

import (
	"github.com/spf13/cobra"
)

func AddStock(portfolioName string, ticker string) {
	portfolio := SelectPortfolio(portfolioName)

	addStock, queryErr := database.Prepare("INSERT INTO stocks (portfolio_id, ticker) VALUES (?, ?)")
	defer addStock.Close()
	cobra.CheckErr(queryErr)

	_, insertErr := addStock.Exec(portfolio.Id, ticker)
	cobra.CheckErr(insertErr)
}
