package handlers

import (
	"log"

	"github.com/Mirsadikovv/idoctor_bot/app/config"
	"github.com/Mirsadikovv/idoctor_bot/app/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func Register(bot *tgbotapi.BotAPI, db *gorm.DB) {
	commands := utils.GetBotCommands()
	_, err := bot.Request(tgbotapi.SetMyCommandsConfig{
		Commands: commands,
	})
	if err != nil {
		log.Println("Failed to register commands:", err)
	}
}

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update, cfg *config.Config, db *gorm.DB, langCache *utils.LanguageCache) *utils.LanguageCache {

	if update.Message.From.LanguageCode != "" && langCache.Get(update.Message.From.ID) == "" {
		langCache.Set(update.Message.From.ID, update.Message.From.LanguageCode)
	}

	if update.Message.Contact != nil && update.Message.Contact.PhoneNumber != "" {
		return SendPhoneNumber(bot, update, db, langCache)

	} else if update.Message != nil && update.Message.IsCommand() {

		switch update.Message.Command() {

		case "start":
			return Start(bot, update, cfg, db, langCache)

		case "help":
			Echo(bot, update, cfg, langCache)

		default:
			Echo(bot, update, cfg, langCache)
		}
	} else if update.Message != nil {
		switch update.Message.Text {

		case "Назад ⬅️", "Ortga ⬅️", "Back ⬅️":
			return Back(bot, update, cfg, db, langCache)

		case "Change language🌐", "Tilni o'zgartirish🌐", "Поменять язык🌐":
			return ChooseLanguage(bot, update, langCache)

		case "🇬🇧 English", "🇷🇺 Русский", "🇺🇿 O'zbek":
			return ChangeLanguage(bot, update, db, langCache)

		default:
			Echo(bot, update, cfg, langCache)
		}
	}
	return langCache
}
