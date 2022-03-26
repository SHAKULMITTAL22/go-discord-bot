package bot_test

import (
	"testing"

	"github.com/helltf/go-discord-bot/bot"
	"github.com/helltf/go-discord-bot/env"
	"github.com/stretchr/testify/assert"
) 

func TestHandlerPrefixCheckHasPrefix(t *testing.T){
	env.LoadTestEnv()

	message := "!hello"
	hasPrefix := bot.HasPrefix(message)

	assert.True(t, hasPrefix)
}

func TestHandlerPrefixCheckHasNoPrefix(t *testing.T){
	env.LoadTestEnv()

	message := "hello"
	hasPrefix := bot.HasPrefix(message)

	assert.False(t, hasPrefix)
}

func TestHandlerPrefixCheckHasPrefixSpace(t *testing.T){
	env.LoadTestEnv()

	message := "! hello"
	hasPrefix := bot.HasPrefix(message)

	assert.True(t, hasPrefix)
}

