package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// flags
var (
	portfolio string
	symbol    string
	today     string
	day       int
	week      int
	month     int
	year      int
)

// Stock Type
type Stock struct {
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
	},
	Example: `budgie stock add
	--portfolio "European Stocks"
	--ticker "MSFT" --date "06.02.2020"
	--price "180"
	--shares "20"
	--currency "USD"
`,
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
	stockCmd.AddCommand(addCmd, removeCmd, reportCmd)

	addCmd.Flags().StringVarP(&portfolio, "portfolio", "p", "", "Portfolio name (required)")
	addCmd.MarkPersistentFlagRequired("portfolio")
	addCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "Company name (required)")
	addCmd.MarkPersistentFlagRequired("symbol")
	addCmd.Flags().StringVarP(&currency, "currency", "c", "", "Stock currency (required)")
	addCmd.MarkPersistentFlagRequired("currency")

	removeCmd.Flags().StringVarP(&portfolio, "portfolio", "p", "", "Portfolio name (required)")
	removeCmd.MarkPersistentFlagRequired("portfolio")
	removeCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "Company name (required)")
	removeCmd.MarkPersistentFlagRequired("symbol")

	reportCmd.Flags().StringVarP(&today, "today", "t", "", "Portfolio name (required)")
	reportCmd.Flags().IntVarP(&day, "day", "d", 1, "Report last given number of days ")
	reportCmd.Flags().IntVarP(&week, "week", "w", 1, "Report last given number of weeks")
	reportCmd.Flags().IntVarP(&month, "month", "m", 1, "Report last given number of months")
	reportCmd.Flags().IntVarP(&year, "year", "y", 1, "Report last given number of years")
}
