package database

import "github.com/spf13/cobra"

func EnableForeignKeys() {
	// https://www.sqlite.org/foreignkeys.html
	enableForeignKeys, _ := database.Prepare("PRAGMA foreign_keys = ON;")

	defer enableForeignKeys.Close()

	_, fkErr := enableForeignKeys.Exec()
	
	cobra.CheckErr(fkErr)
}
