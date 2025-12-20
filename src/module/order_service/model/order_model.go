package order_model

import (
	"time"

	part_model "github.com/Mirsadikovv/idoctor_platform/src/module/part_service/model"
	problem_model "github.com/Mirsadikovv/idoctor_platform/src/module/problem_service/model"
	"gorm.io/gorm"
)

type Order struct {
	ID            int64                   `json:"id" gorm:"primaryKey;autoIncrement"`
	ClientID      *int64                  `json:"client_id" gorm:"index"`
	MasterID      *int64                  `json:"master_id" gorm:"index"`
	Price         float64                 `json:"price" gorm:"type:decimal(10,2);not null"`
	Status        string                  `json:"status" gorm:"type:varchar(50);not null"`
	PaymentType   string                  `json:"payment_type" gorm:"type:varchar(50);not null"`
	PaymentStatus string                  `json:"payment_status" gorm:"type:varchar(50);not null"`
	CreatedAt     *time.Time              `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt     *gorm.DeletedAt         `json:"deleted_at" gorm:"index"`
	Parts         []part_model.Part       `json:"parts" gorm:"many2many:order_parts;"`
	Problems      []problem_model.Problem `json:"problems" gorm:"many2many:order_problems;"`
} // @name Order
