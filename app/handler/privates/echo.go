package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Mirsadikovv/idoctor_bot/app/config"
	"github.com/Mirsadikovv/idoctor_bot/app/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Echo(bot *tgbotapi.BotAPI, update tgbotapi.Update, cfg *config.Config, lang *utils.LanguageCache) *utils.LanguageCache {
	// Пример отправки фото по URL

	{
		txt := "Вы написали: " + update.Message.Text +
			"\n\nНеизвестная команда\n\n\n" +
			"Bot is in development. Please wait for updates.\n\n\n" +
			"Бот находится в разработке. Пожалуйста, ждите обновлений.\n\n\n" +
			"Bot ishlab chiqilmoqda. Iltimos, yangilanishlarni kuting."

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, txt)
		removeKeyboard := tgbotapi.NewRemoveKeyboard(true)
		msg.ReplyMarkup = removeKeyboard

		if _, err := bot.Send(msg); err != nil {
			log.Println("Failed to send msg")
			return lang
		}
	}

	for _, idStr := range cfg.AdminIds {

		adminId, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			log.Printf("Invalid admin ID: %v", err)
			continue
		}

		link := fmt.Sprintf(
			UserEhoToAdmin[lang.Get(update.Message.From.ID)],
			update.Message.From.ID,
			update.Message.From.FirstName,
			update.Message.From.UserName,
			update.Message.Text,
		)

		msg := tgbotapi.NewMessage(adminId, link)
		msg.ParseMode = "HTML"

		_, err = bot.Send(msg)
		if err != nil {
			log.Printf("Failed to send admin notification: %v", err)
		}
	}

	return lang
}
