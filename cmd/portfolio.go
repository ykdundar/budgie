package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/api"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/database/portfolios"
	"github.com/ykdundar/budgie/database/tokens"
	"github.com/ykdundar/budgie/internal/objects"
)

// portfolioCmd represents the portfolio command
var portfolioCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "portfolio is used to create, update, delete and list portfolios",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		portfolios.CreatePortfoliosTable()
	},
}

var addPortfolioCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		portfolios.AddPortfolio(name, currency)
		fmt.Printf("'%s' is created succesfully\n", name)
	},
}

var updatePortfolioCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		portfolios.UpdatePortfolio(name, rename, currency)
		fmt.Printf("'%s' is updated succesfully\n", name)
	},
}

var deletePortfolioCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		portfolios.DeletePortfolio(name)
		fmt.Printf("'%s' is deleted succesfully", name)
	},
}

var listAllPortfoliosCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all portfolios",
	Run: func(cmd *cobra.Command, args []string) {
		selectAllPortfolios := portfolios.SelectAllPortfolios()
		for _, v := range selectAllPortfolios {
			fmt.Println("Id: ", v.Id, "Name: ", v.Name, "Currency: ", v.Currency)
		}
	},
}

var showPortfolioCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows a portfolio",
	PreRun: func(cmd *cobra.Command, args []string) {
		tokens.SetToken()
	},
	Run: func(cmd *cobra.Command, args []string) {
		portfolioId := portfolios.FindPortfolio(name).Id

		records, queryErr := database.DBConnection().Query("SELECT ticker FROM stocks WHERE portfolio_id=?", portfolioId)
		defer records.Close()
		cobra.CheckErr(queryErr)

		stock := objects.Stock{}
		var tickerSlc []string

		for records.Next() {
			scanErr := records.Scan(&stock.Ticker)
			cobra.CheckErr(scanErr)
			tickerSlc = append(tickerSlc, stock.Ticker)
		}

		requests, intradayErr := api.IntradayRequest(tickerSlc)
		cobra.CheckErr(intradayErr)
		fmt.Println(requests)
	},
}

func init() {
	rootCmd.AddCommand(portfolioCmd)
	portfolioCmd.AddCommand(addPortfolioCmd, updatePortfolioCmd, deletePortfolioCmd, listAllPortfoliosCmd, showPortfolioCmd)

	addPortfolioCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	addPortfolioCmd.Flags().StringVarP(&currency, "currency", "c", "USD", "Portfolio currency")
	addPortfolioCmd.MarkFlagRequired("name")

	updatePortfolioCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	updatePortfolioCmd.Flags().StringVarP(&rename, "rename", "r", "", "Update portfolio name")
	updatePortfolioCmd.Flags().StringVarP(&currency, "currency", "c", "USD", "Portfolio currency")
	updatePortfolioCmd.MarkFlagRequired("name")

	deletePortfolioCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	deletePortfolioCmd.MarkFlagRequired("name")

	showPortfolioCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name (required")
	showPortfolioCmd.MarkFlagRequired("name")
}
