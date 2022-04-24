package portfolios

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/internal/objects"
)

func SelectAllPortfolios() []objects.Portfolio {
	records, queryErr := database.DBConnection().Query("SELECT * FROM portfolios")
	cobra.CheckErr(queryErr)
	defer records.Close()

	portfolio := objects.Portfolio{}
	var portfolios []objects.Portfolio

	for records.Next() {
		scanErr := records.Scan(&portfolio.ID, &portfolio.Name)
		cobra.CheckErr(scanErr)

		portfolios = append(portfolios, portfolio)
	}

	return portfolios
}
