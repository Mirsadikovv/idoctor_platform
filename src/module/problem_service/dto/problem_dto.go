package problem_dto

import (
	"time"

	"github.com/Mirsadikovv/shared/response"
)

type ProblemCreate struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required,min=0"`
} // @Name ProblemCreate

type ProblemUpdate struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required,min=0"`
} // @Name ProblemUpdate

type ProblemPage = response.PageData[Problem] // @name ProblemPage

type Problem struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
} // @Name Problem

type ProblemParams struct {
	Name     *string  `query:"name" json:"name"`
	MinPrice *float64 `query:"min_price" json:"min_price"`
	MaxPrice *float64 `query:"max_price" json:"max_price"`
} // @Name ProblemParams
