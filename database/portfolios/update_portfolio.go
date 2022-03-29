package portfolios

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"strings"
)

func UpdatePortfolio(name string, rename string, currency string) error {
	updatePortfolio, prepErr := database.DBConnection.Prepare(buildQuery(name, rename, currency))
	defer updatePortfolio.Close()
	cobra.CheckErr(prepErr)

	_, updateErr := updatePortfolio.Exec()
	cobra.CheckErr(updateErr)

	return updateErr
}

func buildQuery(name string, rename string, currency string) string {
	var querySlc []string

	if rename != "" {
		querySlc = append(querySlc, fmt.Sprintf("name='%s'", rename))
	}

	if currency != "" {
		querySlc = append(querySlc, fmt.Sprintf("currency='%s'", currency))
	}

	return fmt.Sprintf("UPDATE portfolios SET %s WHERE name = '%s'", strings.Join(querySlc[:], ","), name)
}
