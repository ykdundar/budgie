package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// EnableForeignKeys enables foreign key support in sqlite
// For more details: https://www.sqlite.org/foreignkeys.html
func EnableForeignKeys() {
	enableForeignKeys, queryErr := Database.Prepare("PRAGMA foreign_keys = ON;")
	defer enableForeignKeys.Close()
	cobra.CheckErr(queryErr)

	_, fkErr := enableForeignKeys.Exec()
	cobra.CheckErr(fkErr)
}
