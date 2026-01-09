package role_dto

import (
	"github.com/Mirsadikovv/shared/response"
	"github.com/Mirsadikovv/shared/sharedutil"
)

type Permission struct {
	Uuid   string `json:"uuid"`
	Path   string `json:"path"`
	Method string `json:"method"`
} // @Name Permission

type RoleCreate struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Pages       sharedutil.JsonObject `json:"pages"`
	Permissions sharedutil.JsonObject `json:"permissions"`
} // @Name RoleCreate

type RoleUpdate struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Pages       sharedutil.JsonObject `json:"pages"`
	Permissions sharedutil.JsonObject `json:"permissions"`
} // @Name RoleUpdate

type RolePage = response.PageData[Role] // @name RolePage

type Role struct {
	ID          int64                 `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Pages       sharedutil.JsonObject `json:"pages"`
	Permissions sharedutil.JsonObject `json:"permissions"`
} // @Name Role

type RoleParams struct {
	Name           *string `query:"name" json:"name"`
	Description    *string `query:"description" json:"description"`
	OrganizationId *int64  `query:"organization_id" json:"organization_id"`
} // @Name RoleParams
