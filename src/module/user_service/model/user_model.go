package user_model

import (
	"time"
)

type User struct {
	Id             int64      `json:"id" gorm:"primaryKey"`
	Username       string     `json:"username" gorm:"column:username;not null"`
	Password       string     `json:"-" gorm:"column:password"`
	FirstName      string     `json:"firstName" gorm:"column:first_name"`
	LastName       string     `json:"lastName" gorm:"column:last_name"`
	MiddleName     string     `json:"middleName" gorm:"column:middle_name"`
	DateOfBirth    string     `json:"dateOfBirth" gorm:"column:date_of_birth"`
	Gender         string     `json:"gender" gorm:"column:gender"`
	RoleId         int64      `json:"roleId" gorm:"column:role_id"`
	OrganizationId int64      `json:"organizationId" gorm:"column:organization_id"`
	LastVisit      *time.Time `json:"lastVisitTime" gorm:"column:last_visit"`
	CreatedAt      *time.Time `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt      *time.Time `json:"updatedAt" gorm:"autoUpdateTime:true,default:null"`
	BlockedAt      *time.Time `json:"blockedAt" gorm:"column:blocked_at"`
} //@name UserModel

func (User) TableName() string {
	return "users"
}
