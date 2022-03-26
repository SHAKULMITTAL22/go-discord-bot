package env_test

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/helltf/go-discord-bot/env"
)


func TestLoadProdEnv(t *testing.T){
	success := env.LoadEnv("prod")
	assert.False(t, success)
}
