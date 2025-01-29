package main

import (
	"fmt"

	"github.com/spf13/viper"

	settings "github.com/bonanugroho/learn-llm-engineering-go/model"
	ws "github.com/bonanugroho/learn-llm-engineering-go/reader/web_scraper"
)

var Settings = new(settings.Settings)

const Url = "https://edwarddonner.com"

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
	reader := ws.NewWebSummarizerImpl()
	res, err := reader.Read(Url)

	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}
