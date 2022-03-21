package database

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RemoveTransaction(id int) {
	// remove with id?
	removeTransaction, queryErr := database.Prepare(fmt.Sprintf("DELETE FROM transactions WHERE id='%d'", id))
	defer removeTransaction.Close()
	cobra.CheckErr(queryErr)

	_, removeErr := removeTransaction.Exec()
	cobra.CheckErr(removeErr)
	fmt.Printf("'%d' is removed succesfully", id)
}
