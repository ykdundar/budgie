package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/api"
	"github.com/ykdundar/budgie/database/tokens"
	"github.com/ykdundar/budgie/database/transactions"
	"time"
)

// transactionCmd represents the transaction command
var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "transaction command adds, updates, removes and reports a given stock by subcommands",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		transactions.CreateTransactionsTable()
	},
}

var buyTransactionCmd = &cobra.Command{
	Use:   "buy",
	Short: "Saves your stock buys",
	PreRun: func(cmd *cobra.Command, args []string) {
		tokens.CheckToken()
	},
	Run: func(cmd *cobra.Command, args []string) {
		req, reqErr := api.IntradayRequest(ticker)
		cobra.CheckErr(reqErr)

		lastPrice := req.Data[0].Last

		if lastPrice == 0 {
			eodReq, eodReqErr := api.EndOfDayRequest(ticker, "latest")
			lastPrice = eodReq.Data[0].Close
			cobra.CheckErr(eodReqErr)
		}

		addTransaction := transactions.AddTransaction(ticker, price, shares, cmd.Use, date, lastPrice)
		if addTransaction == nil {
			fmt.Printf("'%s' is added succesfully\n", ticker)
		}
	},
	Example: `budgie transaction buy
	--ticker "MSFT"
	--price 180
	--shares 20
	--date "19.01.2022"
`,
}

var sellTransactionCmd = &cobra.Command{
	Use:   "sell",
	Short: "Saves your stock sells",
	PreRun: func(cmd *cobra.Command, args []string) {
		tokens.CheckToken()
	},
	Run: func(cmd *cobra.Command, args []string) {
		req, reqErr := api.IntradayRequest(ticker)
		cobra.CheckErr(reqErr)

		lastPrice := req.Data[0].Last

		if lastPrice == 0 {
			eodReq, eodReqErr := api.EndOfDayRequest(ticker, "latest")
			lastPrice = eodReq.Data[0].Close
			cobra.CheckErr(eodReqErr)
		}
		addTransaction := transactions.AddTransaction(ticker, price, shares, cmd.Use, date, lastPrice)
		if addTransaction == nil {
			fmt.Printf("'%s' is added succesfully\n", ticker)
		}
	},
	Example: `budgie transaction sell
	--ticker "MSFT"
	--price 180
	--shares 20
	--date "19.01.2022"
`,
}

var removeTransactionCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes your stock purchases",
	Run: func(cmd *cobra.Command, args []string) {
		removeTransaction := transactions.RemoveTransaction(id)
		if removeTransaction == nil {
			fmt.Printf("'%d' is removed succesfully", id)
		}
	},
}

var reportTransactionsCmd = &cobra.Command{
	Use:   "report",
	Short: "Reports transaction earnings/losses per stock",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		transactions.ReportRequest(cmd.Use, "")
	},
}

var dayCmd = &cobra.Command{
	Use:   "day",
	Short: "Reports transaction earnings/losses per stock for a given number of days",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		transactions.ReportRequest(cmd.Use, args[0])
	},
}

var monthCmd = &cobra.Command{
	Use:   "month",
	Short: "Reports transaction earnings/losses per stock for a given number of months",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		transactions.ReportRequest(cmd.Use, args[0])
	},
}

var yearCmd = &cobra.Command{
	Use:   "year",
	Short: "Reports transaction earnings/losses per stock for a given number of years",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		transactions.ReportRequest(cmd.Use, args[0])
	},
}

var listAllTransactionsCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your buying and selling records ",
	Run: func(cmd *cobra.Command, args []string) {
		transactions.ListAllTransactions()
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
	buyTransactionCmd.Flags().StringVarP(&date, "date", "d", time.Now().Format("02.01.2006"), "The date stock was bought (required)")
	buyTransactionCmd.MarkFlagRequired("date")
	buyTransactionCmd.Flags().IntVarP(&shares, "shares", "s", 0, "Number of shares (required)")
	buyTransactionCmd.MarkFlagRequired("shares")

	sellTransactionCmd.Flags().StringVarP(&ticker, "ticker", "t", "", "Company name (required)")
	sellTransactionCmd.MarkFlagRequired("ticker")
	sellTransactionCmd.Flags().Float64VarP(&price, "price", "p", 0, "Company price (required)")
	sellTransactionCmd.MarkFlagRequired("price")
	sellTransactionCmd.Flags().StringVarP(&date, "date", "d", time.Now().Format("02.01.2006"), "The date stock was sold (required)")
	sellTransactionCmd.MarkFlagRequired("date")
	sellTransactionCmd.Flags().IntVarP(&shares, "shares", "s", 0, "Number of shares (required)")
	sellTransactionCmd.MarkFlagRequired("shares")

	removeTransactionCmd.Flags().IntVarP(&id, "id", "i", 0, "Transaction ID (required)")
	removeTransactionCmd.MarkFlagRequired("id")
}
