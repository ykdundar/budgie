package database

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/internal/functions"
	"strings"
)

func UpdatePortfolio(name string, rename string, currency string, active bool) {
	activeValue := functions.ConvertBoolToInt(active)

	var queryStr []string

	if rename != "" {
		queryStr = append(queryStr, fmt.Sprintf("name='%s'", rename))
	}

	if currency != "" {
		queryStr = append(queryStr, fmt.Sprintf("currency='%s'", currency))
	}

	queryStr = append(queryStr, fmt.Sprintf("active=%d", activeValue))

	updateSql := strings.Join(queryStr[:], ",")

	updatePortfolio, _ := database.Prepare(
		fmt.Sprintf("UPDATE portfolios SET %s WHERE name = '%s'", updateSql, name),
	)

	defer updatePortfolio.Close()

	_, updateErr := updatePortfolio.Exec()

	cobra.CheckErr(updateErr)

	fmt.Printf("'%s' is updated succesfully\n", name)
}
