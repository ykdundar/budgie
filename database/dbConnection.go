package database

import "database/sql"

var database, _ = sql.Open("sqlite3", "./budgie.db")
