package bot

import (
	"testing"
)

// TestHasPrefix_698ccebb85 is a test function for HasPrefix method
func TestHasPrefix_698ccebb85(t *testing.T) {
	// Test case 1: When message has prefix
	message1 := "prefixHello"
	if !HasPrefix(message1) {
		t.Errorf("HasPrefix failed, expected %v, got %v", true, false)
	} else {
		t.Logf("HasPrefix success, expected %v, got %v", true, true)
	}

	// Test case 2: When message does not have prefix
	message2 := "Hello"
	if HasPrefix(message2) {
		t.Errorf("HasPrefix failed, expected %v, got %v", false, true)
	} else {
		t.Logf("HasPrefix success, expected %v, got %v", false, false)
	}

	// Test case 3: When message is empty
	message3 := ""
	if HasPrefix(message3) {
		t.Errorf("HasPrefix failed, expected %v, got %v", false, true)
	} else {
		t.Logf("HasPrefix success, expected %v, got %v", false, false)
	}
}
