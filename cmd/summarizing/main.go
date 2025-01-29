package main

import (
	"fmt"
	"github.com/bonanugroho/learn-llm-engineering-go/model"

	"github.com/spf13/viper"
)

var Settings = new(model.Settings)

func init() {
	config := viper.New()
	config.AddConfigPath(".")
	config.SetConfigName("config")
	config.SetConfigType("json")

	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err.Error()))
	}

	if err := config.Unmarshal(&Settings); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err.Error()))
	}
}

func main() {
	fmt.Println("Summarizing!")

	fmt.Printf("%v\n", Settings)
}
