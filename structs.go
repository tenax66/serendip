package serendip

// A struct for parsing results of Wikipedia API
type WikipediaRandomResult struct {
	Query struct {
		Random []WikipediaPage `json:"random"`
	} `json:"query"`
}

type WikipediaSearchResult struct {
	Query struct {
		Search []WikipediaSnippet `json:"search"`
	} `json:"query"`
}

type WikipediaSnippet struct {
	Snippet string `json:"snippet"`
}

// A struct for Wikipedia page infomation
type WikipediaPage struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
