/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

func init() {
	rootCmd.AddCommand(portfolioCmd)
	portfolioCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	createCmd.PersistentFlags().StringVarP(&currency, "currency", "c", "USD", "Portfolio currency")
	createCmd.PersistentFlags().BoolVarP(&active, "active", "a", true, "Set to true if default portfolio")
	createCmd.MarkPersistentFlagRequired("name")

	updateCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	updateCmd.PersistentFlags().StringVarP(&rename, "rename", "r", "", "Update portfolio name (required)")
	updateCmd.PersistentFlags().StringVarP(&currency, "currency", "c", "USD", "Portfolio currency")
	updateCmd.PersistentFlags().BoolVarP(&active, "active", "a", true, "Set to true if default portfolio")
	updateCmd.MarkPersistentFlagRequired("name")

}
