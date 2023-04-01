package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const endpoint = "https://ja.wikipedia.org/w/api.php"

func main() {
	TOKEN := os.Getenv("SERENDIP_BOT_TOKEN")

	// DiscordのBot Tokenをセット
	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// !wikipediaコマンドが送信されたときの処理を設定
	dg.AddHandler(onSlashCommand)

	// Discordに接続
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	// スラッシュコマンドの登録
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "wiki",
			Type:        1,
			Description: "Get a link to a random Wikipedia article.",
		},
	}

	_, err = dg.ApplicationCommandBulkOverwrite(dg.State.User.ID, "", commands)
	if err != nil {
		fmt.Println("Error registering slash commands: ", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	// 終了シグナルを待機
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Discordから切断
	dg.Close()
}

func onSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "wiki" {
		// 取得したページのURLを返信
		pageTitle, pageURL, err := getRandomPage()
		if err != nil {
			log.Fatal("Error getting Wikipedia page: ", err)
			return
		}

		content := pageTitle + "\n" + pageURL

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
	params := url.Values{}
	params.Set("action", "query")
	params.Set("format", "json")
	params.Set("list", "random")
	params.Set("rnnamespace", "0")
	params.Set("rnlimit", "1")

	// APIにリクエストを送信
	resp, err := http.Get(endpoint + "?" + params.Encode())
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// レスポンスをパース
	var result WikipediaAPIResult
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

// Wikipedia APIのレスポンスをパースするための型定義
type WikipediaAPIResult struct {
	Query struct {
		Random []WikipediaPage `json:"random"`
	} `json:"query"`
}

// Wikipediaのページ情報を表す型定義
type WikipediaPage struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
