package transactions

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/internal/objects"
)

func ListAllTransactions() []objects.Transaction {
	records, queryErr := database.DBConnection().Query(
		"SELECT id, ticker, transactions_date, price, shares, transaction_category, purchase_value FROM transactions",
	)
	cobra.CheckErr(queryErr)
	defer records.Close()

	transaction := objects.Transaction{}
	var transactions []objects.Transaction

	for records.Next() {
		scanErr := records.Scan(
			&transaction.ID,
			&transaction.Ticker,
			&transaction.TransactionDate,
			&transaction.Price,
			&transaction.Shares,
			&transaction.TransactionCategory,
			&transaction.PurchaseValue,
		)
		cobra.CheckErr(scanErr)
		transactions = append(transactions, transaction)
	}

	return transactions
}
