package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database/portfolios"
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
		addPortfolio := portfolios.AddPortfolio(name, currency, active)

		if addPortfolio == nil{
			fmt.Printf("'%s' is created succesfully\n", name)
		}
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		updatePortfolio := portfolios.UpdatePortfolio(name, rename, currency, active)

		if updatePortfolio == nil{
			fmt.Printf("'%s' is updated succesfully\n", name)
		}
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		deletePortfolio := portfolios.DeletePortfolio(name)
		if deletePortfolio == nil{
			fmt.Printf("'%s' is deleted succesfully", name)
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all portfolios",
	Run: func(cmd *cobra.Command, args []string) {
		portfolios.SelectAllPortfolios()
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows an active portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(portfolios.FindPortfolio(name))
	},
}

func init() {
	rootCmd.AddCommand(portfolioCmd)
	portfolioCmd.AddCommand(addPortfolioCmd, updateCmd, deleteCmd, listCmd, showCmd)

	addPortfolioCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	addPortfolioCmd.Flags().StringVarP(&currency, "currency", "c", "USD", "Portfolio currency")
	addPortfolioCmd.Flags().BoolVarP(&active, "active", "a", true, "Set to true if default portfolio")
	addPortfolioCmd.MarkFlagRequired("name")

	updateCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	updateCmd.Flags().StringVarP(&rename, "rename", "r", "", "Update portfolio name")
	updateCmd.Flags().StringVarP(&currency, "currency", "c", "USD", "Portfolio currency")
	updateCmd.Flags().BoolVarP(&active, "active", "a", true, "Set to true if default portfolio")
	updateCmd.MarkFlagRequired("name")

	deleteCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	deleteCmd.MarkFlagRequired("name")

	// bu zorunlu mu? yoksa aktifi mi gosterecek? aktif ne ki?
	showCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name")
}
