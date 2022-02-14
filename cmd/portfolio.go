package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// flags
var (
	name string 
	currency string 
	active bool 
	rename string
)

// portfolioCmd represents the portfolio command
var portfolioCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "portfolio commend creates, updates, deletes and lists portfolios by using sub commends",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all portfolios",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows an active portfolio",
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("en son ben calistim!")
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("once ben calistim!")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show called")
	},
}

func init() {
	rootCmd.AddCommand(portfolioCmd)
	portfolioCmd.AddCommand(createCmd, updateCmd, deleteCmd, listCmd, showCmd)

	createCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	createCmd.PersistentFlags().StringVarP(&currency, "currency", "c", "USD", "Portfolio currency")
	createCmd.PersistentFlags().BoolVarP(&active, "active", "a", true, "Set to true if default portfolio")
	createCmd.MarkPersistentFlagRequired("name")

	updateCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	updateCmd.PersistentFlags().StringVarP(&rename, "rename", "r", "", "Update portfolio name (required)")
	updateCmd.PersistentFlags().StringVarP(&currency, "currency", "c", "USD", "Portfolio currency")
	updateCmd.PersistentFlags().BoolVarP(&active, "active", "a", true, "Set to true if default portfolio")
	updateCmd.MarkPersistentFlagRequired("name")

	deleteCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	deleteCmd.MarkPersistentFlagRequired("name")

	showCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Portfolio name")
}
