package bot

import "strings"

func handleMessage() string {
	return ""
}

func CheckForPrefix(message string) bool {
	prefix := GetPrefix()
	return strings.HasPrefix(message, prefix)
}