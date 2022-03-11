package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
)

// flags
var (
	portfolio string
	ticker    string
)

// stockCmd represents the stock command
var stockCmd = &cobra.Command{
	Use:   "stock",
	Short: "stock command adds, removes and reports a given stock by subcommands",
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a stock to a given portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		database.AddStock(portfolio, ticker)
	},
	Example: `budgie stock add
	--portfolio "European Stocks"
	--ticker "MSFT"
`,
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a given stock",
	Run: func(cmd *cobra.Command, args []string) {
		database.RemoveStock(portfolio, ticker)
	},
	Example: `budgie stock remove
	--portfolio "European Stocks"
	--ticker "MSFT"
`,
}
	
func init() {
	rootCmd.AddCommand(stockCmd)
	stockCmd.AddCommand(addCmd, removeCmd)

	addCmd.PersistentFlags().StringVarP(&portfolio, "portfolio", "p", "", "Portfolio name (required)")
	addCmd.MarkPersistentFlagRequired("portfolio")
	addCmd.PersistentFlags().StringVarP(&ticker, "ticker", "s", "", "Company name (required)")
	addCmd.MarkPersistentFlagRequired("ticker")

	removeCmd.PersistentFlags().StringVarP(&portfolio, "portfolio", "p", "", "Portfolio name (required)")
	removeCmd.MarkPersistentFlagRequired("portfolio")
	removeCmd.PersistentFlags().StringVarP(&ticker, "ticker", "t", "", "Company name (required)")
	removeCmd.MarkPersistentFlagRequired("ticker")

}
