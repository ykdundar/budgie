package stocks

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/database/portfolios"
)

func AddStock(portfolioName string, ticker string) {
	portfolio := portfolios.FindPortfolio(portfolioName)

	addStock, queryErr := database.DBConnection.Prepare("INSERT INTO stocks (portfolio_id, ticker) VALUES (?, ?)")
	defer addStock.Close()
	cobra.CheckErr(queryErr)

	_, insertErr := addStock.Exec(portfolio.Id, ticker)
	cobra.CheckErr(insertErr)
	fmt.Printf("'%s' is added succesfully\n", ticker)
}
