package transactions

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/internal/objects"
)

type TransactionSum struct {
	Ticker        string
	Shares        int
	PurchaseValue float64
}

func ReportRequest(command string, commandValue string) []TransactionSum {
	var commandValueInt int
	var convErr error

	if command != "report" {
		commandValueInt, convErr = strconv.Atoi(commandValue)
		cobra.CheckErr(convErr)
	}

	var pastTime int64

	switch {
	case command == "day":
		pastTime = time.Now().AddDate(0, 0, -commandValueInt).Unix()
	case command == "month":
		pastTime = time.Now().AddDate(0, -commandValueInt, 0).Unix()
	case command == "year":
		pastTime = time.Now().AddDate(-commandValueInt, 0, 0).Unix()
	}

	var baseQuery = fmt.Sprintf(
		"SELECT ticker, sum(shares), sum(purchase_value), transaction_category FROM transactions WHERE transactions_date > %d GROUP BY transaction_category, ticker", pastTime,
	)

	records, queryErr := database.DBConnection().Query(baseQuery)
	cobra.CheckErr(queryErr)
	defer records.Close()

	dbRecord := objects.Transaction{}

	var transactions []TransactionSum

	for records.Next() {
		scanErr := records.Scan(&dbRecord.Ticker, &dbRecord.Shares, &dbRecord.PurchaseValue, &dbRecord.TransactionCategory)
		cobra.CheckErr(scanErr)

		transactions = append(transactions, TransactionSum{
			Ticker:        dbRecord.Ticker,
			Shares:        dbRecord.Shares,
			PurchaseValue: dbRecord.PurchaseValue,
		})
	}

	return transactions
}
