package main

import "github.com/helltf/go-discord-bot/internal/bot"

func main() {
	bot.CreateNewBotConnection()
	<-make(chan struct{})
	return
}
