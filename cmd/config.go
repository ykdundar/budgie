package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/internal"
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
		err := internal.CheckToken(token)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		homeDir, _ := os.UserHomeDir()
		fileName := ".budgie.yaml"
		configFilePath := filepath.Join(homeDir, fileName)

		var yamlStr = []byte(fmt.Sprintf("token: \"%s\"", token))

		writeErr := ioutil.WriteFile(configFilePath, yamlStr, 0666)

		if writeErr != nil {
			fmt.Printf("Unable to write file: %v", err)
		} else {
			fmt.Println("Token successfully updated!")
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "API token (required)")
	configCmd.MarkPersistentFlagRequired("token")
}
