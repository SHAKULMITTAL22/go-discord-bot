package bot_test

import (
	"testing"

	"github.com/helltf/go-discord-bot/bot"
	"github.com/helltf/go-discord-bot/env"
	"github.com/stretchr/testify/assert"
) 


func TestGetDiscordSecret(t *testing.T){
	env.LoadTestEnv()

	secret := bot.GetDiscordSecret()
	assert.NotEqual(t, secret, "")
}


func TestGetPrefix(t *testing.T){
	env.LoadTestEnv()
	prefix := bot.GetPrefix()

	assert.Equal(t, prefix, "!")
}