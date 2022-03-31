package portfolios

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/internal/objects"
)

func SelectAllPortfolios() []objects.Portfolio {
	records, queryErr := database.DBConnection().Query("SELECT * FROM portfolios")
	defer records.Close()
	cobra.CheckErr(queryErr)

	portfolio := objects.Portfolio{}
	var portfolios []objects.Portfolio

	for records.Next() {
		scanErr := records.Scan(&portfolio.Id, &portfolio.Name, &portfolio.Currency)
		cobra.CheckErr(scanErr)

		portfolios = append(portfolios, portfolio)
	}

	return portfolios
}
