package bot_model

import "time"

type BotUser struct {
	Id           int64      `json:"id" gorm:"primaryKey"`
	TelegramId   int64      `json:"telegramId" gorm:"unique"`
	Username     string     `json:"username" gorm:"unique"`
	Name         string     `json:"name"`
	PhoneNumber  string     `json:"phoneNumber"`
	LanguageCode string     `json:"languageCode"`
	CreatedAt    *time.Time `json:"createdAt" gorm:"autoCreateTime:true;default:now()"`
} // @name BotUser

func (*BotUser) TableName() string {
	return "bot_users"
}
