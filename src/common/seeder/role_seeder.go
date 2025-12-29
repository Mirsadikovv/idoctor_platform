package seeder

import (
	"log"

	role_model "github.com/Mirsadikovv/idoctor_platform/src/module/role_service/model"
	"github.com/Mirsadikovv/shared/sharedutil"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	// Permissions structure example:
	// {
	//   "/api/v1/role/:id": ["GET"],
	//   "/api/v1/user/:id": ["GET", "PATCH"],
	//   "/api/v1/order/page": ["GET"],
	//   ...
	// }
	//
	// Admin role has full access (no permission check)
	// Other roles must have explicit permissions defined

	roles := []role_model.Role{
		{
			ID:          1,
			Name:        "admin",
			Description: "Администратор системы с полным доступом",
			Pages:       sharedutil.JsonObject{},
			Permissions: sharedutil.JsonObject{}, // Empty - admin bypasses permission check
		},
		{
			ID:          2,
			Name:        "user",
			Description: "Пользователь системы",
			Pages:       sharedutil.JsonObject{},
			// Example permissions for user role
			Permissions: sharedutil.JsonObject{
				"/api/v1/user/:id":    []string{"GET", "PATCH"},
				"/api/v1/order/page":  []string{"GET"},
				"/api/v1/order/:id":   []string{"GET"},
				"/api/v1/part/search": []string{"GET"},
			},
		}}

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
			// Роль существует, обновляем name, description и permissions
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
