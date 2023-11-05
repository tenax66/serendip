package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/tenax66/serendip"
)

func main() {
	// Discord ボットのトークンを設定してください
	token := os.Getenv("SERENDIP_BOT_TOKEN")

	// Discord セッションを作成
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println("Error creating Discord session:", err)
		return
	}

	// Discord セッションを開始
	err = dg.Open()
	if err != nil {
		log.Println("Error opening Discord session:", err)
		return
	}

	log.Println("Bot is now running.")

	// テキストチャンネルのIDを指定してください
	channelID := os.Getenv("SERENDIP_TEXT_CHANNEL_ID")

	message, err := serendip.GenerateDiscordMessage()
	if err != nil {
		log.Println("Error generating message")
		return
	}

	// テキストチャンネルにメッセージを投稿します
	if _, err := dg.ChannelMessageSend(channelID, message); err != nil {
		log.Println("Error sending message:", err)
		return
	}

	// Discord セッションをクローズします
	dg.Close()
}
