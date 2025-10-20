package language_model

import "gorm.io/gorm"

type Language struct {
	Id          int64           `json:"id" gorm:"primaryKey"`
	Name        string          `json:"name" gorm:"unique;not null"`
	Description string          `json:"description"`
	DeletedAt   *gorm.DeletedAt `json:"deletedAt"`
}

func (Language) TableName() string {
	return "languages"
}

type LanguageContent struct {
	Id         int64     `json:"id" gorm:"primaryKey"`
	Key        string    `json:"key" gorm:"not null;uniqueIndex:idx_language_key_category"`
	Value      string    `json:"value" gorm:"not null"`
	Category   string    `json:"category" gorm:"not null;uniqueIndex:idx_language_key_category; default:GLOBAL"`
	LanguageId int64     `json:"languageId" gorm:"not null;uniqueIndex:idx_language_key_category"`
	Language   *Language `json:"language,omitempty" gorm:"foreignKey:LanguageId"`
}

func (LanguageContent) TableName() string {
	return "language_contents"
}
