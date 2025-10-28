package bot

import (
	"log"

	"github.com/Mirsadikovv/idoctor_bot/app/config"
	handlers "github.com/Mirsadikovv/idoctor_bot/app/handler/privates"
	"github.com/Mirsadikovv/idoctor_bot/app/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func Start(cfg *config.Config, db *gorm.DB) error {
	log.Println("Starting bot...")
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return err
	}

	bot.Debug = false

	log.Printf("Бот запущен: %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 100

	handlers.Register(bot, db)

	utils.NotifyAdmins(bot, cfg)

	updates := bot.GetUpdatesChan(u)
	langCache := utils.NewLanguageCache()

	for update := range updates {
		if update.Message != nil {
			curLang := handlers.HandleUpdate(bot, update, cfg, db, langCache)
			langCache.Set(update.Message.From.ID, curLang.Get(update.Message.From.ID))
		}
	}

	return nil
}
