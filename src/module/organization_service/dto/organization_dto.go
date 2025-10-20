package organization_dto

import (
	"git.sriss.uz/shared/shared_service/response"
	"git.sriss.uz/shared/shared_service/sharedutil"
	"github.com/lib/pq"
)

type OrganizationPage = response.PageData[Organization]

type Organization struct {
	Id           int64          `json:"id"`
	ParentId     *int64         `json:"parentId"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	SoatoId      int64          `json:"soatoId"`
	Address      string         `json:"address"`
	OrgRoles     *pq.Int64Array `json:"orgRoles" gorm:"type:bigint[]"`
	ReviewRoleId *int64         `json:"reviewRoleId"`
	RegionId     int64          `json:"regionId"`
	DistrictId   int64          `json:"districtId"`
	QuarterId    int64          `json:"quarterId"`
	RegionName   string         `json:"regionName" gorm:"region_name"`
	DistrictName string         `json:"districtName" gorm:"district_name"`
	QuarterName  string         `json:"quarterName" gorm:"quarter_name"`
	SoatoName    string         `json:"soatoName" gorm:"soato_name"`
} // @name Organization

type OrganizationDto struct {
	Id           int64                 `json:"id"`
	ParentId     *int64                `json:"parentId"`
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	LanguageId   int64                 `json:"languageId"`
	SoatoId      int64                 `json:"soatoId"`
	Address      string                `json:"address"`
	OrgRoles     sharedutil.JsonArray  `json:"orgRoles" gorm:"column:org_roles;type:jsonb"`
	ReviewRoleId *int64                `json:"reviewRoleId"`
	RegionId     int64                 `json:"regionId"`
	DistrictId   int64                 `json:"districtId"`
	QuarterId    int64                 `json:"quarterId"`
	RegionName   string                `json:"regionName" gorm:"region_name"`
	DistrictName string                `json:"districtName" gorm:"district_name"`
	QuarterName  string                `json:"quarterName" gorm:"quarter_name"`
	SoatoName    string                `json:"soatoName" gorm:"soato_name"`
	Parent       sharedutil.JsonObject `json:"parent" gorm:"column:parent_organization"`
	ReviewRole   sharedutil.JsonObject `json:"reviewRole" gorm:"column:review_role"`
} // @name OrganizationDto

type OrganizationCreate struct {
	ParentId     int64         `json:"parentId"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	SoatoId      int64         `json:"soatoId"`
	OrgRoles     pq.Int64Array `json:"orgRoles"`
	ReviewRoleId int64         `json:"reviewRoleId"`
	LanguageId   int64         `json:"languageId"`
	RegionId     int64         `json:"regionId"`
	DistrictId   int64         `json:"districtId"`
	QuarterId    int64         `json:"quarterId"`
} // @name OrganizationCreate

type OrganizationUpdate struct {
	ParentId     *int64         `json:"parentId"`
	SoatoId      *int64         `json:"soatoId"`
	OrgRoles     *pq.Int64Array `json:"orgRoles" gorm:"type:bigint[]"`
	ReviewRoleId *int64         `json:"reviewRoleId"`
	RegionId     *int64         `json:"regionId"`
	DistrictId   *int64         `json:"districtId"`
	QuarterId    *int64         `json:"quarterId"`
} // @name OrganizationUpdate

type QueryParams struct {
	Name       string `json:"name" query:"name"`
	ParentId   int64  `json:"parent_id" query:"parent_id"`
	SoatoId    int64  `json:"soato_id" query:"soato_id"`
	LanguageId int64  `json:"language_id" query:"language_id"`
	RegionId   int64  `json:"region_id" query:"region_id"`
	DistrictId int64  `json:"district_id" query:"district_id"`
	QuarterId  int64  `json:"quarter_id" query:"quarter_id"`
} // @name QueryParams

type OrganizationTranslation struct {
	Id             int64  `json:"id"`
	OrganizationId int64  `json:"organizationId"`
	LanguageId     int64  `json:"languageId"`
	Name           string `json:"name"`
	Description    string `json:"description"`
} // @name OrganizationTranslation

type OrganizationTranslationCreate struct {
	OrganizationId int64  `json:"organizationId"`
	LanguageId     int64  `json:"languageId"`
	Name           string `json:"name"`
	Description    string `json:"description"`
} // @name OrganizationTranslationCreate

type OrganizationTranslationUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
} // @name OrganizationTranslationUpdate
