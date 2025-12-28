package problem_handler

import (
	"fmt"

	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	problem_dto "github.com/Mirsadikovv/idoctor_platform/src/module/problem_service/dto"
	problem_service "github.com/Mirsadikovv/idoctor_platform/src/module/problem_service/service"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type problemHandler struct {
	db             *gorm.DB
	log            logger.Logger
	authMiddleware *auth_middleware.AuthMiddleware
	problemService problem_service.ProblemService
}

func NewProblemHandler(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	handler := &problemHandler{
		db:             db,
		log:            log,
		authMiddleware: authMiddleware,
		problemService: problem_service.NewProblemService(db),
	}

	problemServiceMiddleware := handler.authMiddleware.BuildMiddleware()
	problemGroup := router.Group("/api/v1/problem", problemServiceMiddleware)
	{
		problemGroup.GET("/:id", handler.FindById)
		problemGroup.GET("/search", handler.Search)
		problemGroup.GET("/page", handler.Page)
		problemGroup.POST("", handler.Create)
		problemGroup.PUT("/:id", handler.Update)
		problemGroup.DELETE("/:id", handler.Delete)
	}
}

// Search godoc
// @Summary      Get all problems with search
// @Description  Get all problems with search
// @Tags 		 problem
// @Id           search-problem
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        params query problem_dto.ProblemParams false "params"
// @Param        limit  query int    false "Limit the number of results" default(20)
// @Success      200 {object} []problem_dto.Problem "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /problem/search [get]
func (h *problemHandler) Search(c echo.Context) error {
	req := request.RequestWithData[any](c)

	var params problem_dto.ProblemParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		if params.Name != nil {
			name := fmt.Sprintf("%%%s%%", *params.Name)
			tx = tx.Where("problems.name ILIKE ?", name)
		}

		if params.MinPrice != nil {
			tx = tx.Where("problems.price >= ?", *params.MinPrice)
		}

		if params.MaxPrice != nil {
			tx = tx.Where("problems.price <= ?", *params.MaxPrice)
		}

		return tx.Select(
			"problems.id",
			"problems.name",
			"problems.price",
			"problems.created_at",
		).Order("problems.id DESC")
	}

	problems, err := h.problemService.Find(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(problems)
}

// Page 			godoc
// @Summary			problem-service-page
// @Description		problem-service-page
// @Tags 			problem
// @Param       	filter query problem_dto.ProblemParams false "Filter parameters"
// @Param       	page query int false "Page number"
// @Param       	perpage query int false "Items per page"
// @Security     	ApiKeyAuth
// @Success	200 	{object} problem_dto.ProblemPage "Successful operation"
// @Failure	400 	{object} response.HttpSuccess "Bad request"
// @Failure	500 	{object} response.HttpSuccess "Internal server error"
// @Router			/problem/page [get]
func (h *problemHandler) Page(c echo.Context) error {
	req := request.Request(c)

	var params problem_dto.ProblemParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	tx := func(tx *gorm.DB) *gorm.DB {
		if params.Name != nil {
			tx = tx.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *params.Name))
		}

		if params.MinPrice != nil {
			tx = tx.Where("price >= ?", *params.MinPrice)
		}

		if params.MaxPrice != nil {
			tx = tx.Where("price <= ?", *params.MaxPrice)
		}

		return tx.Select(
			"problems.id",
			"problems.name",
			"problems.price",
			"problems.created_at",
		).Order("problems.id DESC")
	}

	items, err := h.problemService.Page(req.Context(), req.NewPaginate(), tx)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(items)
}

// GetById godoc
// @Summary      Get problem by Id
// @Description  Get problem by Id
// @Tags 		 problem
// @Id           get-problem-by-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "problem Id"
// @Success      200 {object} problem_dto.Problem "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /problem/{id} [get]
func (h *problemHandler) FindById(c echo.Context) error {
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

	problem, err := h.problemService.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(problem)
}

// Create godoc
// @Summary      Create problem
// @Description  Create new problem
// @Tags 		 problem
// @Id           create-problem
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        input body problem_dto.ProblemCreate true "problem information"
// @Success      201 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /problem [POST]
func (h *problemHandler) Create(c echo.Context) error {
	req := request.Request(c)

	var problemDto problem_dto.ProblemCreate
	{
		if err := req.BindBody(&problemDto); err != nil {
			return req.BadRequest(err)
		}
	}

	problemDto.Id = 0 // Ensure Id is 0 for creation

	id, err := h.problemService.CreateOrUpdate(c.Request().Context(), &problemDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(response.NewID(id))
}

// Update godoc
// @Summary      Update problem
// @Description  Update existing problem
// @Tags 		 problem
// @Id           update-problem
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "problem Id"
// @Param        input body problem_dto.ProblemUpdate true "problem information"
// @Success      200 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /problem/{id} [PUT]
func (h *problemHandler) Update(c echo.Context) error {
	req := request.Request(c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	var problemDto problem_dto.ProblemUpdate
	{
		if err := req.BindBody(&problemDto); err != nil {
			return req.BadRequest(err)
		}
	}

	createDto := &problem_dto.ProblemCreate{
		Id:    id,
		Name:  problemDto.Name,
		Price: problemDto.Price,
	}

	updatedId, err := h.problemService.CreateOrUpdate(c.Request().Context(), createDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(response.NewID(updatedId))
}

// Delete godoc
// @Summary      Delete problem
// @Description  Delete problem by Id
// @Tags 		 problem
// @Id           delete-problem
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "problem Id"
// @Success      200 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /problem/{id} [delete]
func (h *problemHandler) Delete(c echo.Context) error {
	req := request.Request(c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	err = h.problemService.Delete(req.Context(), id)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.NoContent()
}
