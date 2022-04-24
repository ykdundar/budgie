package tokens

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
)

func AddToken(token string) {
	addToken, queryErr := database.DBConnection().Prepare("INSERT INTO tokens (token) VALUES (?)")
	cobra.CheckErr(queryErr)
	defer addToken.Close()

	_, insertErr := addToken.Exec(token)
	cobra.CheckErr(insertErr)
}
