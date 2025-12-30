package user_handler

import (
	"fmt"

	auth_dto "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/dto"
	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	user_dto "github.com/Mirsadikovv/idoctor_platform/src/module/user_service/dto"
	user_service "github.com/Mirsadikovv/idoctor_platform/src/module/user_service/service"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler interface {
}

type userHandler struct {
	db             *gorm.DB
	log            logger.Logger
	userService    user_service.UserService
	authMiddleware *auth_middleware.AuthMiddleware
}

func NewUserHandler(group *echo.Group, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	handler := &userHandler{
		db:             db,
		log:            log,
		authMiddleware: authMiddleware,
		userService:    user_service.NewUserService(db),
	}

	// userGroup := group.Group("/user", authMiddleware.BuildMiddleware())
	userGroup := group.Group("/user")

	{
		userGroup.POST("", handler.Create)
		userGroup.GET("/page", handler.Page)
		userGroup.GET("/search", handler.Search)
		userGroup.GET("/:id", handler.GetByID)
		userGroup.PATCH("/:id", handler.Update)
		userGroup.DELETE("/:id", handler.Delete)
		userGroup.PATCH("/:id/restore", handler.Restore)
	}
}

// Create godoc
// @Summary      Create user
// @Description  Create a new user
// @Tags 		 user
// @ID           create-user
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        user body user_dto.UserCreate true "User information"
// @Success      201 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /user [post]
func (h *userHandler) Create(ctx echo.Context) error {

	req := request.Request(ctx)

	var userCreate user_dto.UserCreate
	{
		if err := req.BindBody(&userCreate); err != nil {
			return req.BadRequest(err)
		}
	}

	userId, err := h.userService.Create(&userCreate)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.Created(userId)
}

// Delete godoc
// @Summary      Delete user
// @Description  Delete user by ID
// @Tags 		 user
// @ID           delete-user
// @Accept       json
// @Security     ApiKeyAuth
// @Param        id path string true "user ID"
// @Success      204 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /user/{id} [delete]
func (h *userHandler) Delete(ctx echo.Context) error {

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

	if err := h.userService.Delete(ctx, filter); err != nil {
		return req.BadRequest(err)
	}

	return req.NoContent()
}

// Page godoc
// @Summary      GetContent all user with pagination
// @Description  GetContent all user with pagination
// @Tags 		 user
// @ID           get-all-user
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        page query string false "Page number" default(1)
// @Param        perpage query string false "Number of items per page" default(10)
// @Param        params query user_dto.UserParams false "user information"
// @Success      200 {object} []user_dto.User "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /user/page [get]
func (h *userHandler) Page(ctx echo.Context) error {

	req := request.RequestWithData[auth_dto.AuthUser](ctx)

	var params user_dto.UserParams
	{
		if err := req.Bind(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if params.Fullname != nil {
			fullname := fmt.Sprintf("%%%s%%", *params.Fullname)
			tx = tx.Where("(users.first_name ILIKE ? OR users.middle_name ILIKE ? OR users.last_name ILIKE ?)", fullname, fullname, fullname)
		}

		if params.Gender.Valid() {
			tx = tx.Where("users.gender = ?", params.Gender.String())
		}

		allowedSortFields := map[string]string{
			"id":        "users.id",
			"firstName": "users.first_name",
			"lastName":  "users.last_name",
			"createdAt": "users.created_at",
		}

		tx = params.OrderParams.Apply(tx, allowedSortFields, "users.created_at desc")

		return tx
	}

	userPage, err := h.userService.Page(req.Context(), req.NewPaginate(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(userPage)
}

// Search godoc
// @Summary      GetContent all user with pagination
// @Description  GetContent all user with pagination
// @Tags 		 user
// @ID           search-user
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        params query user_dto.UserParams false "user information"
// @Param        limit  query int    false "Limit the number of results" default(20)
// @Success      200 {object} []user_dto.User "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /user/search [get]
func (h *userHandler) Search(ctx echo.Context) error {

	req := request.Request(ctx)

	var params user_dto.UserParams
	{
		if err := req.Bind(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if params.Fullname != nil {
			fullname := fmt.Sprintf("%%%s%%", *params.Fullname)
			tx = tx.Where("(users.first_name ILIKE ? OR users.middle_name ILIKE ? OR users.last_name ILIKE ?)", fullname, fullname, fullname)
		}

		if params.Gender.Valid() {
			tx = tx.Where("users.gender = ?", params.Gender.String())
		}

		return tx
	}

	users, err := h.userService.Find(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(users)
}

// GetById godoc
// @Summary      GetContent user by ID
// @Description  GetContent user by ID
// @Tags 		 user
// @ID           get-user-by-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "user ID"
// @Success      200 {object} user_dto.User "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /user/{id} [get]
func (h *userHandler) GetByID(ctx echo.Context) error {

	req := request.Request(ctx)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.
			Joins("LEFT JOIN roles ON roles.id = users.role_id").
			Joins("LEFT JOIN organizations ON organizations.id = users.organization_id").
			Select(
				"users.id AS id",
				"users.username",
				"users.first_name",
				"users.last_name",
				"users.middle_name",
				"users.date_of_birth",
				"users.gender",
				"users.role_id",
				"organizations.soato_id AS soato_id",
				"organizations.id AS org_id",
				"roles.pages AS pages",
				"roles.permissions AS permissions",
			).Where("users.id = ?", id)
	}

	user, err := h.userService.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(user)
}

// Update godoc
// @Summary      Update user credentials
// @Description  Update username and/or password for a user
// @Tags 		 user
// @ID           update-user
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path int true "User ID"
// @Param        user body user_dto.UserUpdate true "User credentials to update"
// @Success      204 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      404 {object} response.HttpSuccess "User not found"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /user/{id} [patch]
func (h *userHandler) Update(ctx echo.Context) error {

	req := request.Request(ctx)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	var userUpdate user_dto.UserUpdate
	{
		if err := req.BindBody(&userUpdate); err != nil {
			return req.BadRequest(err)
		}
	}

	if err := h.userService.Update(ctx, id, &userUpdate); err != nil {
		return req.BadRequest(err)
	}

	return req.NoContent()
}

// Restore godoc
// @Summary      Restore user by id
// @Description  Restore user by id
// @Tags 		 user
// @ID           restore-user
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id  path     int    true  "User ID"
// @Success      204 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /user/{id}/restore [patch]
func (h *userHandler) Restore(ctx echo.Context) error {

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

	if err := h.userService.DeleteOrRestore(ctx, filter); err != nil {
		return req.BadRequest(err)
	}

	return req.NoContent()
}
