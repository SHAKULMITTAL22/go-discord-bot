package env

import (
	"os"

	"github.com/joho/godotenv"
)


func GetEnvVariable(key string) string {
	res := os.Getenv(key)
	return res  
}

func LoadEnv(environment string) bool {

	envPath := getEnvPath(environment)

	err := godotenv.Load(envPath)

	return err == nil
}

func getEnvPath(env string) string{
	if env == "test"{
		return "../.env.test"
	}
	return ".env"
}