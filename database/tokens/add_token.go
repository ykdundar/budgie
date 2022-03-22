package tokens

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
)

func AddToken(token string) {
	addToken, queryErr := database.DBConnection.Prepare("INSERT INTO tokens (token) VALUES (?)")
	defer addToken.Close()
	cobra.CheckErr(queryErr)

	_, insertErr := addToken.Exec(token)
	cobra.CheckErr(insertErr)
	fmt.Printf("'%s' is added succesfully\n", token)
}
