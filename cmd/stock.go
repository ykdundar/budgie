package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// flags
var (
	portfolio string
	ticker    string
	today     string
	day       int
	week      int
	month     int
	year      int
)

// Stock Type
type Stock struct {
	stockId     int
	name        string
	ticker      string
	buyDate     int
	sellDate    int
	buyPrice    int
	sellPrice   int
	portfolioId int
	shares      int
}

// stockCmd represents the stock command
var stockCmd = &cobra.Command{
	Use:   "stock",
	Short: "stock command adds, removes, buys, sells and reports a given stock by subcommands",
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a stock to a given portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		addStock, _ := dataBase.Prepare(
			"INSERT INTO stocks (portfolio_id, ticker) VALUES (?,?)")
		defer addStock.Close()
		_, insertErr := addStock.Exec(portfolio, ticker)
		cobra.CheckErr(insertErr)
	},

	Example: `budgie stock add
	--portfolio "European Stocks"
	--ticker "MSFT"
	--price "180"
	--shares "20"
	--currency "USD"
`,
}
var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "Add the stock you bought to portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("buy called")
	},
	Example: "TODO",
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a given stock",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
	Example: "TODO",
}

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "reports stock situations by the given time as intiger ",
	Long: `reports stock situations by the given time as intiger
For example:
day 5
week 3
mont 9
year 2`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("report called")
	},
}

func init() {
	rootCmd.AddCommand(stockCmd)
	stockCmd.AddCommand(addCmd, buyCmd, removeCmd, reportCmd)

	addCmd.Flags().StringVarP(&portfolio, "portfolio", "p", "", "Portfolio name (required)")
	addCmd.MarkPersistentFlagRequired("portfolio")
	addCmd.Flags().StringVarP(&ticker, "ticker", "s", "", "Company name (required)")
	addCmd.MarkPersistentFlagRequired("ticker")

	buyCmd.Flags().StringVarP(&portfolio, "portfolio", "p", "", "Portfolio name (required)")

	removeCmd.Flags().StringVarP(&portfolio, "portfolio", "p", "", "Portfolio name (required)")
	removeCmd.MarkPersistentFlagRequired("portfolio")
	removeCmd.Flags().StringVarP(&ticker, "ticker", "s", "", "Company name (required)")
	removeCmd.MarkPersistentFlagRequired("ticker")

	reportCmd.Flags().StringVarP(&today, "today", "t", "", "Portfolio name (required)")
	reportCmd.Flags().IntVarP(&day, "day", "d", 1, "Report last given number of days ")
	reportCmd.Flags().IntVarP(&week, "week", "w", 1, "Report last given number of weeks")
	reportCmd.Flags().IntVarP(&month, "month", "m", 1, "Report last given number of months")
	reportCmd.Flags().IntVarP(&year, "year", "y", 1, "Report last given number of years")
}
