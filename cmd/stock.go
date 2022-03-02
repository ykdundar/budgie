package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
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
		fmt.Println("remove called")
	},
	Example: `budgie stock remove
	--portfolio "European Stocks"
	--ticker "MSFT"
`,
}

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "reports stock situations by the given time as integer ",
	Long: `reports stock situations by the given time as integer
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
	addCmd.Flags().StringVarP(&ticker, "ticker", "s", "", "Company name (required)")
	addCmd.MarkPersistentFlagRequired("ticker")

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
