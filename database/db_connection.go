package database

import "database/sql"

// TODO: Save into home folder instead of repository
var DBConnection, _ = sql.Open("sqlite3", "./budgie.db")
