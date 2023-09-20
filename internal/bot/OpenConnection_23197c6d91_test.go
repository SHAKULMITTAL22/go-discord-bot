package bot

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

// TestOpenConnection_23197c6d91 is a test function for OpenConnection
func TestOpenConnection_23197c6d91(t *testing.T) {
	// Test case 1: successful connection
	t.Run("Success", func(t *testing.T) {
		// TODO: replace with your actual bot token
		token := "your bot token"

		discordBot, _ := discordgo.New("Bot " + token)
		defer discordBot.Close()

		// no panic means success
		defer func() {
			if r := recover(); r != nil {
				t.Error("The code panicked")
			}
		}()

		OpenConnection(discordBot)

		t.Log("TestOpenConnection_23197c6d91 Success case passed")
	})

	// Test case 2: failed connection (invalid token)
	t.Run("Failure", func(t *testing.T) {
		// invalid token
		token := "invalid"

		discordBot, _ := discordgo.New("Bot " + token)
		defer discordBot.Close()

		// expecting a panic
		defer func() {
			if r := recover(); r == nil {
				t.Error("The code did not panic")
			}
		}()

		OpenConnection(discordBot)

		t.Log("TestOpenConnection_23197c6d91 Failure case passed")
	})
}
