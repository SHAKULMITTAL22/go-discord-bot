package bot

import (
	"strings"
	"testing"
)

// TestHasPrefix_698ccebb85 is a test function to test the HasPrefix function in the bot package.
func TestHasPrefix_698ccebb85(t *testing.T) {
	// Test case 1: Message with prefix
	message := "!bot Hello, World!"
	expected := true
	actual := HasPrefix(message)
	if actual != expected {
		t.Errorf("HasPrefix failed for message: %s. Expected %v but got %v", message, expected, actual)
	} else {
		t.Logf("HasPrefix succeeded for message: %s", message)
	}

	// Test case 2: Message without prefix
	message = "Hello, World!"
	expected = false
	actual = HasPrefix(message)
	if actual != expected {
		t.Errorf("HasPrefix failed for message: %s. Expected %v but got %v", message, expected, actual)
	} else {
		t.Logf("HasPrefix succeeded for message: %s", message)
	}

	// Test case 3: Message with prefix but not at the beginning
	message = "Hello, !bot World!"
	expected = false
	actual = HasPrefix(message)
	if actual != expected {
		t.Errorf("HasPrefix failed for message: %s. Expected %v but got %v", message, expected, actual)
	} else {
		t.Logf("HasPrefix succeeded for message: %s", message)
	}

	// Test case 4: Empty message
	message = ""
	expected = false
	actual = HasPrefix(message)
	if actual != expected {
		t.Errorf("HasPrefix failed for message: %s. Expected %v but got %v", message, expected, actual)
	} else {
		t.Logf("HasPrefix succeeded for message: %s", message)
	}
}
