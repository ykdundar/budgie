package portfolios

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/internal/functions"
	"strings"
)

func UpdatePortfolio(name string, rename string, currency string, active bool) error {
	activeValue := functions.BoolConverter(active)

	var queryStr []string

	if rename != "" {
		queryStr = append(queryStr, fmt.Sprintf("name='%s'", rename))
	}

	if currency != "" {
		queryStr = append(queryStr, fmt.Sprintf("currency='%s'", currency))
	}

	queryStr = append(queryStr, fmt.Sprintf("active=%d", activeValue))

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
