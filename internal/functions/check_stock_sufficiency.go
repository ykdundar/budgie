package functions

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/internal/objects"
)

func CheckStockSufficiency(shares int, ticker string) {
	transaction := objects.Transaction{}
	record := database.DBConnection().QueryRow(fmt.Sprintf("SELECT sum(shares) FROM transactions WHERE transaction_category=1 AND ticker='%s'", ticker)).Scan(&transaction.Shares)

	cobra.CheckErr(record)

	if shares > transaction.Shares {
		cobra.CheckErr("You don't have enough stock to sell")
	}
}
