package bot

import (
	"strings"
	"testing"

	"github.com/bwmarrin/discordgo"
)

// TestRunCommand_a2616f765a is a test function for the RunCommand function
func TestRunCommand_a2616f765a(t *testing.T) {
	// Test case 1: Normal case
	t.Run("normal case", func(t *testing.T) {
		result := make(chan string)
		message := &discordgo.MessageCreate{
			Message: &discordgo.Message{
				Content: "Hello, world!",
			},
		}
		go RunCommand(message, result)

		if r := <-result; r != "Hello, world!" {
			t.Errorf("RunCommand(%q) = %q; want %q", message.Content, r, "Hello, world!")
		} else {
			t.Logf("Success: RunCommand(%q) = %q", message.Content, r)
		}
	})

	// Test case 2: Empty message
	t.Run("empty message", func(t *testing.T) {
		result := make(chan string)
		message := &discordgo.MessageCreate{
			Message: &discordgo.Message{
				Content: "",
			},
		}
		go RunCommand(message, result)

		if r := <-result; r != "" {
			t.Errorf("RunCommand(%q) = %q; want %q", message.Content, r, "")
		} else {
			t.Logf("Success: RunCommand(%q) = %q", message.Content, r)
		}
	})
}
