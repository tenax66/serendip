README.md

# Serendip Bot

Serendip is a Discord bot that gets random Wikipedia pages. This bot is built using Go and the [discordgo](https://github.com/bwmarrin/discordgo) library.

## Getting Started

The following instructions assume you already have a working Go environment, if not please see [this page](https://go.dev/doc/install) first.

### Creating a new Discord bot

To set up the bot, you will need to have a Discord account and create a new Discord bot. You can follow the instructions in the [Discord developer portal](https://discord.com/developers/docs/getting-started "Getting Started") to create a new bot and get the bot token.

### Running

Before running, set the environment variable `SERENDIP_BOT_TOKEN` to the bot token obtained in the step above.

```
git clone https://github.com/tenax66/serendip.git
go build -o main
./main
```

## Usage

| Command | Command Description                                                                                                   |
| ------- | --------------------------------------------------------------------------------------------------------------------- |
| /wiki   | Fetches a random Wikipedia page and responds with a title of the page, a summary of the page, and a link to the page. |
