package bot

import "github.com/helltf/go-discord-bot/env"


var(
	envPrefix = "BOT_PREFIX"
)
func GetPrefix() string {
	return env.GetEnvVariable(envPrefix)
}