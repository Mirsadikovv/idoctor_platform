package file_handler

import (
	"fmt"

	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	file_dto "github.com/Mirsadikovv/idoctor_platform/src/module/file_service/dto"
	file_service "github.com/Mirsadikovv/idoctor_platform/src/module/file_service/service"
	log_service "github.com/Mirsadikovv/idoctor_platform/src/module/log_service/service"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type fileHandler struct {
	db          *gorm.DB
	log         logger.Logger
	fileService file_service.FileService
}

func NewFileHandler(group *echo.Group, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	handler := &fileHandler{
		db:          db,
		log:         log,
		fileService: file_service.NewFileService(db),
	}

	fileAuthMiddleware := authMiddleware.BuildMiddleware()
	fileGroup := group.Group("/file")
	{
		fileGroup.POST("/create", handler.Create, fileAuthMiddleware)
		fileGroup.POST("/create_many", handler.CreateMany, fileAuthMiddleware)
		fileGroup.GET("/search", handler.Search)
		fileGroup.GET("/page", handler.Page)
		fileGroup.GET("/download/:id", handler.File)
		fileGroup.GET("/download/:category/:owner", handler.FileByOwner)
		fileGroup.GET("/:id", handler.FindByID)
		fileGroup.DELETE("/:id", handler.Delete, fileAuthMiddleware)
		fileGroup.PATCH("/:id/replace", handler.Replace, fileAuthMiddleware)
	}

}

// Create godoc
// @Summary      file
// @Description  file
// @Tags 		 file
// @ID           create-file
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        input formData file_dto.FileCreate true "file information"
// @Param 		 file formData file true "Order Document"
// @Success      201 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /file/create [post]
func (f *fileHandler) Create(c echo.Context) error {

	req := request.RequestWithData[any](c)

	var fileDto file_dto.FileCreate
	{
		if err := req.BindBody(&fileDto); err != nil {
			return req.BadRequest(err)
		}
	}

	id, err := f.fileService.Create(c, &fileDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	data := map[string]any{
		"id":   id,
		"data": fileDto,
	}

	if _, errLog := log_service.TableCrud(f.db, c, "files", data); errLog != nil {
		fmt.Println("log error:", errLog)
	}

	return req.Created(response.NewID(id))
}

// CreateMany godoc
// @Summary      upload multiple files
// @Description  upload multiple files
// @Tags 		 file
// @ID           create-many-files
// @Accept       multipart/form-data
// @Produce      json
// @Security     ApiKeyAuth
// @Param        files formData []file true "Multiple Files"
// @Param        category formData string true "File category"
// @Param        owner formData string true "File owner"
// @Param        sign formData string false "File sign"
// @Success      201 {object} []response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /file/create_many [post]
func (f *fileHandler) CreateMany(c echo.Context) error {

	req := request.RequestWithData[any](c)

	form, err := c.MultipartForm()
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	files := form.File["files"]
	{
		if len(files) == 0 {
			return req.BadRequest("no files provided")
		}
	}

	var (
		category string
		owner    string
		sign     string
	)

	if categories := form.Value["category"]; len(categories) > 0 {
		category = categories[0]
	}
	if owners := form.Value["owner"]; len(owners) > 0 {
		owner = owners[0]
	}
	if signs := form.Value["sign"]; len(signs) > 0 {
		sign = signs[0]
	}

	fileDto := file_dto.FileCreateMany{
		FileHeaders: files,
		Category:    category,
		Owner:       owner,
		Sign:        sign,
	}

	fileIDs, err := f.fileService.CreateMany(c, &fileDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	{
		logData := map[string]any{
			"category":   fileDto.Category,
			"owner":      fileDto.Owner,
			"sign":       fileDto.Sign,
			"file_count": len(fileDto.FileHeaders),
			"file_ids":   fileIDs,
		}
		log_service.TableCrud(f.db, c, "files", &logData)
	}

	fileIDResponses := make([]response.ID64, len(fileIDs))
	for i, id := range fileIDs {
		fileIDResponses[i] = response.NewID(id)
	}

	return req.Created(fileIDResponses)
}

// GetById godoc
// @Summary      GetContent user by ID
// @Description  GetContent user by ID
// @Tags 		 file
// @ID           download-file-by-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "file ID"
// @Success      200 {object} response.HttpSuccess
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /file/download/{id} [get]
func (f *fileHandler) File(c echo.Context) error {

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

	file, err := f.fileService.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return c.File(file.BuildFilePath())
}

// FileByOwner godoc
// @Summary      GetContent user by ID
// @Description  GetContent user by ID
// @Tags 		 file
// @ID           file-by-owner
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        category path string true "category"
// @Param        owner path string true "owner"
// @Param        sign query string false "sign"
// @Success      200 {object} response.HttpSuccess
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /file/download/{category}/{owner} [get]
func (f *fileHandler) FileByOwner(c echo.Context) error {

	req := request.RequestWithData[any](c)

	var (
		owner string
	)

	category := req.Param("category")
	{
		if category == "" {
			return req.BadRequest("category is required")
		}
	}

	owner = req.Param("owner")
	{
		if owner == "" {
			return req.BadRequest("owner is required")
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if sign := req.Query("sign"); sign != "" {
			tx = tx.Where("sign = ?", sign)
		}

		return tx.Where("category = ?", category).Where("owner = ?", owner)
	}

	file, err := f.fileService.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return c.File(file.BuildFilePath())
}

func (f *fileHandler) Search(c echo.Context) error {
	return nil
}

func (f *fileHandler) FindByID(c echo.Context) error {
	return nil
}

// Page godoc
// @Summary      Get all files with pagination
// @Description  Get all files with pagination
// @Tags         file
// @ID           get-all-files
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        page query string false "Page number" default(1)
// @Param        perpage query string false "Number of items per page" default(10)
// @Param        params query file_dto.FileParams false "Searching by params"
// @Success      200 {object} []file_dto.File "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /file/page [get]
func (f *fileHandler) Page(c echo.Context) error {

	req := request.Request(c)

	var (
		params file_dto.FileParams
	)
	{
		if err := req.Bind(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if params.Category != nil {
			tx = tx.Where("files.category = ?", *params.Category)
		}

		if params.Owner != nil {
			tx = tx.Where("files.owner = ?", *params.Owner)
		}

		if params.Sign != nil {
			tx = tx.Where("files.sign = ?", *params.Sign)
		}

		return tx
	}

	filePage, err := f.fileService.Page(req.Context(), req.NewPaginate(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(filePage)
}

// Delete godoc
// @Summary      Delete file
// @Description  Delete file by ID
// @Tags 		 file
// @ID           delete-file
// @Accept       json
// @Security     ApiKeyAuth
// @Param        id path string true "file ID"
// @Success      204 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /file/{id} [delete]
func (f *fileHandler) Delete(c echo.Context) error {

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

	if err := f.fileService.Delete(c, filter); err != nil {
		return req.BadRequest(err)
	}

	data := map[string]any{
		"id":   id,
		"data": nil,
	}

	if _, errLog := log_service.TableCrud(f.db, c, "files", data); errLog != nil {
		fmt.Println("log error:", errLog)
	}

	return req.NoContent()
}

// Replace godoc
// @Summary      file
// @Description  file
// @Tags 		 file
// @ID           replace-file
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        input formData file_dto.FileReplace true "file information"
// @Param 		 file formData file true "Order Document"
// @Param        id path string true "file ID"
// @Success      201 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /file/{id}/replace [patch]
func (f *fileHandler) Replace(c echo.Context) error {

	req := request.RequestWithData[any](c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	var fileDto file_dto.FileReplace
	{
		if err := req.BindBody(&fileDto); err != nil {
			return req.BadRequest(err)
		}
	}

	err = f.fileService.Replace(c, id, &fileDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}
	data := map[string]any{
		"id":   id,
		"data": fileDto,
	}

	if _, errLog := log_service.TableCrud(f.db, c, "files", data); errLog != nil {
		fmt.Println("log error:", errLog)
	}

	return req.NoContent()
}

// FileByOwner godoc
// @Summary      GetContent user by ID
// @Description  GetContent user by ID
// @Tags 		 file
// @ID           file-by-owner
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        category path string true "category"
// @Param        owner path string true "owner"
// @Param        sign query string false "sign"
// @Success      200 {object} response.HttpSuccess
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /file/download/{category}/{owner} [get]
func (f *fileHandler) FileByEmployeeId(c echo.Context) error {

	req := request.RequestWithData[any](c)

	var (
		owner string
	)

	filter := func(tx *gorm.DB) *gorm.DB {

		if sign := req.Query("sign"); sign != "" {
			tx = tx.Where("files.sign = ?", sign)
		}

		return tx.
			Joins("LEFT JOIN hr_orders ON hr_orders.id = files.owner").
			Joins("LEFT JOIN employees ON employees.id = hr_orders.employee_id").
			Where("files.category = 'hr_orders'").
			Where("employees.id = ?", owner)
	}

	file, err := f.fileService.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return c.File(file.BuildFilePath())
}

// Page godoc
// @Summary      Get all files with pagination
// @Description  Get all files with pagination
// @Tags         file
// @ID           get-files-by-employee
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        page query string false "Page number" default(1)
// @Param        perpage query string false "Number of items per page" default(10)
// @Param        params query file_dto.EmployeeFileParams false "Searching by params"
// @Success      200 {object} []file_dto.File "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /file/page/employee [get]
func (f *fileHandler) PageByEmployee(c echo.Context) error {

	req := request.Request(c)

	var (
		params file_dto.EmployeeFileParams
	)
	{
		if err := req.Bind(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if params.EmployeeId != nil {
			tx = tx.Where("hr_orders.employee_id = ?", *params.EmployeeId)
		}

		if params.Sign != nil {
			tx = tx.Where("files.sign = ?", *params.Sign)
		}

		return tx.
			Select("files.*").
			Joins("LEFT JOIN hr_orders ON hr_orders.id = files.owner::int").
			Where("files.category = 'hr_orders'")
	}

	filePage, err := f.fileService.Page(req.Context(), req.NewPaginate(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(filePage)
}
