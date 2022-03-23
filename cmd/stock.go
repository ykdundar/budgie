package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database/stocks"
)

// stockCmd represents the stock command
var stockCmd = &cobra.Command{
	Use:   "stock",
	Short: "stock command adds, removes and reports a given stock by subcommands",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		stocks.CreateStocksTable()
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a stock to a given portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		addStock := stocks.AddStock(portfolio, ticker)
		if addStock == nil{
			fmt.Printf("'%s' is added succesfully\n", ticker)
		}
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
		removeStock := stocks.RemoveStock(portfolio, ticker)
		if removeStock == nil{
			fmt.Printf("'%s' is removed from '%s' succesfully", ticker, portfolio)
		}
	},
	Example: `budgie stock remove
	--portfolio "European Stocks"
	--ticker "MSFT"
`,
}

func init() {
	rootCmd.AddCommand(stockCmd)
	stockCmd.AddCommand(addCmd, removeCmd)

	addCmd.Flags().StringVarP(&portfolio, "portfolio", "p", "", "Portfolio name (required)")
	addCmd.MarkFlagRequired("portfolio")
	addCmd.Flags().StringVarP(&ticker, "ticker", "s", "", "Company name (required)")
	addCmd.MarkFlagRequired("ticker")

	removeCmd.Flags().StringVarP(&portfolio, "portfolio", "p", "", "Portfolio name (required)")
	removeCmd.MarkFlagRequired("portfolio")
	removeCmd.Flags().StringVarP(&ticker, "ticker", "t", "", "Company name (required)")
	removeCmd.MarkFlagRequired("ticker")

}
