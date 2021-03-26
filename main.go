package main

import (
	"fmt"

	"main/app"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	s := viper.GetViper()

	a := app.App{
		Mux:      r,
		Settings: s,
	}

	a.Run()
	defer a.Close()
}
