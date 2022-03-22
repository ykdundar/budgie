package database

import (
	"database/sql"
	"os"
	"path/filepath"
)

// TODO: Save into home folder instead of repository

var homeDirectory, _ = os.UserHomeDir()
var DBConnection, _ = sql.Open("sqlite3", filepath.Join(homeDirectory, ".budgie.db"))
