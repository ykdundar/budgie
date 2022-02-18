package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName(".budgie") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME") // path to look for the config file in	

	viper.SetEnvPrefix("budgie") // will be uppercased automatically
	viper.BindEnv("apikey")
	os.Setenv("BUDGIE_APIKEY", "13")

	fmt.Println(os.Getenv("BUDGIE_APIKEY"))
	// os.Setenv("SPF_ID", "13") // typically done outside of the app

	// id := Get("id") // 13

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error: %w \n", err))
	}

	// viper.Set("LogFile", LogFile)
	// cmd.Execute()
}
