package order_model

import (
	"time"

	problem_model "github.com/Mirsadikovv/idoctor_platform/src/module/problem_service/model"
	"gorm.io/gorm"
)

type Order struct {
	Id            int64                   `json:"id" gorm:"primaryKey;autoIncrement"`
	ClientId      *int64                  `json:"client_id" gorm:"index"`
	ClientPhone   string                  `json:"client_phone" gorm:"type:varchar(20)"`
	ClientName    string                  `json:"client_name" gorm:"type:varchar(255)"`
	PhonePassword string                  `json:"phone_password" gorm:"type:varchar(50)"`
	MasterId      *int64                  `json:"master_id" gorm:"index"`
	Price         float64                 `json:"price" gorm:"type:decimal(10,2);not null"`
	Status        string                  `json:"status" gorm:"type:varchar(50);not null"`
	PaymentType   string                  `json:"payment_type" gorm:"type:varchar(50);not null"`
	PaymentStatus string                  `json:"payment_status" gorm:"type:varchar(50);not null"`
	Deadline      *time.Time              `json:"deadline" gorm:"type:timestamp"`
	CreatedAt     *time.Time              `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt     *gorm.DeletedAt         `json:"deleted_at" gorm:"index"`
	Problems      []problem_model.Problem `json:"problems" gorm:"many2many:order_problems;"`
} // @name Order
