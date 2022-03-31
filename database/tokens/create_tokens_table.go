package tokens

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
)

// CreateTokensTable creates a table to store API tokens
func CreateTokensTable() {
	createTokensTable, queryErr := database.DBConnection().Prepare(
		"CREATE TABLE IF NOT EXISTS tokens (" +
			"id INTEGER PRIMARY KEY," +
			"token TEXT)",
	)
	defer createTokensTable.Close()
	cobra.CheckErr(queryErr)

	_, tokensErr := createTokensTable.Exec()
	cobra.CheckErr(tokensErr)
}
