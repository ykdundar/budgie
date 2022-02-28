package database

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/internal"
)

func CreatePortfolio(name string, currency string, active bool) {
	activeValue := internal.ConvertBoolToInt(active)

	createPortfolio, _ := database.Prepare(
		"INSERT INTO portfolios (name, currency, active) VALUES (?, ?, ?)",
	)

	defer createPortfolio.Close()

	_, insertErr := createPortfolio.Exec(name, currency, activeValue)
	cobra.CheckErr(insertErr)

	fmt.Printf("'%s' is created succesfully\n", name)
}
