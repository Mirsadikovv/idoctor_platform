package handlers

import (
	"log"
	"strconv"

	"github.com/Mirsadikovv/idoctor_v2/app/config"
	"github.com/Mirsadikovv/idoctor_v2/app/utils"

	"fmt"

	bot_dto "git.sriss.uz/mehnat/inspector_platform/src/module/bot_service/dto"
	keyboard "github.com/Mirsadikovv/idoctor_v2/app/keyboards/defaults"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update, cfg *config.Config, db *gorm.DB, lang *utils.LanguageCache) *utils.LanguageCache {

	var botDto bot_dto.BotUserCreate
	{

		{
			botDto.TelegramId = update.Message.From.ID
			botDto.Name = update.Message.From.FirstName
			botDto.Username = update.Message.From.UserName
			botDto.LanguageCode = update.Message.From.LanguageCode
		}

		result := db.
			Table("bot_users").
			Where("telegram_id = ?", update.Message.From.ID).
			FirstOrCreate(&botDto)

		if result.Error != nil {
			log.Println("Error creating bot user:", result.Error)
		}

		if result.RowsAffected == 1 {

			lang.Set(update.Message.From.ID, update.Message.From.LanguageCode)

			for _, idStr := range cfg.AdminIds {

				adminId, err := strconv.ParseInt(idStr, 10, 64)
				if err != nil {
					log.Printf("Invalid admin ID: %v", err)
					continue
				}

				link := fmt.Sprintf(
					StartUserAdded[update.Message.From.LanguageCode],
					update.Message.From.ID,
					update.Message.From.FirstName,
					update.Message.From.UserName,
				)

				msg := tgbotapi.NewMessage(adminId, link)
				msg.ParseMode = "HTML"

				_, err = bot.Send(msg)
				if err != nil {
					log.Printf("Failed to send admin notification: %v", err)
				}
			}
		} else {
			if botDto.LanguageCode != "" {
				lang.Set(update.Message.From.ID, botDto.LanguageCode)
			}
		}
	}

	{
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, StartWriteYourPhone[lang.Get(update.Message.From.ID)])
		msg.ReplyMarkup = utils.MakeReplyMarkup(keyboard.SendContactMap[lang.Get(update.Message.From.ID)])

		if _, err := bot.Send(msg); err != nil {
			log.Println("Error sending message:", err)
		}

	}

	return lang
}
