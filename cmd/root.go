package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"github.com/ykdundar/budgie/internal"
	"io/ioutil"
	"os"
	"path/filepath"

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
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Version = "[1.0.0]"
}

// Find home directory
var home, _ = os.UserHomeDir()

// Database connection
// TODO: Save into home folder instead of repository
var dataBase, _ = sql.Open("sqlite3", "./budgie.db")

func initConfig() {
	// set config path
	var cfgFileName string = ".budgie"
	var cfgFileType string = "yaml"

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

	// https://www.sqlite.org/datatype3.html
	// https://www.sqlite.org/foreignkeys.html
	enableForeignKeys, _ := dataBase.Prepare(
		"PRAGMA foreign_keys = ON;",
	)
	_, fkErr := enableForeignKeys.Exec()
	cobra.CheckErr(fkErr)

	createPortfolio, _ := dataBase.Prepare(
		"CREATE TABLE IF NOT EXISTS portfolio (" +
			"id INTEGER PRIMARY KEY," +
			"name TEXT," +
			"currency TEXT" +
			"active INTEGER)",
	)

	_, portfolioErr := createPortfolio.Exec()
	cobra.CheckErr(portfolioErr)

	// TODO => date_buy & date_sell degıl buy_date
	createStock, _ := dataBase.Prepare(
		"CREATE TABLE IF NOT EXISTS stocks (" +
			"stockId INTEGER PRIMARY KEY," +
			"stockName TEXT," +
			"ticker TEXT," +
			"date_buy INTEGER," +
			"date_sell INTEGER," +
			"price_buy INTEGER ," +
			"price_sell INTEGER ," +
			"share INTEGER," +
			"portfolio_id INTEGER," +
			"FOREIGN KEY(portfolio_id)" +
			"REFERENCES portfolio(id))",
	)
	_, stockErr := createStock.Exec()
	cobra.CheckErr(stockErr) // TODO: Use helper of cobra to print errors to console

	/*
		createStock, _ = dataBase.Prepare("INSERT INTO stocks (stockName, ticker, date_buy, date_sell, price_buy, price_sell, share, portfolio_id ) VALUES (?,?,?,?,?,?,?,?)")
		createStock.Exec("Apple", "AAPL", 1645549453, 1645549650, 12.4, 15.3, 20, 1)

		stockRows, _ := dataBase.Query("SELECT stockName, ticker, dateBuy, dateSell, priceBuy, priceSell, share, portfolioId FROM stocks")
		var stockId int
		var stockName string
		var ticker string
		var dateBuy int
		var dateSell int
		var priceBuy int
		var priceSell int
		var share int
		var portfolioId int
		for stockRows.Next() {
			stockRows.Scan(&stockId, &stockName, &ticker, &dateBuy, &dateSell, &priceBuy, &priceSell, &share, &portfolioId)
		}
		fmt.Println(stockId, stockName, ticker, dateBuy, dateSell, priceBuy, priceSell, share, portfolioId)

	*/
}
