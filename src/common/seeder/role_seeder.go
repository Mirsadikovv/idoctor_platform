package seeder

import (
	"log"

	role_model "github.com/Mirsadikovv/idoctor_platform/src/module/role_service/model"
	"github.com/Mirsadikovv/shared/sharedutil"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	roles := []role_model.Role{
		{
			ID:          1,
			Name:        "admin",
			Description: "Администратор системы с полным доступом",
			Pages:       sharedutil.JsonObject{},
			Permissions: sharedutil.JsonObject{},
		},
		{
			ID:          2,
			Name:        "master",
			Description: "Мастер системы с доступом к работе с заказами",
			Pages:       sharedutil.JsonObject{},
			Permissions: sharedutil.JsonObject{
				// Auth endpoints
				"/api/v1/auth/me":      []string{"POST"},
				"/api/v1/auth/sign-out": []string{"POST"},

				// Order endpoints - master can view and update assigned orders
				"/api/v1/order/page":   []string{"GET"},
				"/api/v1/order/search": []string{"GET"},
				"/api/v1/order/:id":    []string{"GET", "PUT"},

				// Order parts - master needs to manage parts for repairs
				"/api/v1/order-part":        []string{"POST"},
				"/api/v1/order-part/page":   []string{"GET"},
				"/api/v1/order-part/search": []string{"GET"},
				"/api/v1/order-part/:id":    []string{"GET", "PUT", "DELETE"},

				// Parts - master needs to view available parts
				"/api/v1/part/page":   []string{"GET"},
				"/api/v1/part/search": []string{"GET"},
				"/api/v1/part/:id":    []string{"GET"},

				// Problems - master needs to view and manage problems
				"/api/v1/problem/page":   []string{"GET"},
				"/api/v1/problem/search": []string{"GET"},
				"/api/v1/problem/:id":    []string{"GET"},

				// Devices - master needs to view device information
				"/api/v1/device/page":   []string{"GET"},
				"/api/v1/device/search": []string{"GET"},
				"/api/v1/device/brands": []string{"GET"},
				"/api/v1/device/:id":    []string{"GET"},

				// Suppliers - master may need to view suppliers
				"/api/v1/supplier/page":   []string{"GET"},
				"/api/v1/supplier/search": []string{"GET"},
				"/api/v1/supplier/:id":    []string{"GET"},
			},
		},
		{
			ID:          3,
			Name:        "user",
			Description: "Клиент системы с доступом к своим заказам",
			Pages:       sharedutil.JsonObject{},
			Permissions: sharedutil.JsonObject{
				// Auth endpoints
				"/api/v1/auth/me":              []string{"POST"},
				"/api/v1/auth/sign-out":        []string{"POST"},
				"/api/v1/auth/sign-in-telegram": []string{"POST"},

				// Order endpoints - user can view own orders and create new ones
				"/api/v1/order":        []string{"POST"},
				"/api/v1/order/page":   []string{"GET"},
				"/api/v1/order/search": []string{"GET"},
				"/api/v1/order/:id":    []string{"GET"},

				// Order parts - user can view parts in their orders
				"/api/v1/order-part/page":   []string{"GET"},
				"/api/v1/order-part/search": []string{"GET"},
				"/api/v1/order-part/:id":    []string{"GET"},

				// Devices - user needs to view device information when creating orders
				"/api/v1/device/page":   []string{"GET"},
				"/api/v1/device/search": []string{"GET"},
				"/api/v1/device/brands": []string{"GET"},
				"/api/v1/device/:id":    []string{"GET"},

				// Problems - user needs to select problems when creating orders
				"/api/v1/problem/page":   []string{"GET"},
				"/api/v1/problem/search": []string{"GET"},
				"/api/v1/problem/:id":    []string{"GET"},

			},
		},
	}

	for _, role := range roles {
		var existingRole role_model.Role
		err := db.Where("id = ?", role.ID).First(&existingRole).Error

		if err == gorm.ErrRecordNotFound {
			// Роль не существует, создаем новую
			if err := db.Create(&role).Error; err != nil {
				log.Printf("Ошибка при создании роли %s: %v\n", role.Name, err)
			} else {
				log.Printf("Роль %s успешно создана\n", role.Name)
			}
		} else if err != nil {
			log.Printf("Ошибка при проверке роли %s: %v\n", role.Name, err)
		} else {
			// Роль существует, обновляем
			if err := db.Model(&existingRole).Updates(map[string]interface{}{
				"name":        role.Name,
				"description": role.Description,
				"permissions": role.Permissions,
				"pages":       role.Pages,
			}).Error; err != nil {
				log.Printf("Ошибка при обновлении роли %s: %v\n", role.Name, err)
			} else {
				log.Printf("Роль %s успешно обновлена\n", role.Name)
			}
		}
	}
}
