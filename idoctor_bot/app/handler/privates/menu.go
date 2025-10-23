package handlers

import (
	"log"

	keyboard "github.com/Mirsadikovv/idoctor_v2/app/keyboards/defaults"
	"github.com/Mirsadikovv/idoctor_v2/app/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Menu(bot *tgbotapi.BotAPI, update tgbotapi.Update, lang *utils.LanguageCache) *utils.LanguageCache {

	{
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, BackChooseAction[lang.Get(update.Message.From.ID)])
		msg.ParseMode = "HTML"
		msg.ReplyMarkup = utils.MakeReplyMarkup(keyboard.MainMenuKeyboardMap[lang.Get(update.Message.From.ID)])

		if _, err := bot.Send(msg); err != nil {
			log.Println("Error sending first essage:", err)
		}
	}

	return lang
}
