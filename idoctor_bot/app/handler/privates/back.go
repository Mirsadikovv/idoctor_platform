package handlers

import (
	"log"

	"github.com/Mirsadikovv/idoctor_bot/app/config"
	keyboard "github.com/Mirsadikovv/idoctor_bot/app/keyboards/defaults"
	"github.com/Mirsadikovv/idoctor_bot/app/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func Back(bot *tgbotapi.BotAPI, update tgbotapi.Update, cfg *config.Config, db *gorm.DB, lang *utils.LanguageCache) *utils.LanguageCache {

	userLang := lang.Get(update.Message.From.ID)

	menuKeyboard := keyboard.GetMainMenuKeyboard(userLang, update.Message.From.ID)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, BackChooseAction[userLang])
	msg.ReplyMarkup = utils.MakeReplyMarkup(menuKeyboard)

	if _, err := bot.Send(msg); err != nil {
		log.Println("Error sending message:", err)
	}

	return lang
}
