package part_model

import (
	"time"

	"gorm.io/gorm"
)

type Part struct {
	ID         int64           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string          `json:"name" gorm:"type:varchar(255);not null"`
	DeviceID   int64           `json:"device_id" gorm:"not null;index"`
	SupplierID int64           `json:"supplier_id" gorm:"not null;index"`
	CreatedAt  *time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
} // @name Part
