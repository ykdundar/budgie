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

var addStockCmd = &cobra.Command{
	Use:   "add",
	Short: "add a stock to a given portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		addStock := stocks.AddStock(portfolio, ticker)
		if addStock == nil {
			fmt.Printf("'%s' is added succesfully\n", ticker)
		}
	},
	Example: `budgie stock add
	--portfolio "European Stocks"
	--ticker "MSFT"
`,
}

var removeStockCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a given stock",
	Run: func(cmd *cobra.Command, args []string) {
		removeStock := stocks.RemoveStock(portfolio, ticker)
		if removeStock == nil {
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
	stockCmd.AddCommand(addStockCmd, removeStockCmd)

	addStockCmd.Flags().StringVarP(&portfolio, "portfolio", "p", "", "Portfolio name (required)")
	addStockCmd.MarkFlagRequired("portfolio")
	addStockCmd.Flags().StringVarP(&ticker, "ticker", "s", "", "Company name (required)")
	addStockCmd.MarkFlagRequired("ticker")

	removeStockCmd.Flags().StringVarP(&portfolio, "portfolio", "p", "", "Portfolio name (required)")
	removeStockCmd.MarkFlagRequired("portfolio")
	removeStockCmd.Flags().StringVarP(&ticker, "ticker", "t", "", "Company name (required)")
	removeStockCmd.MarkFlagRequired("ticker")

}
