package cmd

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "budgie",
	Short: "budgie allows you to manage your stock purchases without leaving the command line",
	Long: `You can create portfolios and add, update, remove your stocks into these portfolios and
 check your stocks or stock you are interested in`,
	Args: cobra.NoArgs,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(database.EnableForeignKeys)

	rootCmd.Version = "[1.0.0]"
}
