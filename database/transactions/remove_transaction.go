package transactions

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
)

func RemoveTransaction(id int) {
	removeTransaction, queryErr := database.DBConnection().Prepare(fmt.Sprintf("DELETE FROM transactions WHERE id='%d'", id))
	cobra.CheckErr(queryErr)
	defer removeTransaction.Close()

	_, removeErr := removeTransaction.Exec()
	cobra.CheckErr(removeErr)
}
