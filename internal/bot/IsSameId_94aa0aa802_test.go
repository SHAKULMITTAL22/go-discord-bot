package bot

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

// TestIsSameId_94aa0aa802 is a test function for the IsSameId method
func TestIsSameId_94aa0aa802(t *testing.T) {
	// Creating dummy session and message for testing
	dummySession := &discordgo.Session{State: &discordgo.State{Ready: discordgo.Ready{User: &discordgo.User{ID: "123"}}}}
	dummyMessage := &discordgo.MessageCreate{Message: &discordgo.Message{Author: &discordgo.User{ID: "123"}}}

	// Test case where author ID and user ID are the same
	if !IsSameId(dummySession, dummyMessage) {
		t.Error("Expected true, got false")
	} else {
		t.Log("Success: Test case where author ID and user ID are the same passed")
	}

	// Changing the author ID for the next test case
	dummyMessage.Message.Author.ID = "456"

	// Test case where author ID and user ID are different
	if IsSameId(dummySession, dummyMessage) {
		t.Error("Expected false, got true")
	} else {
		t.Log("Success: Test case where author ID and user ID are different passed")
	}
}
