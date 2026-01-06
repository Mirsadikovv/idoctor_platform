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
			Description: "Мастер системы с доступом к частям системы",
			Pages:       sharedutil.JsonObject{},
			Permissions: sharedutil.JsonObject{},
		},
		{
			ID:          3,
			Name:        "client",
			Description: "Клиент системы с доступом к открытым частям системы",
			Pages:       sharedutil.JsonObject{},
			Permissions: sharedutil.JsonObject{},
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
			// Роль существует, обновляем только если изменилась
			if existingRole.Name != role.Name || existingRole.Description != role.Description {
				if err := db.Model(&existingRole).Updates(map[string]interface{}{
					"name":        role.Name,
					"description": role.Description,
				}).Error; err != nil {
					log.Printf("Ошибка при обновлении роли %s: %v\n", role.Name, err)
				} else {
					log.Printf("Роль %s успешно обновлена\n", role.Name)
				}
			} else {
				log.Printf("Роль %s уже существует и актуальна\n", role.Name)
			}
		}
	}
}
