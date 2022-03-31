package database

import (
	"database/sql"
	"os"
	"path/filepath"
)

func DBConnection() *sql.DB {
	homeDirectory, _ := os.UserHomeDir()
	dbConnection, _ := sql.Open("sqlite3", filepath.Join(homeDirectory, ".budgie.db"))

	return dbConnection
}
