package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/tenax66/serendip"
)

var testOpt = flag.Bool("t", false, "test flag")

func main() {
	flag.Parse()
	if *testOpt {
		content, _ := serendip.GenerateRandomArticleMessage()
		log.Println(content)
		return
	}

	TOKEN := os.Getenv("SERENDIP_BOT_TOKEN")

	// set the bot token
	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		log.Println("Error creating Discord session: ", err)
		return
	}

	// handling /wiki command
	dg.AddHandler(serendip.OnSlashCommand)

	// connect to the server
	err = dg.Open()
	if err != nil {
		log.Println("Error opening Discord session: ", err)
		return
	}

	// register slash commands
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "wiki",
			Type:        1,
			Description: "Get a link to a random Wikipedia article.",
		},
		{
			Name:        "search",
			Type:        1,
			Description: "Search Wikipedia articles by given search words.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "query",
					Type:        3,
					Required:    true,
					Description: "search words",
				},
			},
		},
	}

	_, err = dg.ApplicationCommandBulkOverwrite(dg.State.User.ID, "", commands)
	if err != nil {
		log.Println("Error registering slash commands: ", err)
		return
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")

	// wait for signals
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// disconnect
	dg.Close()
}
