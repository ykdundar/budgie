package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// EnableForeignKeys enables foreign key support in sqlite.
// For more details: https://www.sqlite.org/foreignkeys.
func EnableForeignKeys() {
	enableForeignKeys, queryErr := DBConnection().Prepare("PRAGMA foreign_keys = ON;")
	cobra.CheckErr(queryErr)
	defer enableForeignKeys.Close()

	_, fkErr := enableForeignKeys.Exec()
	cobra.CheckErr(fkErr)
}
