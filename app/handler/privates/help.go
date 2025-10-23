package handlers

import (
	"github.com/Mirsadikovv/idoctor_v2/app/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Help(bot *tgbotapi.BotAPI, update tgbotapi.Update, lang *utils.LanguageCache) *utils.LanguageCache {

	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, HelpText[lang.Get(update.Message.From.ID)]))
	return lang
}
