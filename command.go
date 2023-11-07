package serendip

import (
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func GenerateDiscordMessage() (string, error) {

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

func OnSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	n := i.ApplicationCommandData().Name
	if n == "wiki" {
		// generate title, summary and URL for a random Wikipedia page
		content, err := GenerateDiscordMessage()
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

func makeContent(result PageResult, pageId int, pageTitle string, pageURL string) string {
	extract := result.Query.Pages[strconv.Itoa(pageId)].Extract
	if len(extract) == 0 {
		return pageTitle + "\n" + "<" + pageURL + ">"
	} else {
		// TODO: use something like a template
		return "**" + pageTitle + "**" + "\n\n" + extract + "\n\n" + "<" + pageURL + ">"
	}
}
