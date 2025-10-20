package language_dto

import (
	"time"

	"git.sriss.uz/shared/shared_service/response"
	"git.sriss.uz/shared/shared_service/sharedutil"
)

type LanguagePage = response.PageData[Language]

type Language struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	DeletedAt   *time.Time `json:"deletedAt"`
}

type LanguageCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type LanguageUpdate struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type LanguageContentPage = response.PageData[LanguageContent]
type ContentPage = response.PageData[ContentDto]

type LanguageContent struct {
	ID         int64     `json:"id"`
	Key        string    `json:"key"`
	Value      string    `json:"value"`
	Category   string    `json:"category"`
	LanguageId int64     `json:"languageId"`
	Language   *Language `json:"language,omitempty"`
} // @name LanguageContent

type LanguageContentByKeyResp struct {
	Key      string               `json:"key"`
	Category string               `json:"category"`
	Contents sharedutil.JsonArray `json:"contents" gorm:"column:contents"`
} // @name LanguageContentByKey

type ContentDto struct {
	Id         int64                 `json:"id"`
	Key        string                `json:"key"`
	Value      string                `json:"value"`
	LanguageId int64                 `json:"languageId"`
	Language   sharedutil.JsonObject `json:"language,omitempty"`
} // @name ContentDto

type LanguageContentCreate struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Category   string `json:"category"`
	LanguageId int64  `json:"languageId"`
}

type LanguageContentUpdate struct {
	Key      *string `json:"key"`
	Value    *string `json:"value"`
	Category *string `json:"category"`
}

type LanguageContentsCreateOrUpdate struct {
	LanguageId int64                    `json:"languageId"`
	Category   string                   `json:"category"`
	Contents   []*LanguageContentCreate `json:"contents"`
}

type LanguageContentsUpdate struct {
	LanguageId int64  `json:"languageId"`
	Category   string `json:"category"`
	Contents   []*LanguageContentUpdate
}

type LanguageContentDelete struct {
	Key        string `param:"key"`
	Category   string `param:"category"`
	LanguageId int64  `param:"language_id"`
}

type LanguageResult struct {
	Category   string `json:"category"`
	LanguageId int64  `json:"languageId"`
}

type GetByKeyParams struct {
	Key *string `query:"key" json:"key" validate:"required"`
} // @name GetByKeyParams

type PageByLangParams struct {
	Key   *string `query:"key" json:"key"`
	Value *string `query:"value" json:"value"`
} // @name PageByLangParams
