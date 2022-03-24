package transactions

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
)

func RemoveTransaction(id int) {
	// remove with id?
	removeTransaction, queryErr := database.DBConnection.Prepare(fmt.Sprintf("DELETE FROM transactions WHERE id='%d'", id))
	defer removeTransaction.Close()
	cobra.CheckErr(queryErr)

	_, removeErr := removeTransaction.Exec()
	cobra.CheckErr(removeErr)
}
