package role_model

import "git.sriss.uz/shared/shared_service/sharedutil"

type Role struct {
	ID          int64                 `json:"id" gorm:"primaryKey;unique"`
	Name        string                `json:"name" gorm:"unique;not null"`
	Description string                `json:"description" gorm:"not null"`
	Pages       sharedutil.JsonObject `json:"pages" gorm:"type:jsonb"`
	Permissions sharedutil.JsonObject `json:"permissions" gorm:"type:jsonb"`
} // @name Role
