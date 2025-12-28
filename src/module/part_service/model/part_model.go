package part_model

import (
	"time"

	"gorm.io/gorm"
)

type Part struct {
	Id          int64           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string          `json:"name" gorm:"type:varchar(255);not null"`
	DeviceId    int64           `json:"device_id" gorm:"not null;index"`
	SupplierId  int64           `json:"supplier_id" gorm:"not null;index"`
	MasterId    int64           `json:"master_id" gorm:"not null;index"`
	IncomePrice float64         `json:"income_price" gorm:"type:decimal(10,2);not null"`
	Price       float64         `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedAt   *time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
} // @name Part
