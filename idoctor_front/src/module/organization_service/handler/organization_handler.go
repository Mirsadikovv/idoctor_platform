package organization_handler

import (
	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	organization_dto "github.com/Mirsadikovv/idoctor_platform/src/module/organization_service/dto"
	organization_service "github.com/Mirsadikovv/idoctor_platform/src/module/organization_service/service"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrganizationHandler interface {
}

type organizationHandler struct {
	db                  *gorm.DB
	log                 logger.Logger
	organizationService organization_service.OrganizationService
	authMiddleware      *auth_middleware.AuthMiddleware
}

func NewOrganizationHandler(group *echo.Group, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	handler := &organizationHandler{
		db:                  db,
		log:                 log,
		authMiddleware:      authMiddleware,
		organizationService: organization_service.NewOrganizationService(db),
	}

	organizationAuthMiddleware := authMiddleware.BuildMiddleware()
	organizationGroup := group.Group("/organization", organizationAuthMiddleware)
	{
		organizationGroup.POST("/create", handler.Create)
		organizationGroup.PATCH("/:id", handler.Update)
		organizationGroup.GET("/page", handler.Page)
		organizationGroup.GET("/search", handler.Search)
		organizationGroup.GET("/:id", handler.GetById)
	}

	organizationTranslationGroup := organizationGroup.Group("/:id/translation", organizationAuthMiddleware)
	{
		organizationTranslationGroup.POST("/create", handler.CreateTranslation)
		organizationTranslationGroup.GET("/list", handler.ListTranslation)
		organizationTranslationGroup.GET("/:translation_id", handler.GetTranslationById)
		organizationTranslationGroup.PATCH("/:translation_id", handler.UpdateTranslation)
	}

	invitationGroup := group.Group("/invitation")
	{
		invitationGroup.POST("/create", handler.AddInvitation)
		invitationGroup.PATCH("/delete_or_restore/:id", handler.DeleteOrRestoreInvitation)
		invitationGroup.GET("/search", handler.SearchInvitation)
		invitationGroup.GET("/page", handler.PageInvitation)
		invitationGroup.GET("/:id", handler.GetByIdInvitation)
		invitationGroup.GET("/my", handler.MyInvitations)
	}
}

// Create godoc
// @Summary      organization
// @Description  organization
// @Tags 		 organization
// @ID           create-organization
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        input body organization_dto.OrganizationCreate true "organization information"
// @Success      201 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /organization/create [post]
func (o *organizationHandler) Create(ctx echo.Context) error {

	req := request.Request(ctx)

	var organizationDto organization_dto.OrganizationCreate
	{
		if err := req.BindBody(&organizationDto); err != nil {
			return req.BadRequest(err)
		}
	}

	id, err := o.organizationService.Create(ctx, &organizationDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.Created(response.NewID(id))
}

// Update 		 godoc
// @Summary      organization
// @Description  organization
// @Tags 		 organization
// @ID           update-organization
// @Accept       json
// @Produce      json
// @Param        id path string true "organization ID"
// @Security     ApiKeyAuth
// @Param        input body organization_dto.OrganizationUpdate true "organization information"
// @Success      204 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /organization/{id} [patch]
func (o *organizationHandler) Update(ctx echo.Context) error {

	req := request.Request(ctx)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	var (
		organizationDto organization_dto.OrganizationUpdate
	)
	{
		if err := req.BindBody(&organizationDto); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	if err := o.organizationService.Update(req.Context(), &organizationDto, filter); err != nil {
		return req.BadRequest(err)
	}

	return req.NoContent()
}

// Page 		 godoc
// @Summary      GetContent all organization with pagination
// @Description  GetContent all organization with pagination
// @Tags 		 organization
// @ID           get-all-organization
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        page query string false "Page number" default(1)
// @Param        perpage query string false "Number of items per page" default(10)
// @Param        queryParams query organization_dto.QueryParams false "organization information"
// @Success      200 {object} []organization_dto.Organization "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /organization/page [get]
func (o *organizationHandler) Page(ctx echo.Context) error {

	req := request.Request(ctx)

	var queryParams organization_dto.QueryParams
	{
		if err := req.BindQuery(&queryParams); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if queryParams.Name != "" {
			tx = tx.Where("ot.name ILIKE ?", "%"+queryParams.Name+"%")
		}

		if queryParams.ParentId != 0 {
			tx = tx.Where("organizations.parent_id = ?", queryParams.ParentId)
		}

		if queryParams.RegionId != 0 {
			tx = tx.Where("organizations.region_id = ?", queryParams.RegionId)
		}

		if queryParams.DistrictId != 0 {
			tx = tx.Where("organizations.district_id = ?", queryParams.DistrictId)
		}

		if queryParams.QuarterId != 0 {
			tx = tx.Where("organizations.quarter_id = ?", queryParams.QuarterId)
		}

		if queryParams.SoatoId != 0 {
			tx = tx.Where("organizations.soato_id = ?", queryParams.SoatoId)
		}

		exp := gorm.Expr(`
			SELECT 
				* 
			FROM 
				organization_translations
			WHERE 
				organization_id = organizations.id
			ORDER BY
				(language_id = ?) DESC
			LIMIT 1`,
			queryParams.LanguageId,
		)

		return tx.
			Select(
				"organizations.*",
				"ot.name AS name",
				"ot.description AS description",
				"ot.language_id AS language_id",
			).
			Joins("LEFT JOIN LATERAL (?) AS ot ON true", exp)
	}

	userPage, err := o.organizationService.Page(req.Context(), req.NewPaginate(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(userPage)
}

// Search godoc
// @Summary      GetContent all organization with pagination
// @Description  GetContent all organization with pagination
// @Tags 		 organization
// @ID           search-organization
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        limit  query int    false "Limit the number of results" default(20)
// @Param        params query organization_dto.QueryParams false "organization information"
// @Success      200 {object} []organization_dto.Organization "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /organization/search [get]
func (o *organizationHandler) Search(ctx echo.Context) error {

	req := request.Request(ctx)

	var queryParams organization_dto.QueryParams
	{
		if err := req.BindQuery(&queryParams); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if queryParams.Name != "" {
			tx = tx.Where("ot.name ILIKE ?", "%"+queryParams.Name+"%")
		}

		if queryParams.ParentId != 0 {
			tx = tx.Where("organizations.parent_id = ?", queryParams.ParentId)
		}

		if queryParams.RegionId != 0 {
			tx = tx.Where("organizations.region_id = ?", queryParams.RegionId)
		}

		if queryParams.DistrictId != 0 {
			tx = tx.Where("organizations.district_id = ?", queryParams.DistrictId)
		}

		if queryParams.QuarterId != 0 {
			tx = tx.Where("organizations.quarter_id = ?", queryParams.QuarterId)
		}

		if queryParams.SoatoId != 0 {
			tx = tx.Where("organizations.soato_id = ?", queryParams.SoatoId)
		}

		exp := gorm.Expr(`
			SELECT 
				* 
			FROM 
				organization_translations
			WHERE 
				organization_id = organizations.id
			ORDER BY
				(language_id = ?) DESC
			LIMIT 1`,
			queryParams.LanguageId,
		)

		return tx.
			Select(
				"organizations.*",
				"ot.name AS name",
				"ot.description AS description",
				"ot.language_id AS language_id",
			).
			Joins("LEFT JOIN LATERAL (?) AS ot ON true", exp)
	}

	users, err := o.organizationService.Find(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(users)
}

// GetById godoc
// @Summary      GetContent organization by ID
// @Description  GetContent organization by ID
// @Tags 		 organization
// @ID           get-organization-by-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "organization ID"
// @param        language_id query int false "language_id"
// @Success      200 {object} organization_dto.OrganizationDto "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /organization/{id} [get]
func (o *organizationHandler) GetById(ctx echo.Context) error {

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
		tx = tx.Where("organizations.id = ?", id)

		ot_exp := gorm.Expr(`
			SELECT 
				* 
			FROM 
				organization_translations
			WHERE 
				organization_id = organizations.id
			ORDER BY
				(language_id = ?) DESC
			LIMIT 1`,
			languageId,
		)

		pot_exp := gorm.Expr(`
			SELECT 
				* 
			FROM 
				organization_translations
			WHERE 
				organization_id = parent_organizations.id
			ORDER BY
				(language_id = ?) DESC
			LIMIT 1`,
			languageId,
		)

		rolesExp := gorm.Expr(`
			SELECT 
				COALESCE(
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', r.id,
							'name', r.name,
							'description', r.description
						)
					) FILTER (WHERE r.id IS NOT NULL), 
					'[]'::json
				) AS roles
			FROM 
				unnest(organizations.org_roles) AS role_id
			LEFT JOIN 
				roles r ON r.id = role_id
		`)

		regionExp := gorm.Expr(`
		SELECT 
			* 
		FROM 
			addresses
		WHERE 
			soato_id = organizations.region_id
		AND
			lang_id = ?`, languageId,
		)

		// Район
		districtExp := gorm.Expr(`
		SELECT 
			* 
		FROM 
			addresses
		WHERE 
			soato_id = organizations.district_id
		AND
			lang_id = ?`, languageId,
		)

		// Квартал
		quarterExp := gorm.Expr(`
		SELECT 
			* 
		FROM 
			addresses
		WHERE 
			soato_id = organizations.quarter_id 
		AND
			lang_id = ?
			`, languageId,
		)

		// SOATO
		soatoExp := gorm.Expr(`
		SELECT 
			* 
		FROM 
			addresses
		WHERE 
			soato_id = organizations.soato_id 
		AND
			lang_id = ?`, languageId,
		)

		return tx.
			Select(
				"organizations.id",
				"organizations.parent_id",
				"organizations.soato_id",
				"organizations.region_id",
				"organizations.district_id",
				"organizations.quarter_id",
				"organizations.review_role_id",

				"ot.name AS name",
				"ot.description AS description",
				"ot.language_id AS language_id",
				"roles_data.roles AS org_roles",
				"region_t.title AS region_name",
				"district_t.title AS district_name",
				"quarter_t.title AS quarter_name",
				"soato_t.title AS soato_name",

				`JSONB_BUILD_OBJECT(
					'id', 			  parent_organizations.id,
					'parent_id',      parent_organizations.parent_id,
					'soato_id', 	  parent_organizations.soato_id,
					'region_id', 	  parent_organizations.region_id,
					'district_id', 	  parent_organizations.district_id,
					'quarter_id', 	  parent_organizations.quarter_id,
					'review_role_id', parent_organizations.review_role_id,
					'name', 		  coalesce(pot.name, ''),
					'description', 	  coalesce(pot.description, ''),
					'language_id', 	  coalesce(pot.language_id, 0)
				) AS parent_organization`,

				`JSONB_BUILD_OBJECT(
					'id', 		   review_roles.id,
					'name', 	   review_roles.name,
					'description', review_roles.description
				) AS review_role`,
			).
			Joins("LEFT JOIN organizations AS parent_organizations ON parent_organizations.id = organizations.parent_id").
			Joins("LEFT JOIN roles AS review_roles ON review_roles.id = organizations.review_role_id").
			Joins("LEFT JOIN LATERAL (?) AS ot ON true", ot_exp).
			Joins("LEFT JOIN LATERAL (?) AS pot ON true", pot_exp).
			Joins("LEFT JOIN LATERAL (?) AS roles_data ON true", rolesExp).
			Joins("LEFT JOIN LATERAL (?) AS region_t ON true", regionExp).
			Joins("LEFT JOIN LATERAL (?) AS district_t ON true", districtExp).
			Joins("LEFT JOIN LATERAL (?) AS quarter_t ON true", quarterExp).
			Joins("LEFT JOIN LATERAL (?) AS soato_t ON true", soatoExp)
	}

	user, err := o.organizationService.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(user)
}

// CreateTranslation godoc
// @Summary      Create organization translation
// @Description  Create organization translation
// @Tags 		 organization-translation
// @ID           create-organization-translation
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path int true "organization ID"
// @Param        input body organization_dto.OrganizationTranslationCreate true "organization information"
// @Success      201 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /organization/{id}/translation/create [post]
func (h *organizationHandler) CreateTranslation(ctx echo.Context) error {

	req := request.Request(ctx)

	organizationId, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	var translation organization_dto.OrganizationTranslationCreate
	{
		if err := req.Bind(&translation); err != nil {
			return req.BadRequest(err)
		}

		translation.OrganizationId = organizationId
	}

	id, err := h.organizationService.CreateTranslation(req.Context(), translation)
	if err != nil {
		return req.BadRequest(err)
	}

	return req.Created(response.ID64{ID: id})
}

// GetTranslations godoc
// @Summary      GetContent organization translations
// @Description  GetContent organization translations
// @Tags 		 organization-translation
// @ID           get-organization-translations
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "organization ID"
// @Success      200 {object} []organization_dto.OrganizationTranslation "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /organization/{id}/translation/list [get]
func (h *organizationHandler) ListTranslation(ctx echo.Context) error {

	req := request.Request(ctx)

	organizationId, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("organization_translations.organization_id = ?", organizationId)
	}

	tranlations, err := h.organizationService.ListTranslation(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}

		if tranlations == nil {
			return req.OK([]organization_dto.OrganizationTranslation{})
		}
	}

	return req.OK(tranlations)
}

// GetTranslationById godoc
// @Summary      GetContent organization translation by ID
// @Description  GetContent organization translation by ID
// @Tags 		 organization-translation
// @ID           get-organization-translation-by-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "organization ID"
// @Param        translation_id path string true "organization translation ID"
// @Success      200 {object} organization_dto.OrganizationTranslation "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /organization/{id}/translation/{translation_id} [get]
func (h *organizationHandler) GetTranslationById(ctx echo.Context) error {

	req := request.Request(ctx)

	organizationId, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	tranlationId, err := req.ParamToInt("translation_id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("organization_translations.organization_id = ? AND organization_translations.id = ?", organizationId, tranlationId)
	}

	tranlation, err := h.organizationService.GetTranslationById(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}

		if tranlation == nil {
			return req.OK(organization_dto.OrganizationTranslation{})
		}
	}

	return req.OK(tranlation)
}

// UpdateTranslation godoc
// @Summary      Update organization translation
// @Description  Update organization translation
// @Tags 		 organization-translation
// @ID           update-organization-translation
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "organization ID"
// @Param        translation_id path string true "organization translation ID"
// @Param        input body organization_dto.OrganizationTranslationUpdate true "organization information"
// @Success      204 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /organization/{id}/translation/{translation_id} [patch]
func (h *organizationHandler) UpdateTranslation(ctx echo.Context) error {

	req := request.Request(ctx)

	organizationId, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	translationId, err := req.ParamToInt("translation_id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	var (
		translation organization_dto.OrganizationTranslationUpdate
	)
	{
		if err := req.BindBody(&translation); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("organization_translations.organization_id = ? AND organization_translations.id = ?", organizationId, translationId)
	}

	if err := h.organizationService.UpdateTranslation(req.Context(), translation, filter); err != nil {
		return req.BadRequest(err)
	}

	return req.NoContent()
}
