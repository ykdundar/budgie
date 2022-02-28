package database

import (
	"fmt"
)

func AddStock(name string, ticker string) {
	portfolio := SelectPortfolio(name)

	fmt.Println(portfolio)
	/*
		addStock, _ := dataBase.Prepare("INSERT INTO stocks (portfolio_id, ticker) VALUES (?, ?)")

		defer addStock.Close()

		_, insertErr := addStock.Exec(portfolio.Id, ticker)

		cobra.CheckErr(insertErr)

	*/
}
