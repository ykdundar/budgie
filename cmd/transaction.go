package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/api"
	"github.com/ykdundar/budgie/database"
	"time"
)

//flags
var price float64
var shares int
var date string

// transactionCmd represents the transaction command
var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "transaction command adds, updates, removes and reports a given stock by subcommands",
}

var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "Add the stock you bought to transactions table",
	Run: func(cmd *cobra.Command, args []string) {
		req, reqErr := api.IntradayRequest(ticker)
		cobra.CheckErr(reqErr)

		lastPrice := req.Data[0].Last

		if lastPrice == 0 {
			eodReq, eodReqErr := api.EndOfDayRequest(ticker)
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
	Short: "Add the stock you bought to transactions table",
	Run: func(cmd *cobra.Command, args []string) {
		req, reqErr := api.IntradayRequest(ticker)
		cobra.CheckErr(reqErr)

		lastPrice := req.Data[0].Last

		if lastPrice == 0 {
			eodReq, eodReqErr := api.EndOfDayRequest(ticker)
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

func init() {
	rootCmd.AddCommand(transactionCmd)
	transactionCmd.AddCommand(buyCmd, sellCmd)

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
}
