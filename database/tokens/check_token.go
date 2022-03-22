package tokens

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func CheckToken(){
	token := FindToken()

	if token == "" {
		cobra.CheckErr("Please enter an API token with config command!")
	}

	viper.Set("token", token)
}
