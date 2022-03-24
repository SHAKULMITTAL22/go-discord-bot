package bot

import(
	"github.com/helltf/go-discord-bot/env"
	"github.com/bwmarrin/discordgo"
)

var (
	botToken = "DISCORD_BOT_TOKEN"
)

func CreateBot() *discordgo.Session {
	discordSecret := env.GetEnvVariable(botToken)
	bot, err := discordgo.New("Bot " + discordSecret )
	
	if(err != nil){
		return nil
	}

	return bot
}
