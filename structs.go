package serendip

// A struct for parsing results of Wikipedia API
type WikipediaRandomResult struct {
	Query struct {
		Random []WikipediaPage `json:"random"`
	} `json:"query"`
}

type PageResult struct {
	Query struct {
		Pages map[string]Page `json:"pages"`
	} `json:"query"`
}

type Page struct {
	PageID  int    `json:"pageid"`
	NS      int    `json:"ns"`
	Title   string `json:"title"`
	Extract string `json:"extract"`
}

// A struct for Wikipedia page infomation
type WikipediaPage struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
