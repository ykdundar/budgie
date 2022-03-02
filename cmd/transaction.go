
package cmd

import (

	"github.com/spf13/cobra"
)
//flags
var price int
var shares int
var date int

// transactionCmd represents the transaction command
var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "transaction command adds, updates, removes and reports a given stock by subcommands",
}

var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "Add the stock you bought to transactions table",
	Run: func(cmd *cobra.Command, args []string) {
	},
	Example: `budgie transaction buy
	--ticker "MSFT"
	--price "180"
	--shares "20"
	--date "19.01.2022"
`,
}

func init() {
	rootCmd.AddCommand(transactionCmd)
	transactionCmd.AddCommand(buyCmd)

	buyCmd.Flags().StringVarP(&ticker, "ticker", "s", "", "Company name (required)")
	buyCmd.MarkPersistentFlagRequired("ticker")
	buyCmd.Flags().IntVarP(&price, "price", "p", 0, "Company price (required)")
	buyCmd.MarkPersistentFlagRequired("price")
	buyCmd.Flags().IntVarP(&date, "date", "d", 0, "The date stock was bought (required)")
	buyCmd.MarkPersistentFlagRequired("date")
	buyCmd.Flags().IntVarP(&shares, "shares", "h", 0, "Number of shares (required)")
	buyCmd.MarkPersistentFlagRequired("shares")
	}
