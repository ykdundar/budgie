package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// flags
var (
portfolio string
symbol string
today string
day int
week int
month int
year int

)

// stockCmd represents the stock command
var stockCmd = &cobra.Command{
	Use:   "stock",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("report called")
	},
}



/*
var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "Aliases of add command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

var sellCmd = &cobra.Command{
	Use:   "sell",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}*/


func init() {
	rootCmd.AddCommand(stockCmd)
	stockCmd.AddCommand(addCmd)
	stockCmd.AddCommand(removeCmd)
	stockCmd.AddCommand(reportCmd)


	addCmd.Flags().StringVarP(&portfolio,"portfolio", "p", "", "Portfolio name (required)")
	addCmd.MarkPersistentFlagRequired("portfolio")
	addCmd.Flags().StringVarP(&symbol,"symbol", "s", "", "Company name (required)")
	addCmd.MarkPersistentFlagRequired("symbol")
	addCmd.Flags().StringVarP(&currency,"currency", "c", "", "Stock currency (required)")
	addCmd.MarkPersistentFlagRequired("currency")

	removeCmd.Flags().StringVarP(&portfolio,"portfolio", "p", "", "Portfolio name (required)")
	removeCmd.MarkPersistentFlagRequired("portfolio")
	removeCmd.Flags().StringVarP(&symbol,"symbol", "s", "", "Company name (required)")
	removeCmd.MarkPersistentFlagRequired("symbol")

	reportCmd.Flags().StringVarP(&today,"today", "t", "", "Portfolio name (required)")
	reportCmd.Flags().IntVarP(&day,"day", "d", 1, "Report last given number of days ")
	reportCmd.Flags().IntVarP(&week,"week", "w", 1, "Report last given number of weeks")
	reportCmd.Flags().IntVarP(&month,"month", "m", 1, "Report last given number of months")
	reportCmd.Flags().IntVarP(&year,"year", "y", 1, "Report last given number of years")
}