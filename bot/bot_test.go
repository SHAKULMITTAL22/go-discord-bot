package bot_test

import (
	"testing"

	"github.com/helltf/go-discord-bot/bot"
	"github.com/helltf/go-discord-bot/env"
	"github.com/stretchr/testify/assert"
) 


func TestGetDiscordSecret(t *testing.T){
	env.LoadEnv("test")

	secret := bot.GetDiscordSecret()
	assert.NotEqual(t, secret, "")
}


