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

// stockOperationsCmd represents the stockOperations command
var stockOperationsCmd = &cobra.Command{
	Use:   "stockOperations",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var addStockCmd = &cobra.Command{
	Use:   "addStock",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addStock called")
	},
}

var removeStockCmd = &cobra.Command{
	Use:   "removeStock",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("removeStock called")
	},
}

func init() {
	rootCmd.AddCommand(stockOperationsCmd)
	stockOperationsCmd.AddCommand(addStockCmd)
	stockOperationsCmd.AddCommand(removeStockCmd)

	addStockCmd.Flags().StringVarP(&portfolio,"portfolio", "p", "", "Portfolio name (required)")
	addStockCmd.MarkPersistentFlagRequired("portfolio")
	addStockCmd.Flags().StringVarP(&ticker,"ticker", "t", "", "Company name (required)")
	addStockCmd.MarkPersistentFlagRequired("ticker")
	addStockCmd.Flags().StringVarP(&currency,"currency", "c", "", "Stock currency")

	removeStockCmd.Flags().StringVarP(&portfolio,"portfolio", "p", "", "Portfolio name (required)")
	removeStockCmd.MarkPersistentFlagRequired("portfolio")
	removeStockCmd.Flags().StringVarP(&ticker,"ticker", "t", "", "Company name (required)")
	removeStockCmd.MarkPersistentFlagRequired("ticker")








	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stockOperationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stockOperationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
