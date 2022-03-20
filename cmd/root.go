package cmd

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/internal"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "budgie",
	Short: "budgie allows you to manage your stock purchases without leaving the command line",
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
	cobra.OnInitialize(
		database.EnableForeignKeys,
		database.CreatePortfolioTable,
		database.CreateStocksTable,
		database.CreateTransactionsTable,
		initConfig,
	)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Version = "[1.0.0]"
}

// Find home directory
var home, _ = os.UserHomeDir()

func initConfig() {
	// set config path
	var cfgFileName = ".budgie"
	var cfgFileType = "yaml"

	// tell viper where to find config files
	viper.AddConfigPath(home)
	viper.SetConfigType(cfgFileType)
	viper.SetConfigName(cfgFileName)

	// Find and read the config file
	cfgErr := viper.ReadInConfig()

	// Create config file if it doesn't exist
	if cfgErr != nil {
		configFilePath := filepath.Join(home, cfgFileName+"."+cfgFileType)

		var yamlStr = []byte("token: \"REPLACE_WITH_YOUR_TOKEN\"")
		writeErr := ioutil.WriteFile(configFilePath, yamlStr, 0666)

		if writeErr != nil {
			fmt.Printf("Can not set config automatically. Please create it manually. Unable to write file: %v", writeErr)
			os.Exit(1)
		}
	}

	if token == "" {
		err := internal.CheckToken(viper.GetString("token"))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
