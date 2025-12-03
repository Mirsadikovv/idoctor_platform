package user_model

import (
	"time"
)

type User struct {
	Id               int64      `json:"id" gorm:"primaryKey"`
	Username         string     `json:"username" gorm:"column:username"`
	Password         string     `json:"-" gorm:"column:password"`
	FirstName        string     `json:"firstName" gorm:"column:first_name"`
	LastName         string     `json:"lastName" gorm:"column:last_name"`
	MiddleName       string     `json:"middleName" gorm:"column:middle_name"`
	DateOfBirth      string     `json:"dateOfBirth" gorm:"column:date_of_birth"`
	Gender           string     `json:"gender" gorm:"column:gender"`
	RoleId           int64      `json:"roleId" gorm:"column:role_id"`
	OrganizationId   int64      `json:"organizationId" gorm:"column:organization_id"`
	TelegramName     *string    `json:"telegramName,omitempty" gorm:"column:telegram_name"`
	TelegramId       *int64     `json:"telegramId,omitempty" gorm:"column:telegram_id;unique"`
	TelegramUsername *string    `json:"telegramUsername,omitempty" gorm:"column:telegram_username"`
	PhoneNumber      *string    `json:"phoneNumber,omitempty" gorm:"column:phone_number"`
	LanguageCode     *string    `json:"languageCode,omitempty" gorm:"column:language_code"`
	LastVisit        *time.Time `json:"lastVisitTime" gorm:"column:last_visit"`
	CreatedAt        *time.Time `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt        *time.Time `json:"updatedAt" gorm:"autoUpdateTime:true,default:null"`
	BlockedAt        *time.Time `json:"blockedAt" gorm:"column:blocked_at"`
} //@name UserModel

func (User) TableName() string {
	return "users"
}
