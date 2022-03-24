package main

import (
	"fmt"
	"github.com/helltf/go-discord-bot/bot"
	"github.com/helltf/go-discord-bot/env"
)


func main() {
	success := env.LoadEnv()

	if !success {
		fmt.Println("Error while loading env file")
	}

	bot.CreateBot()
	<-make(chan struct{})
	return
}

