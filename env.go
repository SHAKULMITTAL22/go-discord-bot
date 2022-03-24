package main

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {
	return os.Getenv(key)
}

func LoadEnv() bool {
	err := godotenv.Load(".env")
	return err != nil
}
