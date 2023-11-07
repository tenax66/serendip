package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/tenax66/serendip"
)

func main() {
	token := os.Getenv("SERENDIP_BOT_TOKEN")

	// create a discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println("Error creating Discord session:", err)
		return
	}

	// start a discord session
	err = dg.Open()
	if err != nil {
		log.Println("Error opening Discord session:", err)
		return
	}

	log.Println("Bot is now running.")

	channelID := os.Getenv("SERENDIP_TEXT_CHANNEL_ID")

	message, err := serendip.GenerateRandomArticleMessage()
	if err != nil {
		log.Println("Error generating message")
		return
	}

	// post the message
	if _, err := dg.ChannelMessageSend(channelID, message); err != nil {
		log.Println("Error sending message:", err)
		return
	}

	// disconnect
	dg.Close()
}
