package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "budgie",
	Short: "budgie is a cli tool which follow your stock market transactions",
	Long: `You can create portfolios and add, update, remove your stocks into these portfolios and
 check your stocks or stock you are interested in`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Version = "[1.0.0]"
}


