package cmd

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/database/stocks"
	"github.com/ykdundar/budgie/database/transactions"
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
	cobra.OnInitialize(
		database.EnableForeignKeys,
		database.CreateTokensTable,
		database.CreatePortfoliosTable,
		stocks.CreateStocksTable,
		transactions.CreateTransactionsTable,
		initConfig,
	)

	rootCmd.Version = "[1.0.0]"
}

func initConfig() {
	// check if a token flag exists or not
	if token == "" {
		tokenRecord := database.FindToken()

		if tokenRecord == "" {
			cobra.CheckErr("Please enter an API token with config command!")
		}
		viper.Set("token", tokenRecord)
	}
}
