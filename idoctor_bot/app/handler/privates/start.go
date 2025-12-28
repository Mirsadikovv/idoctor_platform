package handlers

import (
	"log"
	"strconv"

	"github.com/Mirsadikovv/idoctor_bot/app/config"
	"github.com/Mirsadikovv/idoctor_bot/app/utils"

	"fmt"

	keyboard "github.com/Mirsadikovv/idoctor_bot/app/keyboards/defaults"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func clauseOnConflict() clause.OnConflict {
	return clause.OnConflict{
		Columns: []clause.Column{{Name: "telegram_id"}},
		DoUpdates: clause.Assignments(map[string]any{
			"telegram_id":       gorm.Expr("excluded.telegram_id"),
			"telegram_username": gorm.Expr("excluded.telegram_username"),
			"telegram_name":     gorm.Expr("excluded.telegram_name"),
			"language_code":     gorm.Expr("excluded.language_code"),
		}),
	}
}

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update, cfg *config.Config, db *gorm.DB, lang *utils.LanguageCache) *utils.LanguageCache {

	type TelegramUser struct {
		TelegramId       int64  `gorm:"column:telegram_id"`
		TelegramUsername string `gorm:"column:telegram_username"`
		TelegramName     string `gorm:"column:telegram_name"`
		LanguageCode     string `gorm:"column:language_code"`
	}

	var user TelegramUser
	{
		telegramId := update.Message.From.ID
		telegramUsername := update.Message.From.UserName
		telegramName := update.Message.From.FirstName
		languageCode := update.Message.From.LanguageCode

		newUser := TelegramUser{
			TelegramId:       telegramId,
			TelegramUsername: telegramUsername,
			TelegramName:     telegramName,
			LanguageCode:     languageCode,
		}

		result := db.Table("users").
			Clauses(clauseOnConflict()).
			Create(&newUser)
		{
			if result.Error != nil {
				log.Println("Error creating telegram user:", result.Error)
			}
		}

		if result.RowsAffected == 1 {
			lang.Set(update.Message.From.ID, languageCode)

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
			if user.LanguageCode != "" {
				lang.Set(update.Message.From.ID, user.LanguageCode)
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
