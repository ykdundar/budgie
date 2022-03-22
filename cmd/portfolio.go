package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database/portfolios"
)

// portfolioCmd represents the portfolio command
var portfolioCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "portfolio commend creates, updates, deletes and lists portfolios by using sub commends",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		portfolios.CreatePortfoliosTable()
	},
}

var createCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		portfolios.AddPortfolio(name, currency, active)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		portfolios.UpdatePortfolio(name, rename, currency, active)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		portfolios.DeletePortfolio(name)
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
	portfolioCmd.AddCommand(createCmd, updateCmd, deleteCmd, listCmd, showCmd)

	createCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	createCmd.Flags().StringVarP(&currency, "currency", "c", "USD", "Portfolio currency")
	createCmd.Flags().BoolVarP(&active, "active", "a", true, "Set to true if default portfolio")
	createCmd.MarkFlagRequired("name")

	updateCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	updateCmd.Flags().StringVarP(&rename, "rename", "r", "", "Update portfolio name (required)")
	updateCmd.Flags().StringVarP(&currency, "currency", "c", "USD", "Portfolio currency")
	updateCmd.Flags().BoolVarP(&active, "active", "a", true, "Set to true if default portfolio")
	updateCmd.MarkFlagRequired("name")

	deleteCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	deleteCmd.MarkFlagRequired("name")

	showCmd.Flags().StringVarP(&name, "name", "n", "", "Portfolio name")
}
