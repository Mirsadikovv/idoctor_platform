package src

import (
	"log"

	"github.com/Mirsadikovv/idoctor_bot/app/bot"
	"github.com/Mirsadikovv/idoctor_bot/app/config"
	"github.com/Mirsadikovv/shared/pg"
)

type Env struct {
	POSTGRES_HOST      string `env:"POSTGRES_HOST"`
	POSTGRES_USER      string `env:"POSTGRES_USER"`
	POSTGRES_PASSWORD  string `env:"POSTGRES_PASSWORD"`
	POSTGRES_DB        string `env:"POSTGRES_DB"`
	POSTGRES_PORT      int    `env:"POSTGRES_PORT"`
	POSTGRES_SSL_MODE  string `env:"POSTGRES_SSL_MODE" default:"disable"`
	POSTGRES_TIME_ZONE string `env:"POSTGRES_TIME_ZONE" default:"UTC"`
}

func Exec(env *Env) {
	log.Println("Bot starting...")
	log.Printf("Database config: Host=%s, User=%s, DB=%s, Port=%d, SSLMode=%s, TimeZone=%s",
		env.POSTGRES_HOST, env.POSTGRES_USER, env.POSTGRES_DB, env.POSTGRES_PORT, env.POSTGRES_SSL_MODE, env.POSTGRES_TIME_ZONE)

	gormConfig := &pg.GormConfig{
		SkipDefaultTransaction: true,
	}

	pgConfig := pg.ConnectionConfig{
		Host:     env.POSTGRES_HOST,
		User:     env.POSTGRES_USER,
		Password: env.POSTGRES_PASSWORD,
		DBName:   env.POSTGRES_DB,
		Port:     env.POSTGRES_PORT,
		SSLMode:  env.POSTGRES_SSL_MODE,
		TimeZone: env.POSTGRES_TIME_ZONE,
	}

	log.Println("Initializing database connection...")
	db := pg.Primary(gormConfig, pgConfig)
	log.Println("Database connection established")

	log.Println("Loading bot configuration...")
	cfg, err := config.Load()
	if err != nil {
		log.Printf("Failed to load bot configuration: %v", err)
		return
	}
	log.Printf("Bot configuration loaded successfully. Token: %s", cfg.BotToken)

	log.Println("Starting bot...")
	bot.Start(cfg, db)
}
