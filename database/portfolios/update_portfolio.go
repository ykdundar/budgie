package portfolios

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"strings"
)

func UpdatePortfolio(name string, rename string, currency string) error {

	var queryStr []string

	if rename != "" {
		queryStr = append(queryStr, fmt.Sprintf("name='%s'", rename))
	}

	if currency != "" {
		queryStr = append(queryStr, fmt.Sprintf("currency='%s'", currency))
	}


	updateSql := strings.Join(queryStr[:], ",")

	updatePortfolio, prepErr := database.DBConnection.Prepare(
		fmt.Sprintf("UPDATE portfolios SET %s WHERE name = '%s'", updateSql, name),
	)
	defer updatePortfolio.Close()
	cobra.CheckErr(prepErr)

	_, updateErr := updatePortfolio.Exec()

	cobra.CheckErr(updateErr)
	return updateErr
}
