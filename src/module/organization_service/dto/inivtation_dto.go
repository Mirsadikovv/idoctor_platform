package organization_dto

import (
	"time"

	"github.com/Mirsadikovv/idoctor_platform/src/common/enum"
	"github.com/Mirsadikovv/shared/response"
	"github.com/Mirsadikovv/shared/sharedutil"
)

type InvitationPage = response.PageData[InvitationDto]

type InvitationCreate struct {
	OrganizationId int64  `json:"organizationId"`
	Pin            int64  `json:"pin"`
	RoleId         int64  `json:"roleId"`
	Description    string `json:"description"`
} // @name Invitation

type InvitationDto struct {
	Id             int64                 `json:"id"`
	OrganizationId int64                 `json:"organizationId" gorm:"column:organization_id"`
	Organization   sharedutil.JsonObject `json:"organization"`
	Pin            int64                 `json:"pin"`
	RoleId         int64                 `json:"roleId"`
	Role           sharedutil.JsonObject `json:"role"`
	User           sharedutil.JsonObject `json:"user"`
	Description    string                `json:"description"`
	BlockedAt      *time.Time            `json:"blockedAt"`
} // @name InvitationDto

type InvitationParams struct {
	OrganizationId *int64    `query:"organization_id" json:"organization_id"`
	Pin            *string   `query:"pin" json:"pin"`
	RoleId         *int64    `query:"role_id" json:"role_id"`
	LanguageId     *int64    `query:"language_id" json:"language_id"`
	UserId         *int64    `query:"user_id" json:"user_id"`
	Blocked        enum.Bool `query:"blocked" json:"blocked"`
} // @name InvitationParams
