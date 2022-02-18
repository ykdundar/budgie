package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/ykdundar/budgie/cmd"
)

func main() {
	// position of this command is wrong!
	cmd.Execute()

	home_dir, _ := os.UserHomeDir()

	// // name of config file (without extension)
	viper.SetConfigName(".budgie")
	// // format of the config file
	viper.SetConfigType("yaml")
	// // path to look for the config file in
	viper.AddConfigPath(home_dir)

	// // Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		errMsg := errors.New("Please add your API key as explained in the README")
		fmt.Println(errMsg)
		os.Exit(1)
	}
}
