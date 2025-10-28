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

	db := pg.Primary(gormConfig, pgConfig)
	{
		if db == nil {
			log.Fatal("Failed to establish database connection")
			return
		}
	}

	cfg, err := config.Load()
	if err != nil {
		return
	}

	bot.Start(cfg, db)
}
