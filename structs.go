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

// search results
type SearchResult struct {
    PrefixSearch []PrefixSearch `json:"prefixsearch"`
}

type PrefixSearch struct {
    NS     int    `json:"ns"`
    Title  string `json:"title"`
    PageID int    `json:"pageid"`
}

type SearchContinue struct {
    PSOffset int    `json:"psoffset"`
    Continue string `json:"continue"`
}

type SearchResponse struct {
    BatchComplete string   `json:"batchcomplete"`
    Continue      SearchContinue `json:"continue"`
    Query         SearchResult    `json:"query"`
}
