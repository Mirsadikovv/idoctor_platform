package organization_service

import (
	"context"

	organization_dto "git.sriss.uz/mehnat/inspector_platform/src/module/organization_service/dto"
	organization_model "git.sriss.uz/mehnat/inspector_platform/src/module/organization_service/model"
	"github.com/labstack/echo/v4"

	"git.sriss.uz/shared/shared_service/pg"
	"git.sriss.uz/shared/shared_service/request"
	"gorm.io/gorm"
)

func (s *organizationService) AddInvitation(ctx echo.Context, invitationDto *organization_dto.InvitationCreate) (int64, error) {

	invitationModel := &organization_model.Invitation{
		OrganizationId: invitationDto.OrganizationId,
		Pin:            invitationDto.Pin,
		RoleId:         invitationDto.RoleId,
	}

	if err := pg.Create(s.db.WithContext(ctx.Request().Context()), invitationModel, "id"); err != nil {
		return 0, err
	}

	return invitationModel.Id, nil
}

func (s *organizationService) PageInvitation(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*organization_dto.InvitationPage, error) {
	return pg.PageWithScan[organization_model.Invitation, organization_dto.InvitationDto](s.db, paginate, filter)
}

func (s *organizationService) FindInvitation(ctx context.Context, filter pg.Filter) ([]organization_dto.InvitationDto, error) {
	users, err := pg.FindWithScan[organization_model.Invitation, organization_dto.InvitationDto](s.db, filter)
	{
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

func (s *organizationService) FindOneInvitation(ctx context.Context, filter pg.Filter) (*organization_dto.InvitationDto, error) {
	return pg.FindOneWithScan[organization_model.Invitation, organization_dto.InvitationDto](s.db, filter)
}

func (s *organizationService) MyInvitations(ctx context.Context, filter pg.Filter) ([]organization_dto.InvitationDto, error) {

	invitations, err := pg.FindWithScan[organization_model.Invitation, organization_dto.InvitationDto](s.db.WithContext(ctx), filter)
	{
		if err != nil {
			return nil, err
		}
	}

	return invitations, nil
}

func (s *organizationService) DeleteOrRestoreInvitation(ctx echo.Context, filter pg.Filter) error {

	restoreData := map[string]any{
		"blocked_at": gorm.Expr("CASE WHEN blocked_at IS NULL THEN now() ELSE null END"),
	}

	if _, err := pg.Update[organization_model.Invitation](s.db.WithContext(ctx.Request().Context()).Unscoped(), restoreData, filter); err != nil {
		return err
	}

	return nil
}
