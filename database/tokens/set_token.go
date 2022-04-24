package tokens

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ykdundar/budgie/internal/functions"
)

func SetToken() {
	token := FindToken()
	tokenErr := functions.ValidateToken(token)

	if tokenErr != nil {
		cobra.CheckErr(tokenErr)
	}

	viper.Set("token", token)
}
