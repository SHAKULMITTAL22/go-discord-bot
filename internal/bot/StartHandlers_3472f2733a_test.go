package bot

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

// TestStartHandlers_3472f2733a is a test function for the StartHandlers function
func TestStartHandlers_3472f2733a(t *testing.T) {
	// Creating a mock bot
	mockBot := &discordgo.Session{}

	// Test case 1: Check if the function runs without error
	t.Run("Test case 1: Check if the function runs without error", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Error("The function panicked")
			}
		}()
		StartHandlers(mockBot)
		t.Log("Test case 1 passed")
	})

	// Test case 2: Check if the messageHandler is added
	t.Run("Test case 2: Check if the messageHandler is added", func(t *testing.T) {
		StartHandlers(mockBot)
		if len(mockBot.MessageCreateHandlers) != 1 {
			t.Error("messageHandler was not added")
		} else {
			t.Log("Test case 2 passed")
		}
	})

	// Test case 3: Check if the messageHandler is not added more than once
	t.Run("Test case 3: Check if the messageHandler is not added more than once", func(t *testing.T) {
		StartHandlers(mockBot)
		if len(mockBot.MessageCreateHandlers) > 1 {
			t.Error("messageHandler was added more than once")
		} else {
			t.Log("Test case 3 passed")
		}
	})
}
