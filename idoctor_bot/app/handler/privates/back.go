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

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, BackChooseAction[lang.Get(update.Message.From.ID)])
	msg.ReplyMarkup = utils.MakeReplyMarkup(keyboard.MainMenuKeyboardMap[lang.Get(update.Message.From.ID)])

	if _, err := bot.Send(msg); err != nil {
		log.Println("Error sending message:", err)
	}

	return lang
}
