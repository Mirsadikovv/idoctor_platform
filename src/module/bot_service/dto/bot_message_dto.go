package bot_dto

import (
	"time"

	appeal_dto "git.sriss.uz/mehnat/inspector_platform/src/module/appeal_service/dto"
	appeal_model "git.sriss.uz/mehnat/inspector_platform/src/module/appeal_service/model"
	"git.sriss.uz/shared/shared_service/response"
)

type BotMessagePage = response.PageData[BotMessageDto] // @name BotMessagePage

type BotMessageDto struct {
	Id         int64     `json:"id"`
	AppealId   int64     `json:"appealId"`
	TelegramId int64     `json:"telegramId"`
	Message    string    `json:"message"`
	SentAt     time.Time `json:"sentAt"`
	IsBlocked  bool      `json:"isBlocked"`
	CreatedAt  int64     `json:"createdAt"`
	UpdatedAt  int64     `json:"updatedAt"`
} // @name BotMessage

type BotMessageCreate struct {
	AppealId   int64  `json:"appealId"`
	TelegramId int64  `json:"telegramId"`
	Message    string `json:"message"`
} // @name BotMessageCreate

type BotMessageUpdate struct {
	TelegramId int64     `json:"telegramId"`
	Message    string    `json:"message"`
	SentAt     time.Time `json:"sentAt"`
	IsBlocked  bool      `json:"isBlocked"`
} // @name BotMessageUpdate

type BotMessageParams struct {
	TelegramId *int64  `query:"telegram_id" json:"telegram_id"`
	Message    *string `query:"message" json:"message"`
} // @name BotMessageParams

type SendMessageRequest struct {
	TelegramId int64  `json:"telegramId" validate:"required"`
	Text       string `json:"text" validate:"required"`
} // @name SendMessageRequest

type SendMessageByStatusRequest struct {
	TelegramId     int64
	Status         appeal_model.AppealWorkflowAction
	LatestWorkflow appeal_dto.LatestWorkflow
}
