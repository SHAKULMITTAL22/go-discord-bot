package bot

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

// TestRunCommand_a2616f765a is a test function for the RunCommand function,
// it covers both the success and failure scenarios
func TestRunCommand_a2616f765a(t *testing.T) {
	// Creating a channel to receive the result
	result := make(chan string)

	// Creating a discord message
	m := &discordgo.Message{
		Content: "Test message",
	}

	// Running the command
	go RunCommand(m, result)

	// Checking the result
	res := <-result
	if res != m.Content {
		t.Error("Expected ", m.Content, ", but got ", res)
		t.Log("Failed! Arguments were: ", m.Content)
	} else {
		t.Log("Success!")
	}

	// Creating a discord message with empty content
	m = &discordgo.Message{
		Content: "",
	}

	// Running the command
	go RunCommand(m, result)

	// Checking the result
	res = <-result
	if res != m.Content {
		t.Error("Expected ", m.Content, ", but got ", res)
		t.Log("Failed! Arguments were: ", m.Content)
	} else {
		t.Log("Success!")
	}
}
