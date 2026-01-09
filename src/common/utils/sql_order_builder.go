// utils/order.go
package utils

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type OrderParams struct {
	SortBy string `json:"sort_by" query:"sort_by"`
	Order  string `json:"order" query:"order"`
}

func (o *OrderParams) Apply(tx *gorm.DB, allowed map[string]string, defaultSort string) *gorm.DB {

	fields := strings.Split(o.SortBy, ",")
	orders := strings.Split(o.Order, ",")

	added := false
	for i, field := range fields {
		field = strings.TrimSpace(field)
		if col, ok := allowed[field]; ok {
			dir := "asc"
			if i < len(orders) && strings.ToLower(strings.TrimSpace(orders[i])) == "desc" {
				dir = "desc"
			}
			tx = tx.Order(fmt.Sprintf("%s %s", col, dir))
			added = true
		}
	}

	if !added && defaultSort != "" {
		tx = tx.Order(defaultSort)
	}

	return tx
}
