package handlers

import (
	"log"

	keyboard "github.com/Mirsadikovv/idoctor_bot/app/keyboards/defaults"
	"github.com/Mirsadikovv/idoctor_bot/app/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Menu(bot *tgbotapi.BotAPI, update tgbotapi.Update, lang *utils.LanguageCache) *utils.LanguageCache {

	{
		userLang := lang.Get(update.Message.From.ID)
		menuKeyboard := keyboard.GetMainMenuKeyboard(userLang, update.Message.From.ID)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, BackChooseAction[userLang])
		msg.ParseMode = "HTML"
		msg.ReplyMarkup = utils.MakeReplyMarkup(menuKeyboard)

		if _, err := bot.Send(msg); err != nil {
			log.Println("Error sending first essage:", err)
		}
	}

	return lang
}
