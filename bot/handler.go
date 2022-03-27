package bot

import (
	"strings"
	"github.com/bwmarrin/discordgo"
)

func handleMessage() string {
	return ""
}

func HasPrefix(message string) bool {
	prefix := GetPrefix()
	return strings.HasPrefix(message, prefix)
}

func IsSameId(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return m.Author.ID == s.State.User.ID
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	result := make(chan string)

	if HasPrefix(m.Content) || IsSameId(s, m){
		 return
	} 
	
	go RunCommand(m, result)
	response := <- result
	s.ChannelMessageSend(m.ChannelID, response)
}

func RunCommand(m *discordgo.MessageCreate, result chan string) {
	result <- m.Content
}

