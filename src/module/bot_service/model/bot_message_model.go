package bot_model

import "time"

type BotMessage struct {
	Id         int64     `json:"id" gorm:"primaryKey"`
	TelegramId int64     `json:"telegramId" gorm:"index"`
	AppealId   int64     `json:"appealId" gorm:"index"`
	Message    string    `json:"message" gorm:"type:text"`
	SentAt     time.Time `json:"sentAt" gorm:"type:timestamp;default:null"`
	IsBlocked  bool      `json:"isBlocked" gorm:"default:false"`
	CreatedAt  int64     `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt  int64     `json:"updatedAt" gorm:"autoUpdateTime:true"`
}

func (*BotMessage) TableName() string {
	return "bot_messages"
}
