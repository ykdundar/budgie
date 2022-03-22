package portfolios

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/internal/objects"
)

func SelectAllPortfolios() {
	records, queryErr := database.DBConnection.Query("SELECT * FROM portfolios")
	defer records.Close()
	cobra.CheckErr(queryErr)

	portfolio := objects.Portfolio{}
	var portfolios []objects.Portfolio

	for records.Next() {
		scanErr := records.Scan(&portfolio.Id, &portfolio.Name, &portfolio.Currency, &portfolio.Active)
		cobra.CheckErr(scanErr)

		portfolios = append(portfolios, portfolio)
	}

	for _, v := range portfolios {
		fmt.Println("Id: ", v.Id, "Name: ", v.Name, "Currency: ", v.Currency, "Active: ", v.Active)
	}
}
