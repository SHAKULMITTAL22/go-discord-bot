package bot

import(
	"github.com/helltf/go-discord-bot/env"
)

var (
	botToken = "DISCORD_BOT_TOKEN"
)

func CreateBot() *Bot {
	discordSecret := env.GetEnvVariable(botToken)
	return &Bot{discordSecret, "!"}
}
