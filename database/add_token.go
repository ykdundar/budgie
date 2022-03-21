package database

import (
	"fmt"
	"github.com/spf13/cobra"
)

func AddToken(token string) {
	addToken, queryErr := Database.Prepare("INSERT INTO tokens (token) VALUES (?)")
	defer addToken.Close()
	cobra.CheckErr(queryErr)

	_, insertErr := addToken.Exec(token)
	cobra.CheckErr(insertErr)
	fmt.Printf("'%s' is added succesfully\n", token)
}
