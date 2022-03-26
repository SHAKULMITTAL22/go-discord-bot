package bot

import "strings"

func handleMessage() string {
	return ""
}

func HasPrefix(message string) bool {
	prefix := GetPrefix()
	return strings.HasPrefix(message, prefix)
}