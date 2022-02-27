package database

import "database/sql"

func CreateStocksTable() (*sql.Stmt, error) {
	createStocksTable, err := database.Prepare(
		"CREATE TABLE IF NOT EXISTS stocks (" +
			"stockId INTEGER PRIMARY KEY," +
			"name TEXT," +
			"ticker TEXT," +
			"buy_date INTEGER," +
			"sell_date INTEGER," +
			"buy_price INTEGER ," +
			"sell_price INTEGER ," +
			"shares INTEGER," +
			"portfolio_id INTEGER," +
			"FOREIGN KEY(portfolio_id)" +
			"REFERENCES portfolio(id))",
	)
	return createStocksTable, err
}
