package role_handler

import (
	"fmt"
	"strings"

	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	role_dto "github.com/Mirsadikovv/idoctor_platform/src/module/role_service/dto"
	role_service "github.com/Mirsadikovv/idoctor_platform/src/module/role_service/service"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/response"
	"github.com/Mirsadikovv/shared/sharedutil"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type roleHandler struct {
	db             *gorm.DB
	log            logger.Logger
	authMiddleware *auth_middleware.AuthMiddleware
	roleService    role_service.RoleService
}

func NewRoleHandler(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	handler := &roleHandler{
		db:             db,
		log:            log,
		authMiddleware: authMiddleware,
		roleService:    role_service.NewRoleService(db),
	}

	roleServiceMiddleware := handler.authMiddleware.BuildMiddleware()
	roleGroup := router.Group("/api/v1/role", roleServiceMiddleware)
	{
		roleGroup.GET("/permissions", handler.Permissions)
		roleGroup.GET("/:id", handler.FindByID)
		roleGroup.GET("/search", handler.Search)
		roleGroup.GET("/page", handler.Page)
		roleGroup.POST("", handler.Create)
		roleGroup.PATCH("/:id", handler.Update)
		roleGroup.PATCH("/:id/restore", handler.Restore)
		roleGroup.DELETE("/:id", handler.Delete)
	}
}

// Permissions 				godoc
// @Summary 				get-all-permissions
// @Description 			get-all-permissions
// @Tags 					role
// @ID 						get-all-permissions
// @Accept 					json
// @Produce 				json
// @Security     			ApiKeyAuth
// @Param       			page query int false "Page number"
// @Param       			perpage query int false "Items per page"
// @Param 	 				search query string false "Search term"
// @Param       			without_paginate query bool false "Without paginate"
// @Success 				200 {array} role_dto.Permission "List of permissions"
// @Failure 				500 {object} map[string]string "Internal server error"
// @Router 					/role/permissions [get]
func (h *roleHandler) Permissions(c echo.Context) error {
	var (
		req             = request.Request(c)
		search          = c.QueryParam("search")
		withoutPaginate = c.QueryParam("without_paginate") == "true"
		router          = c.Echo().Router()
		allPermissions  []role_dto.Permission
	)

	for _, route := range router.Routes() {
		permission := role_dto.Permission{
			Uuid:   sharedutil.Join(route.Method, "_", route.Path),
			Path:   route.Path,
			Method: route.Method,
		}

		if route.Method == "echo_route_not_found" {
			continue
		}

		if search == "" ||
			strings.Contains(strings.ToLower(route.Method), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(route.Path), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(fmt.Sprintf("%s %s", route.Method, route.Path)), strings.ToLower(search)) {
			allPermissions = append(allPermissions, permission)
		}
	}

	if withoutPaginate {
		return req.OK(map[string]any{
			"items":       allPermissions,
			"page":        1,
			"per_page":    len(allPermissions),
			"total_items": int64(len(allPermissions)),
		})
	}

	var (
		paginate = req.NewPaginate()
		start    = (paginate.Page() - 1) * paginate.PerPage()
		end      = start + paginate.PerPage()
	)

	if start > len(allPermissions) {
		return req.OK(map[string]any{
			"items":       []role_dto.Permission{},
			"page":        paginate.Page(),
			"per_page":    paginate.PerPage(),
			"total_items": int64(len(allPermissions)),
		})
	}

	if end > len(allPermissions) {
		end = len(allPermissions)
	}

	permissions := allPermissions[start:end]

	result := map[string]any{
		"items":       permissions,
		"page":        paginate.Page(),
		"per_page":    paginate.PerPage(),
		"total_items": int64(len(allPermissions)),
	}

	return req.OK(result)
}

// Search godoc
// @Summary      GetContent all role with pagination
// @Description  GetContent all role with pagination
// @Tags 		 role
// @ID           search-role
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        params query role_dto.RoleParams false "params"
// @Param        limit  query int    false "Limit the number of results" default(20)
// @Success      200 {object} []role_dto.Role "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /role/search [get]
func (h *roleHandler) Search(c echo.Context) error {

	req := request.RequestWithData[any](c)

	var params role_dto.RoleParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if params.Name != nil {
			name := fmt.Sprintf("%%%s%%", *params.Name)
			tx = tx.Where("roles.name ILIKE ?", name)
		}

		if params.Description != nil {
			description := fmt.Sprintf("%%%s%%", *params.Description)
			tx = tx.Where("roles.description ILIKE ?", description)
		}

		if params.OrganizationId != nil {

			tx = tx.
				Joins("LEFT JOIN organizations ON organizations.id = ?", *params.OrganizationId).
				Where("roles.id = ANY(organizations.org_roles)")
		}

		return tx.Select(
			"roles.id",
			"roles.name",
			"roles.description",
		).Group("roles.id").
			Order("roles.id DESC")

	}

	roles, err := h.roleService.Find(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(roles)
}

// Page 			godoc
// @Summary			role-service-page
// @Description		role-service-page
// @Tags 			role
// @Param       	filter query role_dto.RoleParams false "Filter parameters"
// @Param       	page query int false "Page number"
// @Param       	perpage query int false "Items per page"
// @Security     	ApiKeyAuth
// @Success	200 	{object} role_dto.RolePage "Successful operation"
// @Failure	400 	{object} response.HttpSuccess "Bad request"
// @Failure	500 	{object} response.HttpSuccess "Internal server error"
// @Router			/role/page [get]
func (h *roleHandler) Page(c echo.Context) error {
	req := request.Request(c)

	var params role_dto.RoleParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	tx := func(tx *gorm.DB) *gorm.DB {
		if params.Name != nil {
			tx = tx.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *params.Name))
		}

		if params.Description != nil {
			tx = tx.Where("description ILIKE ?", fmt.Sprintf("%%%s%%", *params.Description))
		}

		if params.OrganizationId != nil {
			tx = tx.Joins("INNER JOIN organizations ON ? = ANY(organizations.org_roles)", *params.OrganizationId)
		}

		return tx.Select(
			"roles.id",
			"roles.name",
			"roles.description",
		).Group("roles.id").
			Order("roles.id DESC")
	}

	items, err := h.roleService.Page(req.Context(), req.NewPaginate(), tx)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(items)
}

// GetById godoc
// @Summary      GetContent role by ID
// @Description  GetContent role by ID
// @Tags 		 role
// @ID           get-role-by-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "role ID"
// @Success      200 {object} role_dto.Role "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /role/{id} [get]
func (h *roleHandler) FindByID(c echo.Context) error {
	req := request.RequestWithData[any](c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	role, err := h.roleService.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(role)
}

// Create godoc
// @Summary      Create role
// @Description  Create a new role
// @Tags 		 role
// @ID           create-role
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        input body role_dto.RoleCreate true "role information"
// @Success      201 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /role [POST]
func (h *roleHandler) Create(c echo.Context) error {

	req := request.Request(c)

	var roleDto role_dto.RoleCreate
	{
		if err := req.BindBody(&roleDto); err != nil {
			return req.BadRequest(err)
		}
	}

	id, err := h.roleService.Create(c.Request().Context(), &roleDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(response.NewID(id))
}

func (h *roleHandler) Update(c echo.Context) error {
	return nil
}

func (h *roleHandler) Restore(c echo.Context) error {
	return nil
}

func (h *roleHandler) Delete(c echo.Context) error {
	return nil
}
