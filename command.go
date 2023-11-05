package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/net/html"
)

const endpoint = "https://ja.wikipedia.org/w/api.php"

func paramsTemplate() url.Values {
	params := url.Values{}
	params.Set("action", "query")
	params.Set("format", "json")
	return params
}

func getWikipediaAPI(params url.Values) (*http.Response, error) {
	// APIにリクエストを送信

	resp, err := http.Get(endpoint + "?" + params.Encode())
	if err != nil {
		return resp, err
	}

	return resp, err
}

func makeContent(result WikipediaSearchResult, pageTitle string, pageURL string) (string, error) {
	if len(result.Query.Search) == 0 {
		return pageTitle + "\n" + "<" + pageURL + ">", nil
	} else {
		snippet := &result.Query.Search[0].Snippet
		rawSnippet, err := html.Parse(strings.NewReader(*snippet))

		if err != nil {
			return "", err
		}
		pageDetail := removeTagFromText(rawSnippet)
		return pageTitle + "\n\n" + pageDetail + "\n\n" + "<" + pageURL + ">", err
	}
}

func generateDiscordMessage() (string, error) {

	pageTitle, pageURL, err := getRandomPage()

	if err != nil {
		log.Fatal("Error getting Wikipedia page: ", err)
		return "", err
	}

	result, err := getSearchPage(pageTitle)
	if err != nil {
		log.Fatal("Error getting Wikipedia details: ", err)
		return "", err
	}

	content, err := makeContent(result, pageTitle, pageURL)
	return content, err
}

func onSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "wiki" {
		// 取得したページのURLを返信
		content, err := generateDiscordMessage()
		if err != nil {
			return
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
	}
}

// Wikipediaのランダムなページを取得する関数
func getRandomPage() (string, string, error) {
	// リクエストパラメータを作成
	params := paramsTemplate()
	params.Set("list", "random")
	params.Set("rnnamespace", "0")
	params.Set("rnlimit", "1")

	resp, err := getWikipediaAPI(params)

	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()

	// レスポンスをパース
	var result WikipediaRandomResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", err
	}

	// 結果を返す
	if len(result.Query.Random) == 0 {
		return "", "", fmt.Errorf("no random pages found")
	}

	page := &result.Query.Random[0]
	return page.Title, "https://ja.wikipedia.org/wiki/" + url.PathEscape(page.Title), nil
}

func getSearchPage(pageTitle string) (WikipediaSearchResult, error) {
	params := paramsTemplate()
	params.Set("list", "search")
	params.Set("srsearch", pageTitle)
	params.Set("srlimit", "1")

	var result WikipediaSearchResult

	resp, err := getWikipediaAPI(params)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

func removeTagFromText(n *html.Node) string {

	if n.Type == html.TextNode {
		return n.Data
	}

	var str_build strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		str_build.WriteString(removeTagFromText(c))
	}

	return strings.Join(strings.Fields(str_build.String()), " ")
}

// Wikipedia APIのレスポンスをパースするための型定義
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

// Wikipediaのページ情報を表す型定義
type WikipediaPage struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
