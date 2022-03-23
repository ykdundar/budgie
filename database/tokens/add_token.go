package tokens

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
)

func AddToken(token string) error {
	addToken, queryErr := database.DBConnection.Prepare("INSERT INTO tokens (token) VALUES (?)")
	defer addToken.Close()
	cobra.CheckErr(queryErr)

	_, insertErr := addToken.Exec(token)
	cobra.CheckErr(insertErr)

	return insertErr
}
