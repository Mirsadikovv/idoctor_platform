package order_part_model

import (
	"time"

	order_model "github.com/Mirsadikovv/idoctor_platform/src/module/order_service/model"
	part_model "github.com/Mirsadikovv/idoctor_platform/src/module/part_service/model"
	supplier_model "github.com/Mirsadikovv/idoctor_platform/src/module/supplier_service/model"

	"gorm.io/gorm"
)

type OrderPart struct {
	Id          int64                    `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderId     int64                    `json:"order_id" gorm:"not null;index"`
	PartId      int64                    `json:"part_id" gorm:"not null;index"`
	SupplierId  int64                    `json:"supplier_id" gorm:"not null;index"`
	IncomePrice float64                  `json:"income_price" gorm:"type:decimal(10,2);not null"`
	Price       float64                  `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedAt   *time.Time               `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt   *gorm.DeletedAt          `json:"deleted_at" gorm:"index"`
	Order       *order_model.Order       `json:"order,omitempty" gorm:"foreignKey:OrderId;references:Id"`
	Part        *part_model.Part         `json:"part,omitempty" gorm:"foreignKey:PartId;references:Id"`
	Supplier    *supplier_model.Supplier `json:"supplier,omitempty" gorm:"foreignKey:SupplierId;references:ID"`
} // @name OrderPart

func (OrderPart) TableName() string {
	return "order_parts"
}
