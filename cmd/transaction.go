package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/api"
	"github.com/ykdundar/budgie/database"
	"time"
)

// flags
var price float64
var shares int
var date string
var id int
var day int
var month int
var year int

// transactionCmd represents the transaction command
var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "transaction command adds, updates, removes and reports a given stock by subcommands",
}

var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "Saves your stock buys",
	Run: func(cmd *cobra.Command, args []string) {
		req, reqErr := api.IntradayRequest(ticker)
		cobra.CheckErr(reqErr)

		lastPrice := req.Data[0].Last

		if lastPrice == 0 {
			eodReq, eodReqErr := api.EndOfDayRequest(ticker, "latest")
			lastPrice = eodReq.Data[0].Close
			cobra.CheckErr(eodReqErr)
		}

		database.AddTransaction(ticker, price, shares, cmd.Use, date, lastPrice)
	},
	Example: `budgie transaction buy
	--ticker "MSFT"
	--price 180
	--shares 20
	--date "19.01.2022"
`,
}

var sellCmd = &cobra.Command{
	Use:   "sell",
	Short: "Saves your stock sells",
	Run: func(cmd *cobra.Command, args []string) {
		req, reqErr := api.IntradayRequest(ticker)
		cobra.CheckErr(reqErr)

		lastPrice := req.Data[0].Last

		if lastPrice == 0 {
			eodReq, eodReqErr := api.EndOfDayRequest(ticker, "latest")
			lastPrice = eodReq.Data[0].Close
			cobra.CheckErr(eodReqErr)
		}

		database.AddTransaction(ticker, price, shares, cmd.Use, date, lastPrice)
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
		database.RemoveTransaction(id)
	},
}

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Reports transaction earnings/losses per stock",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		database.ReportRequest(cmd.Use, "")
	},
}

var dayCmd = &cobra.Command{
	Use:   "day",
	Short: "Reports transaction earnings/losses per stock for a given number of days",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		database.ReportRequest(cmd.Use, args[0])
	},
}

var monthCmd = &cobra.Command{
	Use:   "month",
	Short: "Reports transaction earnings/losses per stock for a given number of months",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		database.ReportRequest(cmd.Use, args[0])
	},
}

var yearCmd = &cobra.Command{
	Use:   "year",
	Short: "Reports transaction earnings/losses per stock for a given number of years",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		database.ReportRequest(cmd.Use, args[0])
	},
}

var listAllTransactionsCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your buying and selling records ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(database.ListAllTransactions())
	},
}

func init() {
	rootCmd.AddCommand(transactionCmd)
	transactionCmd.AddCommand(buyCmd, sellCmd, removeTransactionCmd, reportCmd, listAllTransactionsCmd)
	reportCmd.AddCommand(dayCmd, monthCmd, yearCmd)

	buyCmd.PersistentFlags().StringVarP(&ticker, "ticker", "t", "", "Company name (required)")
	buyCmd.MarkPersistentFlagRequired("ticker")
	buyCmd.PersistentFlags().Float64VarP(&price, "price", "p", 0, "Company price (required)")
	buyCmd.MarkPersistentFlagRequired("price")
	buyCmd.PersistentFlags().StringVarP(&date, "date", "d", time.Now().Format("02.01.2006"), "The date stock was bought (required)")
	buyCmd.MarkPersistentFlagRequired("date")
	buyCmd.PersistentFlags().IntVarP(&shares, "shares", "s", 0, "Number of shares (required)")
	buyCmd.MarkPersistentFlagRequired("shares")

	sellCmd.PersistentFlags().StringVarP(&ticker, "ticker", "t", "", "Company name (required)")
	sellCmd.MarkPersistentFlagRequired("ticker")
	sellCmd.PersistentFlags().Float64VarP(&price, "price", "p", 0, "Company price (required)")
	sellCmd.MarkPersistentFlagRequired("price")
	sellCmd.PersistentFlags().StringVarP(&date, "date", "d", time.Now().Format("02.01.2006"), "The date stock was sold (required)")
	sellCmd.MarkPersistentFlagRequired("date")
	sellCmd.PersistentFlags().IntVarP(&shares, "shares", "s", 0, "Number of shares (required)")
	sellCmd.MarkPersistentFlagRequired("shares")

	removeTransactionCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "Transaction ID (required)")
	removeTransactionCmd.MarkPersistentFlagRequired("id")
}
