package problem_model

import "time"

type Problem struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Price     float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
} // @name Problem
