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