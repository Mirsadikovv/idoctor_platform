package user_model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id             int64           `json:"id" gorm:"primaryKey"`
	Pin            int64           `json:"pin" gorm:"column:pin;uniqueIndex;not null"`
	Username       string          `json:"username" gorm:"column:username;not null"`
	Valid          string          `json:"valid" gorm:"column:valid;not null"`
	PassportNumber string          `json:"passportNumber" gorm:"column:passport_number;not null"`
	FirstName      string          `json:"firstName" gorm:"column:first_name;not null"`
	LastName       string          `json:"lastName" gorm:"column:last_name;not null"`
	MiddleName     string          `json:"middleName" gorm:"column:middle_name;not null"`
	DateOfBirth    string          `json:"dateOfBirth" gorm:"column:date_of_birth;not null"`
	UserType       string          `json:"userType" gorm:"column:user_type;not null"`
	LastVisit      *time.Time      `json:"lastVisitTime" gorm:"column:last_visit"`
	CreatedAt      *time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt      *time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true,default:null"`
	DeletedAt      *gorm.DeletedAt `json:"-" swaggerignore:"true"`
	Gender         string          `json:"gender" gorm:"column:gender"`
	Email          string          `json:"email" gorm:"column:email"`
	SessionId      string          `json:"sessionId" gorm:"column:session_id"`
	Photo          string          `json:"photo" gorm:"column:photo"`
	Nationality    string          `json:"nationality" gorm:"column:nationality"`
	PlaceOfBirth   string          `json:"placeOfBirth" gorm:"column:place_of_birth"`
	CountryOfBirth string          `json:"countryOfBirth" gorm:"column:country_of_birth"`
	Citizenship    string          `json:"citizenship" gorm:"column:citizenship"`
	RoleId         int64           `json:"roleId" gorm:"column:role_id"`
	OrganizationId int64           `json:"organizationId" gorm:"column:organization_id"`
} //@name UserModel

func (User) TableName() string {
	return "users"
}
