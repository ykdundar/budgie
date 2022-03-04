package database

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/internal"
)

func DeleteStock(portfolio string, ticker string) {
	portfolioName := internal.Portfolio{}

	deleteStock, queryErr := database.Prepare(fmt.Sprintf("DELETE FROM stocks WHERE ticker='%s' AND portfolio_id=%s", ticker, portfolioName.Id))
	defer deleteStock.Close()
	cobra.CheckErr(queryErr)


	_, deleteErr := deleteStock.Exec()
	cobra.CheckErr(deleteErr)
	fmt.Printf("'%s' is deleted from '%s' succesfully", ticker, portfolio)
}
