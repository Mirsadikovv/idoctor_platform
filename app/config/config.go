package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
	Debug    bool
	AdminIds []string
	Timeout  int
}

func Load() (*Config, error) {
	_ = godotenv.Load(".env")

	adminEnv := os.Getenv("ADMINS")
	admins := strings.Split(adminEnv, ",")

	return &Config{
		BotToken: os.Getenv("BOT_TOKEN"),
		AdminIds: admins,
	}, nil
}
