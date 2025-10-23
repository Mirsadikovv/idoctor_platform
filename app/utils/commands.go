package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetBotCommands() []tgbotapi.BotCommand {
	return []tgbotapi.BotCommand{
		{Command: "/start", Description: "Начало"},
		{Command: "/help", Description: "Помощь"},
	}
}
