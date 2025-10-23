package utils

import (
	"log"
	"strconv"

	"github.com/Mirsadikovv/idoctor_v2/app/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NotifyAdmins(bot *tgbotapi.BotAPI, cfg *config.Config) {
	photo := "https://repository-images.githubusercontent.com/735638594/8a5e89ee-83e4-40c9-b735-5cedfe4901f7"
	caption := "Бот начал работу  /start\n\nКоманды:\n/help - Помощь\n"

	for _, idStr := range cfg.AdminIds {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			log.Printf("Invalid admin ID: %v", err)
			continue
		}

		msg := tgbotapi.NewPhoto(id, tgbotapi.FileURL(photo))
		msg.Caption = caption

		_, err = bot.Send(msg)
		if err != nil {
			log.Printf("Failed to send admin notification: %v", err)
		}
	}
}
