package handlers

import (
	"fmt"
	"log"

	keyboard "github.com/Mirsadikovv/idoctor_bot/app/keyboards/defaults"
	"github.com/Mirsadikovv/idoctor_bot/app/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func SendPhoneNumber(bot *tgbotapi.BotAPI, update tgbotapi.Update, db *gorm.DB, lang *utils.LanguageCache) *utils.LanguageCache {
	if update.Message.Contact != nil && update.Message.Contact.PhoneNumber != "" {

		if err := db.Table("bot_users").
			Where("telegram_id = ?", update.Message.From.ID).
			Update("phone_number", update.Message.Contact.PhoneNumber).Error; err != nil {
			log.Println("Error updating bot user state:", err)
		}

		userLang := lang.Get(update.Message.From.ID)

		menuKeyboard := keyboard.GetMainMenuKeyboard(userLang, update.Message.From.ID)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(MenuText[userLang], update.Message.From.ID, update.Message.From.FirstName))
		msg.ParseMode = "HTML"
		msg.ReplyMarkup = utils.MakeReplyMarkup(menuKeyboard)
		if _, err := bot.Send(msg); err != nil {
			log.Println("Error sending message:", err)
		}
	} else {
		Start(bot, update, nil, db, lang)
	}

	return lang
}
