package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func printViperInfo(filepath string) {
	viper.SetConfigName(filepath)
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(viper.GetViper())
	fmt.Println(viper.Get("sources"))
	fmt.Println(viper.Get("language"))
}
