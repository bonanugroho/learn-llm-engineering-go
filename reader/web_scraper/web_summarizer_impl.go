package web_scraper

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"

	model "github.com/bonanugroho/learn-llm-engineering-go/model/summarizer"
)

type WebSummarizerImpl struct {
	colly      *colly.Collector
	summarizer *model.Summarizer
}

func NewWebSummarizerImpl() *WebSummarizerImpl {
	c := colly.NewCollector()
	s := model.Summarizer{
		Title: "",
		Text:  "",
		Links: make([]model.Link, 0),
	}
	return &WebSummarizerImpl{
		colly:      c,
		summarizer: &s,
	}
}

func (impl *WebSummarizerImpl) timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Scraping took: %s\n", time.Since(start))
	}
}

func (impl *WebSummarizerImpl) Read(location string) (any, error) {

	impl.colly.OnHTML("head title", func(e *colly.HTMLElement) {
		impl.summarizer.Title = e.Text
	})

	impl.colly.OnHTML("h1", func(e *colly.HTMLElement) {
		impl.summarizer.Text = impl.summarizer.Text + fmt.Sprintf("%s \n", e.Text)
	})

	impl.colly.OnHTML("p", func(e *colly.HTMLElement) {
		impl.summarizer.Text = impl.summarizer.Text + fmt.Sprintf("%s \n", e.Text)
	})

	impl.colly.OnHTML("a", func(e *colly.HTMLElement) {
		impl.summarizer.Text = impl.summarizer.Text + fmt.Sprintf("%s \n", e.Text)

		link := model.Link{
			Title: e.Text,
			Url:   e.Attr("href"),
		}

		impl.summarizer.Links = append(impl.summarizer.Links, link)
	})

	defer impl.timer()
	
	if err := impl.colly.Visit(location); err != nil {
		return nil, err
	}

	impl.colly.Wait()

	return impl.summarizer, nil
}
