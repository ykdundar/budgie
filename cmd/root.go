package cmd

import (
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "budgie",
	Short: "budgie allows you to manage your stock portfolios without leaving the command line",
	Args:  cobra.NoArgs,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(database.EnableForeignKeys)

	rootCmd.Version = "[1.0.0]"
}
