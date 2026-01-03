package part_model

import (
	"time"

	device_model "github.com/Mirsadikovv/idoctor_platform/src/module/device_service/model"
	supplier_model "github.com/Mirsadikovv/idoctor_platform/src/module/supplier_service/model"
	"gorm.io/gorm"
)

type Part struct {
	Id          int64                    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string                   `json:"name" gorm:"type:varchar(255);not null"`
	DeviceId    int64                    `json:"device_id" gorm:"not null;index"`
	SupplierId  int64                    `json:"supplier_id" gorm:"not null;index"`
	IncomePrice float64                  `json:"income_price" gorm:"type:decimal(10,2)"`
	Price       float64                  `json:"price" gorm:"type:decimal(10,2)"`
	CreatedAt   *time.Time               `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt   *gorm.DeletedAt          `json:"deleted_at" gorm:"index"`
	Device      *device_model.Device     `json:"device,omitempty" gorm:"foreignKey:DeviceId;references:ID"`
	Supplier    *supplier_model.Supplier `json:"supplier,omitempty" gorm:"foreignKey:SupplierId;references:ID"`
} // @name Part
