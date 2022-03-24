package database

import (
	"database/sql"
	"os"
	"path/filepath"
)

var homeDirectory, _ = os.UserHomeDir()
var DBConnection, _ = sql.Open("sqlite3", filepath.Join(homeDirectory, ".budgie.db"))
