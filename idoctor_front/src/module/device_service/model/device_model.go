package device_model

import (
	"time"

	"gorm.io/gorm"
)

type Device struct {
	Id        int64           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string          `json:"name" gorm:"type:varchar(255);not null"`
	BrandName string          `json:"brand_name" gorm:"type:varchar(255)"`
	CreatedAt *time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
} // @name Device
