package main

import (
	"github.com/helltf/go-discord-bot/bot"
)

func main() {
	discord_secret := GetEnvVariable("DISCORD_SECRET")
	bot.StartBot(discord_secret)
}
