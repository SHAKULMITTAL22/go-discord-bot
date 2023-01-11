package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func CreateNewBotConnection() {
	discordSecret := ""
	discordBot := CreateNewBot(discordSecret)
	OpenConnection(discordBot)
}

func StartHandlers(bot *discordgo.Session) {
	bot.AddHandler(messageHandler)
}

func OpenConnection(discordBot *discordgo.Session) {
	errOpen := discordBot.Open()

	if errOpen != nil {
		panic(errOpen)
	}
}

func CreateNewBot(discordSecret string) *discordgo.Session {
	discordBot, err := discordgo.New("Bot " + discordSecret)

	if err != nil {
		panic(err)
	}

	log.Println("Succesfully created new Session")

	return discordBot
}
