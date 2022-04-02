package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/api"
	"github.com/ykdundar/budgie/database/stocks"
	"github.com/ykdundar/budgie/database/tokens"
	"github.com/ykdundar/budgie/internal/functions"
)

// stockCmd represents the stock command
var stockCmd = &cobra.Command{
	Use:   "stock",
	Short: "stock command is allows you to watch stocks that you are interested in",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		stocks.CreateStocksTable()
	},
}

var addStockCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a stock to a given portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		stocks.AddStock(portfolio, ticker)
		fmt.Printf("'%s' is added succesfully\n", ticker)
	},
	Example: `budgie stock add
	--portfolio "European Stocks"
	--ticker "MSFT"
`,
}

var removeStockCmd = &cobra.Command{
	Use:   "remove",
	Short: "removes a stock from a given portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		stocks.RemoveStock(portfolio, ticker)
		fmt.Printf("'%s' is removed from '%s' succesfully", ticker, portfolio)
	},
	Example: `budgie stock remove
	--portfolio "European Stocks"
	--ticker "MSFT"
`,
}

var showStockCmd = &cobra.Command{
	Use:   "show",
	Short: "shows a stock or comma seperated stocks",
	PreRun: func(cmd *cobra.Command, args []string) {
		tokens.SetToken()
	},
	Run: func(cmd *cobra.Command, args []string) {
		requests, intradayErr := api.IntradayRequest([]string{ticker})
		cobra.CheckErr(intradayErr)
		functions.ShowCmdPrinter(requests, ticker)
	},
	Example: `budgie stock show
	--ticker "MSFT, AAPL"
`,
}

var searchStockCmd = &cobra.Command{
	Use:   "search",
	Short: "fetch information of a given stock",
	PreRun: func(cmd *cobra.Command, args []string) {
		tokens.SetToken()
	},
	Run: func(cmd *cobra.Command, args []string) {
		request, tickerErr := api.TickerRequest(name)
		cobra.CheckErr(tickerErr)
		
		functions.SearchStockPrinter(request, name)
	},
	Example: `budgie search 
	--name="Apple"
`,
}

func init() {
	rootCmd.AddCommand(stockCmd)
	stockCmd.AddCommand(addStockCmd, removeStockCmd, showStockCmd, searchStockCmd)

	addStockCmd.Flags().StringVarP(&portfolio, "portfolio", "p", "", "Portfolio name (required)")
	addStockCmd.MarkFlagRequired("portfolio")
	addStockCmd.Flags().StringVarP(&ticker, "ticker", "t", "", "Company symbol (required)")
	addStockCmd.MarkFlagRequired("ticker")

	removeStockCmd.Flags().StringVarP(&portfolio, "portfolio", "p", "", "Portfolio name (required)")
	removeStockCmd.MarkFlagRequired("portfolio")
	removeStockCmd.Flags().StringVarP(&ticker, "ticker", "t", "", "Company symbol (required)")
	removeStockCmd.MarkFlagRequired("ticker")

	showStockCmd.Flags().StringVarP(&ticker, "ticker", "t", "", "Company symbol (required)")
	showStockCmd.MarkFlagRequired("ticker")

	searchStockCmd.Flags().StringVarP(&name, "name", "n", "", "Company name (required)")
	searchStockCmd.MarkFlagRequired("name")
}
