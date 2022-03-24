package bot

import "fmt"

type Bot struct {
	Token     string
	BotPrefix string
}

func (b Bot) Start() {
	fmt.Println("Starting")
}