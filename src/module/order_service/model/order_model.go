package order_model

import (
	"time"

	part_model "github.com/Mirsadikovv/idoctor_platform/src/module/part_service/model"
	problem_model "github.com/Mirsadikovv/idoctor_platform/src/module/problem_service/model"
	user_model "github.com/Mirsadikovv/idoctor_platform/src/module/user_service/model"
	"gorm.io/gorm"
)

type Order struct {
	Id            int64                   `json:"id" gorm:"primaryKey;autoIncrement"`
	ClientId      *int64                  `json:"client_id" gorm:"index"`
	MasterId      *int64                  `json:"master_id" gorm:"index"`
	Price         float64                 `json:"price" gorm:"type:decimal(10,2);not null"`
	Status        string                  `json:"status" gorm:"type:varchar(50);not null"`
	PaymentType   string                  `json:"payment_type" gorm:"type:varchar(50);not null"`
	PaymentStatus string                  `json:"payment_status" gorm:"type:varchar(50);not null"`
	CreatedAt     *time.Time              `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt     *gorm.DeletedAt         `json:"deleted_at" gorm:"index"`
	Client        *user_model.User        `json:"client,omitempty" gorm:"foreignKey:ClientId;references:Id"`
	Master        *user_model.User        `json:"master,omitempty" gorm:"foreignKey:MasterId;references:Id"`
	Parts         []part_model.Part       `json:"parts" gorm:"many2many:order_parts;"`
	Problems      []problem_model.Problem `json:"problems" gorm:"many2many:order_problems;"`
} // @name Order
