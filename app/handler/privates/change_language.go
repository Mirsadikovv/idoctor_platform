package handlers

import (
	"log"

	keyboard "github.com/Mirsadikovv/idoctor_v2/app/keyboards/defaults"
	"github.com/Mirsadikovv/idoctor_v2/app/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func ChooseLanguage(bot *tgbotapi.BotAPI, update tgbotapi.Update, lang *utils.LanguageCache) *utils.LanguageCache {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, ChangeLanguageText[update.Message.From.LanguageCode])
	msg.ReplyMarkup = utils.MakeReplyMarkup(keyboard.LanguageKeyboard)

	if _, err := bot.Send(msg); err != nil {
		log.Println("Error sending message:", err)
	}

	return lang
}

func ChangeLanguage(bot *tgbotapi.BotAPI, update tgbotapi.Update, db *gorm.DB, langCache *utils.LanguageCache) *utils.LanguageCache {
	var lang string
	if update.Message.Text == "🇷🇺 Русский" {
		lang = "ru"
	} else if update.Message.Text == "🇺🇿 O'zbek" {
		lang = "uz"
	} else if update.Message.Text == "🇬🇧 English" {
		lang = "en"
	} else {
		ChooseLanguage(bot, update, langCache)
	}

	if err := db.Table("bot_users").
		Where("telegram_id = ?", update.Message.From.ID).
		Update("language_code", lang).Error; err != nil {
		log.Println("Error updating bot user state:", err)
	}

	langCache.Set(update.Message.From.ID, lang)

	return Menu(bot, update, langCache)
}
