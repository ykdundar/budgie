package database

import "database/sql"

func CreatePortfolioTable() (*sql.Stmt, error) {
	createPortfolios, err := database.Prepare(
		"CREATE TABLE IF NOT EXISTS portfolios (" +
			"id INTEGER PRIMARY KEY," +
			"name TEXT," +
			"currency TEXT," +
			"active INTEGER)",
	)

	return createPortfolios, err
}
