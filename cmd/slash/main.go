package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/tenax66/serendip"
)

var testOpt = flag.Bool("t", false, "テスト用です。投稿する文章をDiscordを経由せず表示します")

func main() {
	flag.Parse()
	if *testOpt {
		content, _ := serendip.GenerateDiscordMessage()
		fmt.Println(content)
		return
	}

	TOKEN := os.Getenv("SERENDIP_BOT_TOKEN")

	// DiscordのBot Tokenをセット
	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		log.Println("Error creating Discord session: ", err)
		return
	}

	// /wikiコマンドが送信されたときの処理を設定
	dg.AddHandler(serendip.OnSlashCommand)

	// Discordに接続
	err = dg.Open()
	if err != nil {
		log.Println("Error opening Discord session: ", err)
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
		log.Println("Error registering slash commands: ", err)
		return
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")

	// 終了シグナルを待機
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Discordから切断
	dg.Close()
}
