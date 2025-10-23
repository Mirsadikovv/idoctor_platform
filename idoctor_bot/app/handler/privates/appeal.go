package handlers

import (
	"log"

	keyboard "github.com/Mirsadikovv/idoctor_bot/app/keyboards/defaults"
	"github.com/Mirsadikovv/idoctor_bot/app/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Appeal(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Нажмите кнопку ниже, чтобы подать заявку:")
	msg.ReplyMarkup = utils.MakeReplyMarkup(keyboard.AppealKeyboard)

	if _, err := bot.Send(msg); err != nil {
		log.Println("Error sending message:", err)
	}
}
