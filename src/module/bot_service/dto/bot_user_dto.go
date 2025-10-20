package bot_dto

import (
	"git.sriss.uz/shared/shared_service/response"
)

type BotUserPage = response.PageData[BotUserDto]

type BotUserDto struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	TelegramId   int64  `json:"telegramId"`
	Username     string `json:"username"`
	PhoneNumber  string `json:"phoneNumber"`
	LanguageCode string `json:"languageCode"`
} // @name Bot

type BotUserCreate struct {
	Name         string `json:"name"`
	TelegramId   int64  `json:"telegramId"`
	Username     string `json:"username"`
	PhoneNumber  string `json:"phoneNumber"`
	LanguageCode string `json:"languageCode"`
} // @name BotCreate

type BotUserUpdate struct {
	Name         *string `json:"name"`
	TelegramId   *int64  `json:"telegramId"`
	Username     *string `json:"username"`
	PhoneNumber  *string `json:"phoneNumber"`
	LanguageCode *string `json:"languageCode"`
} // @name BotUpdate

type BotUserParams struct {
	Name        *string `query:"name" json:"name"`
	TelegramId  *int64  `query:"telegram_id" json:"telegram_id"`
	Username    *string `query:"username" json:"username"`
	PhoneNumber *string `query:"phone_number" json:"phone_number"`
} // @name BotParams
