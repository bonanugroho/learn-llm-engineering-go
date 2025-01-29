package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/spf13/viper"

	settings "github.com/bonanugroho/learn-llm-engineering-go/model"
	summarizer "github.com/bonanugroho/learn-llm-engineering-go/model/summarizer"
	"github.com/bonanugroho/learn-llm-engineering-go/reader/web_scraper"
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
	fmt.Println("Summarizing!")

	fmt.Printf("%v\n", Settings)

	c := colly.NewCollector()
	tempSummarizer := summarizer.Summarizer{
		Title: "",
		Text:  "",
		Links: make([]summarizer.Link, 0),
	}

	reader := web_scraper.NewWebSummarizerImpl(c, &tempSummarizer)
	res, err := reader.Read(Url)

	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}
