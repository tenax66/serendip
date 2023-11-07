package serendip

import (
	"github.com/bwmarrin/discordgo"
)

func OnSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	n := i.ApplicationCommandData().Name
	if n == "wiki" {
		// generate title, summary and URL for a random Wikipedia page
		content, err := GenerateRandomArticleMessage()
		if err != nil {
			return
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
	} else if n == "search" {
		options := i.ApplicationCommandData().Options

		// convert the slice into a map
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		if query, ok := optionMap["query"]; ok {
			content, err := GenerateSearchResultMessage(query.StringValue())
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

}
