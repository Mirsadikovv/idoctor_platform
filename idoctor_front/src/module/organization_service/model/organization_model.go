package organization_model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Organization struct {
	Id           int64         `json:"id" gorm:"primaryKey"`
	ParentId     int64         `json:"parentId" gorm:"default:0"`
	SoatoId      int64         `json:"soatoId"`
	OrgRoles     pq.Int64Array `json:"orgRoles" gorm:"type:bigint[]"`
	RegionId     int64         `json:"regionId"`
	DistrictId   int64         `json:"districtId"`
	QuarterId    int64         `json:"quarterId"`
	ReviewRoleId int64         `json:"reviewRoleId"`
} // @name Organization

func (Organization) TableName() string {
	return "organizations"
}

type OrganizationTranslation struct {
	Id             int64         `json:"id" gorm:"primaryKey"`
	OrganizationId int64         `json:"organizationId" gorm:"not null;uniqueIndex:idx_organization_lang"`
	Organization   *Organization `json:"organization,omitempty" gorm:"foreignKey:OrganizationId"`
	LanguageId     int64         `json:"languageId" gorm:"not null;uniqueIndex:idx_organization_lang"`
	Name           string        `json:"name" gorm:"default:null"`
	Description    string        `json:"description"`
} // @name OrganizationTranslation

func (OrganizationTranslation) TableName() string {
	return "organization_translations"
}

type Invitation struct {
	Id             int64           `json:"id" gorm:"primaryKey"`
	OrganizationId int64           `json:"organizationId" gorm:"not null; index:idx_organization_id_pin"`
	Pin            int64           `json:"pin" gorm:"not null; index:idx_organization_id_pin"`
	RoleId         int64           `json:"roleId" gorm:"not null"`
	Description    string          `json:"description"`
	BlockedAt      *gorm.DeletedAt `json:"blockedAt" gorm:"default:null"`
} // @name Invitation

func (Invitation) TableName() string {
	return "invitations"
}
