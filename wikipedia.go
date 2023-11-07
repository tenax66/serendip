package serendip

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const ENDPOINT = "https://ja.wikipedia.org/w/api.php"

// get a random page of Wikipedia
func GetRandomPage() (int, string, string, error) {
	// create request parameters
	params := createParamsTemplate()
	params.Set("list", "random")
	params.Set("rnnamespace", "0")
	params.Set("rnlimit", "1")

	resp, err := requestWikipediaAPI(params)

	if err != nil {
		return 0, "", "", err
	}

	defer resp.Body.Close()

	var result WikipediaRandomResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, "", "", err
	}

	if len(result.Query.Random) == 0 {
		return 0, "", "", fmt.Errorf("no random pages found")
	}

	page := &result.Query.Random[0]
	return page.Id, page.Title, "https://ja.wikipedia.org/wiki/" + url.PathEscape(page.Title), nil
}

func GetPageContent(pageId int) (PageResult, error) {
	params := createParamsTemplate()
	params.Set("prop", "extracts")
	params.Set("explaintext", "")
	params.Set("exintro", "")
	params.Set("redirects", "1")
	params.Set("pageids", strconv.Itoa(pageId))
	params.Set("utf8", "")

	var result PageResult

	resp, err := requestWikipediaAPI(params)

	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

func createParamsTemplate() url.Values {
	params := url.Values{}
	params.Set("action", "query")
	params.Set("format", "json")
	return params
}

func requestWikipediaAPI(params url.Values) (*http.Response, error) {
	url := ENDPOINT + "?" + params.Encode()
	log.Println("url:", url)
	resp, err := http.Get(ENDPOINT + "?" + params.Encode())
	if err != nil {
		return resp, err
	}

	return resp, err
}
