package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/api"
	"github.com/ykdundar/budgie/database/tokens"
	"github.com/ykdundar/budgie/database/transactions"
	"github.com/ykdundar/budgie/internal/functions"
	"time"
)

// transactionCmd represents the transaction command
var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "transaction command adds, updates, removes and reports a given stock by subcommands",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		transactions.CreateTransactionsTable()
		tokens.SetToken()
	},
}

var buyTransactionCmd = &cobra.Command{
	Use:   "buy",
	Short: "Saves a transaction (buy)",
	Run: func(cmd *cobra.Command, args []string) {
		req, reqErr := api.IntradayRequest([]string{ticker})
		cobra.CheckErr(reqErr)

		lastPrice := req.Data[0].Last

		if lastPrice == 0 {
			eodReq, eodReqErr := api.EndOfDayRequest(ticker, "latest")
			cobra.CheckErr(eodReqErr)

			lastPrice = eodReq.Data[0].Close
		}

		transactions.AddTransaction(ticker, price, shares, cmd.Use, date, lastPrice)
		fmt.Printf("'%s' is added succesfully\n", ticker)
	},
	Example: `budgie transaction buy
	--ticker="MSFT"
	--price=180
	--shares=20
	--date="19.01.2022"
`,
}

var sellTransactionCmd = &cobra.Command{
	Use:   "sell",
	Short: "Saves a transaction (sell)",
	Run: func(cmd *cobra.Command, args []string) {
		req, reqErr := api.IntradayRequest([]string{ticker})
		cobra.CheckErr(reqErr)

		lastPrice := req.Data[0].Last

		if lastPrice == 0 {
			eodReq, eodReqErr := api.EndOfDayRequest(ticker, "latest")
			cobra.CheckErr(eodReqErr)

			lastPrice = eodReq.Data[0].Close
		}
		transactions.AddTransaction(ticker, price, shares, cmd.Use, date, lastPrice)
		fmt.Printf("'%s' is added succesfully\n", ticker)
	},
	Example: `budgie transaction sell
	--ticker="MSFT"
	--price=180
	--shares=20
	--date="19.01.2022"
`,
}

var listAllTransactionsCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your transactions",
	Run: func(cmd *cobra.Command, args []string) {
		transactions := transactions.ListAllTransactions()
		functions.ListTransactionPrinter(transactions, "My Transactions")
	},
}

var removeTransactionCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a transaction",
	Run: func(cmd *cobra.Command, args []string) {
		transactions.RemoveTransaction(id)
		fmt.Printf("'%d' is removed succesfully\n", id)
	},
}

var reportTransactionsCmd = &cobra.Command{
	Use:   "report",
	Short: "Reports transaction earnings/losses per stock",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		transactions := transactions.ReportRequest(cmd.Use, "")
		functions.ReportTransactionPrinter(transactions, "My Earnings/Losses")
	},
}

var dayCmd = &cobra.Command{
	Use:   "day",
	Short: "Earnings/losses per stock for a given number of days",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		transactions := transactions.ReportRequest(cmd.Use, args[0])
		functions.ReportTransactionPrinter(transactions, "My Earnings/Losses")
	},
}

var monthCmd = &cobra.Command{
	Use:   "month",
	Short: "Earnings/losses per stock for a given number of months",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		transactions := transactions.ReportRequest(cmd.Use, args[0])
		functions.ReportTransactionPrinter(transactions, "My Earnings/Losses")

	},
}

var yearCmd = &cobra.Command{
	Use:   "year",
	Short: "Earnings/losses per stock for a given number of years",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		transactions := transactions.ReportRequest(cmd.Use, args[0])
		functions.ReportTransactionPrinter(transactions, "My Earnings/Losses")
	},
}

func init() {
	rootCmd.AddCommand(transactionCmd)
	transactionCmd.AddCommand(buyTransactionCmd, sellTransactionCmd, removeTransactionCmd, reportTransactionsCmd, listAllTransactionsCmd)
	reportTransactionsCmd.AddCommand(dayCmd, monthCmd, yearCmd)

	buyTransactionCmd.Flags().StringVarP(&ticker, "ticker", "t", "", "Company name (required)")
	buyTransactionCmd.MarkFlagRequired("ticker")
	buyTransactionCmd.Flags().Float64VarP(&price, "price", "p", 0, "Company price (required)")
	buyTransactionCmd.MarkFlagRequired("price")
	buyTransactionCmd.Flags().StringVarP(&date, "date", "d", time.Now().Format("02.01.2006"), "The date stock was bought")
	buyTransactionCmd.Flags().IntVarP(&shares, "shares", "s", 0, "Number of shares (required)")
	buyTransactionCmd.MarkFlagRequired("shares")

	sellTransactionCmd.Flags().StringVarP(&ticker, "ticker", "t", "", "Company name (required)")
	sellTransactionCmd.MarkFlagRequired("ticker")
	sellTransactionCmd.Flags().Float64VarP(&price, "price", "p", 0, "Company price (required)")
	sellTransactionCmd.MarkFlagRequired("price")
	sellTransactionCmd.Flags().StringVarP(&date, "date", "d", time.Now().Format("02.01.2006"), "The date stock was sold")
	sellTransactionCmd.Flags().IntVarP(&shares, "shares", "s", 0, "Number of shares (required)")
	sellTransactionCmd.MarkFlagRequired("shares")

	removeTransactionCmd.Flags().IntVarP(&id, "id", "i", 0, "Transaction ID (required)")
	removeTransactionCmd.MarkFlagRequired("id")
}
