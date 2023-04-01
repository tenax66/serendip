package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// DiscordのBot Tokenをセット
	dg, err := discordgo.New("Bot " + os.Getenv("SERENDIP_BOT_TOKEN"))
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// !wikipediaコマンドが送信されたときの処理を設定
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Content == "!wikipedia" {
			// ランダムなWikipedia記事のURLを取得
			url := "https://ja.wikipedia.org/wiki/Special:Random"

			// チャンネルにURLを投稿
			s.ChannelMessageSend(m.ChannelID, url)
		}
	})

	// Discordに接続
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
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
