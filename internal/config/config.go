package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func MustLoad() {
	configName := os.Getenv("CONFIG_PATH")
	if configName == "" {
		configName = ".env"
	}
	err := godotenv.Load(configName)
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}
}
