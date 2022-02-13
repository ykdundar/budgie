package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// flags
var (
portfolio string
ticker string
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

func init() {
	rootCmd.AddCommand(stockCmd)
	stockCmd.AddCommand(addCmd)
	stockCmd.AddCommand(removeCmd)

	addCmd.Flags().StringVarP(&portfolio,"portfolio", "p", "", "Portfolio name (required)")
	addCmd.MarkPersistentFlagRequired("portfolio")
	addCmd.Flags().StringVarP(&ticker,"ticker", "t", "", "Company name (required)")
	addCmd.MarkPersistentFlagRequired("ticker")
	addCmd.Flags().StringVarP(&currency,"currency", "c", "", "Stock currency")

	removeCmd.Flags().StringVarP(&portfolio,"portfolio", "p", "", "Portfolio name (required)")
	removeCmd.MarkPersistentFlagRequired("portfolio")
	removeCmd.Flags().StringVarP(&ticker,"ticker", "t", "", "Company name (required)")
	removeCmd.MarkPersistentFlagRequired("ticker")
}