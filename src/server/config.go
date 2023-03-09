package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DiscordWebhookUrl string
}

func loadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	getConfigValue := func(name string) string {
		val := os.Getenv(name)

		if val == "" {
			log.Fatalf("Missing .env config for %v\n", name)
		}

		return val
	}

	return Config{
		DiscordWebhookUrl: getConfigValue("DISCORD_WEBHOOK_URL"),
	}
}
