package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// token represents the API token obtained from marketstack
var token string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config is used to add or update an API token",
	Long: `First of all, you need to obtain an API token
from marketstack, and then you can use the 'token' flag to save it.
If you run the command again with a different token, your existing
token will be updated.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "API token (required)")
	configCmd.MarkPersistentFlagRequired("token")
}
