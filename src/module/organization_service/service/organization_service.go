package organization_service

import (
	"context"

	organization_dto "github.com/Mirsadikovv/idoctor_platform/src/module/organization_service/dto"
	organization_model "github.com/Mirsadikovv/idoctor_platform/src/module/organization_service/model"
	"github.com/labstack/echo/v4"

	"github.com/Mirsadikovv/shared/pg"
	"github.com/Mirsadikovv/shared/request"
	"gorm.io/gorm"
)

type OrganizationService interface {
	Create(ctx echo.Context, in *organization_dto.OrganizationCreate) (int64, error)
	Update(ctx context.Context, in *organization_dto.OrganizationUpdate, filter pg.Filter) (err error)
	Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*organization_dto.OrganizationPage, error)
	Find(ctx context.Context, filter pg.Filter) ([]organization_dto.Organization, error)
	FindOne(ctx context.Context, filter pg.Filter) (*organization_dto.OrganizationDto, error)

	CreateTranslation(ctx context.Context, in organization_dto.OrganizationTranslationCreate) (int64, error)
	ListTranslation(ctx context.Context, filter pg.Filter) ([]organization_dto.OrganizationTranslation, error)
	GetTranslationById(ctx context.Context, filter pg.Filter) (*organization_dto.OrganizationTranslation, error)
	UpdateTranslation(ctx context.Context, in organization_dto.OrganizationTranslationUpdate, filter pg.Filter) error

	AddInvitation(ctx echo.Context, invitationDto *organization_dto.InvitationCreate) (int64, error)
	PageInvitation(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*organization_dto.InvitationPage, error)
	FindInvitation(ctx context.Context, filter pg.Filter) ([]organization_dto.InvitationDto, error)
	FindOneInvitation(ctx context.Context, filter pg.Filter) (*organization_dto.InvitationDto, error)
	MyInvitations(ctx context.Context, filter pg.Filter) ([]organization_dto.InvitationDto, error)
	DeleteOrRestoreInvitation(ctx echo.Context, filter pg.Filter) error
}

type organizationService struct {
	db *gorm.DB
}

func NewOrganizationService(db *gorm.DB) OrganizationService {
	return &organizationService{
		db: db,
	}
}

func (s *organizationService) Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*organization_dto.OrganizationPage, error) {
	return pg.PageWithScan[organization_model.Organization, organization_dto.Organization](s.db, paginate, filter)
}

func (s *organizationService) Find(ctx context.Context, filter pg.Filter) ([]organization_dto.Organization, error) {
	users, err := pg.FindWithScan[organization_model.Organization, organization_dto.Organization](s.db, filter)
	{
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

func (s *organizationService) FindOne(ctx context.Context, filter pg.Filter) (*organization_dto.OrganizationDto, error) {
	return pg.FindOneWithScan[organization_model.Organization, organization_dto.OrganizationDto](s.db, filter)
}

func (s *organizationService) Create(ctx echo.Context, in *organization_dto.OrganizationCreate) (int64, error) {

	organizationModel := &organization_model.Organization{
		ParentId:     in.ParentId,
		SoatoId:      in.SoatoId,
		OrgRoles:     in.OrgRoles,
		ReviewRoleId: in.ReviewRoleId,
		RegionId:     in.RegionId,
		DistrictId:   in.DistrictId,
		QuarterId:    in.QuarterId,
	}

	if err := pg.Transaction(
		s.db,

		func(tx *gorm.DB) error {

			if err := pg.Create(tx.WithContext(ctx.Request().Context()), organizationModel, "id"); err != nil {
				return err
			}

			var translation organization_model.OrganizationTranslation
			{
				translation.OrganizationId = organizationModel.Id
				translation.LanguageId = in.LanguageId
				translation.Name = in.Name
				translation.Description = in.Description
			}

			if err := pg.Create(tx.WithContext(ctx.Request().Context()), &translation); err != nil {
				return err
			}

			return nil
		},
	); err != nil {
		return 0, err
	}

	return organizationModel.Id, nil
}

func (s *organizationService) Update(ctx context.Context, in *organization_dto.OrganizationUpdate, filter pg.Filter) (err error) {

	_, err = pg.Update[organization_model.Organization](s.db.WithContext(ctx), &in, filter)

	return
}

func (s *organizationService) CreateTranslation(ctx context.Context, in organization_dto.OrganizationTranslationCreate) (int64, error) {

	var translation organization_model.OrganizationTranslation
	{
		translation.OrganizationId = in.OrganizationId
		translation.LanguageId = in.LanguageId
		translation.Name = in.Name
		translation.Description = in.Description
	}

	err := pg.Create(s.db.WithContext(ctx), &translation)
	{
		if err != nil {
			return 0, err
		}
	}

	return translation.Id, nil
}

func (s *organizationService) ListTranslation(ctx context.Context, filter pg.Filter) ([]organization_dto.OrganizationTranslation, error) {
	return pg.FindWithScan[organization_model.OrganizationTranslation, organization_dto.OrganizationTranslation](s.db.WithContext(ctx), filter)
}

func (s *organizationService) GetTranslationById(ctx context.Context, filter pg.Filter) (*organization_dto.OrganizationTranslation, error) {
	return pg.FindOneWithScan[organization_model.OrganizationTranslation, organization_dto.OrganizationTranslation](s.db.WithContext(ctx), filter)
}

func (s *organizationService) UpdateTranslation(ctx context.Context, in organization_dto.OrganizationTranslationUpdate, filter pg.Filter) (err error) {
	_, err = pg.Update[organization_model.OrganizationTranslation](s.db.WithContext(ctx), &in, filter)
	return
}
