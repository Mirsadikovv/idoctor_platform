package supplier_model

import (
	"time"

	"gorm.io/gorm"
)

type Supplier struct {
	ID        int64           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string          `json:"name" gorm:"type:varchar(255);not null"`
	CreatedAt *time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
} // @name Supplier
