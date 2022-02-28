package database

import (
	"fmt"
	"github.com/spf13/cobra"
)

func DeletePortfolio(name string) {
	deletePortfolio, _ := database.Prepare(
		fmt.Sprintf("DELETE FROM portfolios WHERE name= '%s'", name),
	)
	defer deletePortfolio.Close()
	_, deleteErr := deletePortfolio.Exec()
	cobra.CheckErr(deleteErr)
	fmt.Printf("'%s' is deleted succesfully", name)
}
