package organization_handler

import (
	"fmt"

	auth_dto "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/dto"
	organization_dto "github.com/Mirsadikovv/idoctor_platform/src/module/organization_service/dto"

	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// AddInvitation godoc
// @Summary      invitation
// @Description  invitation
// @Tags 		 invitation
// @ID           create-invitation
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        input body organization_dto.InvitationCreate true "invitation information"
// @Success      201 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /invitation/create [post]
func (o *organizationHandler) AddInvitation(ctx echo.Context) error {

	req := request.Request(ctx)

	var invitationDto organization_dto.InvitationCreate
	{
		if err := req.BindBody(&invitationDto); err != nil {
			return req.BadRequest(err)
		}
	}

	if len(fmt.Sprint(invitationDto.Pin)) != 14 {
		return req.BadRequest("Pin must be 14 characters long")
	}

	if invitationDto.OrganizationId == 0 {
		return req.BadRequest("Organization ID is required")
	}

	if invitationDto.RoleId == 0 {
		return req.BadRequest("Role ID is required")
	}

	id, err := o.organizationService.AddInvitation(ctx, &invitationDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.Created(response.ID64{ID: id})
}

// Page 		 godoc
// @Summary      GetContent all invitation with pagination
// @Description  GetContent all invitation with pagination
// @Tags 		 invitation
// @ID           get-all-invitation
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        page query string false "Page number" default(1)
// @Param        perpage query string false "Number of items per page" default(10)
// @Param        queryParams query organization_dto.InvitationParams false "queryParams"
// @Success      200 {object} organization_dto.InvitationPage "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /invitation/page [get]
func (o *organizationHandler) PageInvitation(ctx echo.Context) error {

	req := request.Request(ctx)

	var queryParams organization_dto.InvitationParams
	{
		if err := req.BindQuery(&queryParams); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if queryParams.OrganizationId != nil {
			tx = tx.Where("invitations.organization_id = ?", *queryParams.OrganizationId)
		}

		if queryParams.RoleId != nil {
			tx = tx.Where("invitations.role_id = ?", *queryParams.RoleId)
		}

		if queryParams.Pin != nil {
			pin := fmt.Sprintf("%s%%", *queryParams.Pin)
			tx = tx.Where("invitations.pin::text ILIKE ?", pin)
		}

		if queryParams.Blocked.IsValid() {
			if queryParams.Blocked.Bool() {
				tx = tx.Unscoped().Where("blocked_at IS NOT NULL")
			} else if !queryParams.Blocked.Bool() {
				tx = tx.Unscoped().Where("blocked_at IS NULL")
			}
		} else {
			tx = tx.Unscoped()
		}

		orgExp := gorm.Expr(`
			SELECT 
				* 
			FROM 
				organization_translations
			WHERE
				organization_id = organizations.id
			ORDER BY 
				(language_id = ?) DESC
			LIMIT 1`, queryParams.LanguageId)

		return tx.
			Joins("LEFT JOIN organizations ON organizations.id = invitations.organization_id").
			Joins("LEFT JOIN LATERAL (?) AS ot ON true", orgExp).
			Joins("LEFT JOIN roles ON roles.id = invitations.role_id").
			Joins("LEFT JOIN users ON users.pin = invitations.pin").
			Select("invitations.*",
				`JSONB_BUILD_OBJECT(
						'id', 		   	  organizations.id,
						'parent_id',   	  organizations.parent_id,
						'name', 		  ot.name,
						'description',	  ot.description,
						'soato_id',    	  organizations.soato_id,
						'org_roles',   	  organizations.org_roles,
						'review_role_id', organizations.review_role_id
			) AS organization`,

				`JSONB_BUILD_OBJECT(
						'id', 		   roles.id,
						'name', 	   roles.name,
						'description', roles.description
			) AS role`,
				`JSONB_BUILD_OBJECT(
					'id', 		   users.id,
					'firstName',   users.first_name,
					'lastName',    users.last_name,
					'middleName',  users.middle_name,
					'gender',      users.gender,
					'pin',         users.pin
				) AS user`,
			).Order("invitations.id DESC")
	}

	userPage, err := o.organizationService.PageInvitation(req.Context(), req.NewPaginate(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(userPage)
}

// Search godoc
// @Summary      GetContent all invitation with pagination
// @Description  GetContent all invitation with pagination
// @Tags 		 invitation
// @ID           search-invitation
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        limit  query int    false "Limit the number of results" default(20)
// @Param        params query organization_dto.InvitationParams false "params"
// @Success      200 {object} []organization_dto.InvitationDto "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /invitation/search [get]
func (o *organizationHandler) SearchInvitation(ctx echo.Context) error {

	req := request.Request(ctx)

	var queryParams organization_dto.InvitationParams
	{
		if err := req.BindQuery(&queryParams); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if queryParams.OrganizationId != nil {
			tx = tx.Where("invitations.organization_id = ?", *queryParams.OrganizationId)
		}

		if queryParams.RoleId != nil {
			tx = tx.Where("invitations.role_id = ?", *queryParams.RoleId)
		}

		if queryParams.Pin != nil {
			pin := fmt.Sprintf("%s%%", *queryParams.Pin)
			tx = tx.Where("invitations.pin::TEXT ILIKE ?", pin)
		}

		if queryParams.UserId != nil {
			tx = tx.Where("users.id = ?", *queryParams.UserId)
		}

		if queryParams.Blocked.IsValid() {
			if queryParams.Blocked.Bool() {
				tx = tx.Unscoped().Where("blocked_at IS NOT NULL")
			} else if !queryParams.Blocked.Bool() {
				tx = tx.Unscoped().Where("blocked_at IS NULL")
			}
		} else {
			tx = tx.Unscoped()
		}

		orgExp := gorm.Expr(`
			SELECT 
				* 
			FROM 
				organization_translations
			WHERE
				organization_id = organizations.id
			ORDER BY 
				(language_id = ?) DESC
			LIMIT 1`, queryParams.LanguageId)

		return tx.
			Joins("LEFT JOIN organizations ON organizations.id = invitations.organization_id").
			Joins("LEFT JOIN LATERAL (?) AS ot ON true", orgExp).
			Joins("LEFT JOIN roles ON roles.id = invitations.role_id").
			Joins("LEFT JOIN users ON users.pin = invitations.pin").
			Select("invitations.*",
				`JSONB_BUILD_OBJECT(
						'id', 		   	  organizations.id,
						'parent_id',   	  organizations.parent_id,
						'name', 		  ot.name,
						'description',	  ot.description,
						'soato_id',    	  organizations.soato_id,
						'org_roles',   	  organizations.org_roles,
						'review_role_id', organizations.review_role_id
			) AS organization`,

				`JSONB_BUILD_OBJECT(
						'id', 		   roles.id,
						'name', 	   roles.name,
						'description', roles.description
			) AS role`,
				`JSONB_BUILD_OBJECT(
					'id', 		   users.id,
					'firstName',   users.first_name,
					'lastName',    users.last_name,
					'middleName',  users.middle_name,
					'gender',      users.gender,
					'pin',         users.pin
				) AS user`,
			)
	}

	users, err := o.organizationService.FindInvitation(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(users)
}

// MyInvitations godoc
// @Summary      Get my invitations
// @Description  Get all invitations belonging to the authenticated user
// @Tags 		 invitation
// @ID           get-my-invitations
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        language_id query string false "Language ID"
// @Param        limit  query int    false "Limit the number of results" default(20)
// @Success      200 {object} []organization_dto.InvitationDto "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /invitation/my [get]
func (o *organizationHandler) MyInvitations(ctx echo.Context) error {

	req := request.RequestWithData[auth_dto.AuthUser](ctx)
	user := req.AuthUser()

	languageId, err := req.QueryToInt("language_id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		tx = tx.Where("users.id = ?", user.Id)

		orgExp := gorm.Expr(`
			SELECT 
				* 
			FROM 
				organization_translations
			WHERE
				organization_id = organizations.id
			ORDER BY 
				(language_id = ?) DESC
			LIMIT 1`, languageId)

		return tx.
			Joins("LEFT JOIN organizations ON organizations.id = invitations.organization_id").
			Joins("LEFT JOIN LATERAL (?) AS ot ON true", orgExp).
			Joins("LEFT JOIN roles ON roles.id = invitations.role_id").
			Joins("LEFT JOIN users ON users.pin = invitations.pin").
			Select("invitations.*",
				`JSONB_BUILD_OBJECT(
						'id', 		   	  organizations.id,
						'name', 		  ot.name
			) AS organization`,
				`JSONB_BUILD_OBJECT(
						'id', 		   roles.id,
						'name', 	   roles.name
			) AS role`,
				`JSONB_BUILD_OBJECT(
					'id', 		   users.id,
					'firstName',   users.first_name,
					'lastName',    users.last_name,
					'middleName',  users.middle_name,
					'gender',      users.gender
				) AS user`,
			)
	}

	invitations, err := o.organizationService.MyInvitations(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(invitations)
}

// GetById godoc
// @Summary      GetContent invitation by ID
// @Description  GetContent invitation by ID
// @Tags 		 invitation
// @ID           get-invitation-by-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "invitation ID"
// @Param        language_id query string false "Language ID"
// @Success      200 {object} organization_dto.InvitationDto "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /invitation/{id} [get]
func (o *organizationHandler) GetByIdInvitation(ctx echo.Context) error {

	req := request.Request(ctx)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	languageId, err := req.QueryToInt("language_id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		tx = tx.Unscoped()

		orgExp := gorm.Expr(`
			SELECT 
				* 
			FROM 
				organization_translations
			WHERE
				organization_id = organizations.id
			ORDER BY 
				(language_id = ?) DESC
			LIMIT 1`, languageId)

		return tx.
			Joins("LEFT JOIN organizations ON organizations.id = invitations.organization_id").
			Joins("LEFT JOIN LATERAL (?) AS ot ON true", orgExp).
			Joins("LEFT JOIN roles ON roles.id = invitations.role_id").
			Joins("LEFT JOIN users ON users.pin = invitations.pin").
			Select("invitations.*",
				`JSONB_BUILD_OBJECT(
						'id', 		   	  organizations.id,
						'parent_id',   	  organizations.parent_id,
						'name', 		  ot.name,
						'description',	  ot.description,
						'soato_id',    	  organizations.soato_id,
						'org_roles',   	  organizations.org_roles,
						'review_role_id', organizations.review_role_id
			) AS organization`,

				`JSONB_BUILD_OBJECT(
						'id', 		   roles.id,
						'name', 	   roles.name,
						'description', roles.description
			) AS role`,
				`JSONB_BUILD_OBJECT(
					'id', 		   users.id,
					'firstName',   users.first_name,
					'lastName',    users.last_name,
					'middleName',  users.middle_name,
					'gender',      users.gender,
					'pin',         users.pin
				) AS user`).
			Where("invitations.id = ?", id)
	}

	user, err := o.organizationService.FindOneInvitation(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(user)
}

// Restore       godoc
// @Summary      invitation
// @Description  invitation
// @Tags 		 invitation
// @ID           restore-invitation
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "invitation ID"
// @Success      204 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /invitation/delete_or_restore/{id} [patch]
func (o *organizationHandler) DeleteOrRestoreInvitation(ctx echo.Context) error {

	req := request.Request(ctx)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	if err := o.organizationService.DeleteOrRestoreInvitation(ctx, filter); err != nil {
		return req.BadRequest(err)
	}

	return req.NoContent()
}
