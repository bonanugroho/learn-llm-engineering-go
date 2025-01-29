package web_scraper

import (
	"fmt"

	"github.com/gocolly/colly"

	model "github.com/bonanugroho/learn-llm-engineering-go/model/summarizer"
)

type WebSummarizerImpl struct {
	colly      *colly.Collector
	summarizer *model.Summarizer
}

func NewWebSummarizerImpl(colly *colly.Collector, summarizer *model.Summarizer) *WebSummarizerImpl {
	return &WebSummarizerImpl{colly: colly, summarizer: summarizer}
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

	if err := impl.colly.Visit(location); err != nil {
		return nil, err
	}

	impl.colly.Wait()

	return impl.summarizer, nil
}
