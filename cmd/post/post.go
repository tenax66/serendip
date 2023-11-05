package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Discord ボットのトークンを設定してください
	token := "YOUR_DISCORD_BOT_TOKEN"

	// Discord セッションを作成します
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// Discord セッションを開始します
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	// テキストチャンネルに定期的にメッセージを投稿するループ
	for {
		// テキストチャンネルのIDを指定してください
		channelID := "YOUR_TEXT_CHANNEL_ID"

		// テキストチャンネルにメッセージを投稿します
		_, err := dg.ChannelMessageSend(channelID, "Hello, World!")
		if err != nil {
			fmt.Println("Error sending message:", err)
		}

		// 10秒ごとにメッセージを投稿します
		time.Sleep(10 * time.Second)
	}

	// 終了シグナルを受信するまで待機します
	// CTRL+C を押すとプログラムが終了します
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Discord セッションをクローズします
	dg.Close()
}
