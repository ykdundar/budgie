package database

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/internal"
)

func SelectAllPortfolios() {
	records, queryErr := database.Query("SELECT * FROM portfolios")
	cobra.CheckErr(queryErr)

	defer records.Close()

	portfolio := internal.Portfolio{}
	var portfolios []internal.Portfolio

	for records.Next() {
		scanErr := records.Scan(&portfolio.Id, &portfolio.Name, &portfolio.Currency, &portfolio.Active)
		cobra.CheckErr(scanErr)

		portfolios = append(portfolios, portfolio)
	}

	for _, v := range portfolios {
		fmt.Println("Id: ", v.Id, "Name: ", v.Name, "Currency: ", v.Currency, "Active: ", v.Active)
	}
}
