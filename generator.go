package serendip

import (
	"log"
	"net/url"
	"strconv"
)

func GenerateRandomArticleMessage() (string, error) {

	pageId, pageTitle, pageURL, err := GetRandomPage()

	if err != nil {
		log.Fatal("Error getting Wikipedia page: ", err)
		return "", err
	}

	result, err := GetPageContent(pageId)
	if err != nil {
		log.Fatal("Error getting Wikipedia details: ", err)
		return "", err
	}

	content := makeContent(result, pageId, pageTitle, pageURL)
	return content, nil
}

func GenerateSearchResultMessage(query string) (string, error) {
	searchRes, err := SearchArticle(query)
	if err != nil {
		log.Fatal("Error searching Wikipedia articles: ", err)
		return "", err
	}

	var content string
	for _, ps := range searchRes.Query.PrefixSearch {
		// pass an empty string because summary is not needed
		content += buildPost(ps.Title, "", "https://ja.wikipedia.org/wiki/"+url.PathEscape(ps.Title))
		content += "\n"
		log.Println("content: ", content)
	}

	return content, nil
}

func makeContent(result PageResult, pageId int, pageTitle string, pageURL string) string {
	extract := result.Query.Pages[strconv.Itoa(pageId)].Extract
	return buildPost(pageTitle, extract, pageURL)
}

// Build a discord post. This formats text using markdown.
func buildPost(title string, text string, url string) string {
	// TODO: simplify this branching
	if text == "" {
		return title + "\n" + "<" + url + ">"
	} else {
		return "**" + title + "**" + "\n\n" + text + "\n\n" + "<" + url + ">"
	}
}
