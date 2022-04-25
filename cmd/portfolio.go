package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/api"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/database/portfolios"
	"github.com/ykdundar/budgie/database/tokens"
	"github.com/ykdundar/budgie/internal/objects"
	"github.com/ykdundar/budgie/internal/tableprinters"
)

// portfolioCmd represents the portfolio command.
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
		portfolios.AddPortfolio(name)
		fmt.Printf("'%s' is created database.DBConnection()\n", name)
	},
}

var updatePortfolioCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		portfolios.UpdatePortfolio(name, rename)
		fmt.Printf("'%s' is updated successfully\n", name)
	},
}

var deletePortfolioCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		portfolios.DeletePortfolio(name)
		fmt.Printf("'%s' is deleted successfully\n", name)
	},
}

var listAllPortfoliosCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all portfolios",
	Run: func(cmd *cobra.Command, args []string) {
		selectAllPortfolios := portfolios.SelectAllPortfolios()
		tableprinters.ListPortfolioPrinter(selectAllPortfolios, "My Portfolios")
	},
}

var showPortfolioCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows a portfolio",
	PreRun: func(cmd *cobra.Command, args []string) {
		tokens.SetToken()
	},
	Run: func(cmd *cobra.Command, args []string) {
		portfolioID := portfolios.FindPortfolio(name).ID

		records, queryErr := database.DBConnection().Query("SELECT ticker FROM stocks WHERE portfolio_id=?", portfolioID)
		cobra.CheckErr(queryErr)
		defer records.Close()

		stock := objects.Stock{}
		var tickerSlc []string

		for records.Next() {
			scanErr := records.Scan(&stock.Ticker)
			cobra.CheckErr(scanErr)
			tickerSlc = append(tickerSlc, stock.Ticker)
		}

		if tickerSlc == nil {
			cobra.CheckErr("No stocks have been found for this portfolio")
		}

		requests, intradayErr := api.IntradayRequest(tickerSlc)
		cobra.CheckErr(intradayErr)

		tableprinters.ShowCmdPrinter(requests, name)
	},
}

func init() {
	rootCmd.AddCommand(portfolioCmd)
	portfolioCmd.AddCommand(addPortfolioCmd, updatePortfolioCmd, deletePortfolioCmd, listAllPortfoliosCmd, showPortfolioCmd)

	addPortfolioCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	addPortfolioCmd.MarkFlagRequired("name")

	updatePortfolioCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	updatePortfolioCmd.Flags().StringVarP(&rename, "rename", "r", "", "Update portfolio name")
	updatePortfolioCmd.MarkFlagRequired("name")

	deletePortfolioCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	deletePortfolioCmd.MarkFlagRequired("name")

	showPortfolioCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	showPortfolioCmd.MarkFlagRequired("name")
}
