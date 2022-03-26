package main

import (
	"fmt"
	"github.com/helltf/go-discord-bot/bot"
	"github.com/helltf/go-discord-bot/env"
)


func main() {
	success := env.LoadEnv("prod")

	if !success {
		fmt.Println("Error while loading env file")
	}

	bot.CreateNewBotConnection()
	<-make(chan struct{})
	return
}

