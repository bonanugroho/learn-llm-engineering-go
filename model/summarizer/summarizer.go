package summarizer

type Link struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type Summarizer struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Links []Link `json:"links"`
}
