package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/helltf/go-discord-bot/env"
)

var (
	botToken = "DISCORD_BOT_TOKEN"
)

func CreateBot()  {
	discordSecret := env.GetEnvVariable(botToken)
	discordBot, err := discordgo.New("Bot " + discordSecret )
	
	if(err != nil){
		fmt.Printf(err.Error())
	}
	discordBot.AddHandler(messageHandler)

	errOpen := discordBot.Open()

	if errOpen != nil {
		fmt.Println(errOpen.Error())
	}
}


func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate){
	fmt.Println(m.Message)
}