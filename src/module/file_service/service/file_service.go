package file_service

import (
	"context"
	"errors"
	"io"
	"os"
	"path"

	file_dto "github.com/Mirsadikovv/idoctor_platform/src/module/file_service/dto"
	file_model "github.com/Mirsadikovv/idoctor_platform/src/module/file_service/model"
	"github.com/labstack/echo/v4"

	"github.com/Mirsadikovv/shared/pg"
	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/sharedutil"
	"gorm.io/gorm"
)

type FileService interface {
	Create(ctx echo.Context, fileDto *file_dto.FileCreate) (int64, error)
	CreateMany(ctx echo.Context, fileDto *file_dto.FileCreateMany) ([]int64, error)
	FindOne(ctx context.Context, filter pg.Filter) (*file_dto.File, error)
	Find(ctx context.Context, filter pg.Filter) ([]file_dto.File, error)
	Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*file_dto.FilePage, error)
	Delete(ctx echo.Context, filter pg.Filter) error
	Replace(ctx echo.Context, fileID int64, fileDto *file_dto.FileReplace) error
}

type fileService struct {
	db *gorm.DB
}

func NewFileService(db *gorm.DB) FileService {
	return &fileService{
		db: db,
	}
}

func (f *fileService) FindOne(ctx context.Context, filter pg.Filter) (*file_dto.File, error) {
	return pg.FindOneWithScan[file_model.File, file_dto.File](f.db, filter)
}

func (f *fileService) Find(ctx context.Context, filter pg.Filter) ([]file_dto.File, error) {
	return pg.FindWithScan[file_model.File, file_dto.File](f.db, filter)
}

func (f *fileService) Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*file_dto.FilePage, error) {
	return pg.PageWithScan[file_model.File, file_dto.FileDto](f.db, paginate, filter)
}

func (s *fileService) Create(ctx echo.Context, fileDto *file_dto.FileCreate) (int64, error) {
	const FILES = "files"

	var (
		fileHeader = fileDto.FileHeader
		category   = fileDto.Category
		owner      = fileDto.Owner
		sign       = fileDto.Sign
	)

	ownerDirname := path.Join("./", FILES, category, owner)
	{
		if err := os.MkdirAll(ownerDirname, 0755); err != nil {
			return 0, err
		}
	}

	ownerFilename := sharedutil.RndWithExt(fileHeader.Filename)

	ownerFullFilePath := path.Join(ownerDirname, ownerFilename)

	multipart, err := fileHeader.Open()
	{
		if err != nil {
			return 0, err
		}
		defer multipart.Close()
	}

	ownerFile, err := os.Create(ownerFullFilePath)
	{
		if err != nil {
			return 0, err
		}
		defer ownerFile.Close()
	}

	fileModel := &file_model.File{
		DirName:  FILES,
		Category: category,
		Filename: path.Base(ownerFilename),
		Name:     fileHeader.Filename,
		Sign:     sign,
		MimeType: path.Ext(fileHeader.Filename),
		Owner:    owner,
		Size:     fileHeader.Size,
	}

	err = pg.Transaction(s.db, func(tx *gorm.DB) error {

		if err := tx.WithContext(ctx.Request().Context()).Create(fileModel).Error; err != nil {
			return err
		}

		if _, err := io.Copy(ownerFile, multipart); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		// Удаляем файл, если копирование не удалось
		os.Remove(ownerFullFilePath)
		return 0, err
	}

	return fileModel.Id, nil
}

func (s *fileService) Delete(ctx echo.Context, filter pg.Filter) error {

	return pg.Transaction(s.db, func(tx *gorm.DB) error {

		var fileModel file_model.File
		{
			if err := pg.Delete(tx.WithContext(ctx.Request().Context()), &fileModel, filter, "id", "dir_name", "category", "owner", "filename"); err != nil {
				return err
			}
		}

		if !fileModel.CanDelete() {
			return errors.New("can not delete")
		}

		return fileModel.Remove()
	})

}

func (s *fileService) Replace(ctx echo.Context, fileID int64, fileDto *file_dto.FileReplace) error {
	const FILES = "files"

	var (
		existingFile file_model.File
		oldFilePath  string
	)
	{
		if err := s.db.Where("id = ?", fileID).First(&existingFile).Error; err != nil {
			return err
		}

		oldFilePath = path.Join("./", FILES, existingFile.Category, existingFile.Owner, existingFile.Filename)
	}

	var (
		fileHeader = fileDto.FileHeader
		category   = existingFile.Category
		owner      = existingFile.Owner
		sign       = existingFile.Sign
	)

	ownerDirname := path.Join("./", FILES, category, owner)
	{
		if err := os.MkdirAll(ownerDirname, 0755); err != nil {
			return err
		}
	}

	ownerFilename := sharedutil.RndWithExt(fileHeader.Filename)

	ownerFullFilePath := path.Join(ownerDirname, ownerFilename)

	multipart, err := fileHeader.Open()
	{
		if err != nil {
			return err
		}
		defer multipart.Close()
	}

	ownerFile, err := os.Create(ownerFullFilePath)
	{
		if err != nil {
			return err
		}
		defer ownerFile.Close()
	}

	fileModel := &file_model.File{
		DirName:  FILES,
		Category: category,
		Filename: path.Base(ownerFilename),
		Sign:     sign,
		MimeType: path.Ext(fileHeader.Filename),
		Owner:    owner,
		Size:     fileHeader.Size,
	}

	err = pg.Transaction(s.db, func(tx *gorm.DB) error {

		if err := tx.WithContext(ctx.Request().Context()).Model(&file_model.File{}).Where("id = ?", fileID).Updates(fileModel).Error; err != nil {
			return err
		}

		if _, err := io.Copy(ownerFile, multipart); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		os.Remove(ownerFullFilePath)
		return err
	}

	if oldFilePath != "" {
		os.Remove(oldFilePath)
	}

	return nil
}

func (s *fileService) CreateMany(ctx echo.Context, fileDto *file_dto.FileCreateMany) ([]int64, error) {
	const FILES = "files"

	var (
		fileHeaders = fileDto.FileHeaders
		category    = fileDto.Category
		owner       = fileDto.Owner
		sign        = fileDto.Sign
		fileIDs     = make([]int64, 0, len(fileHeaders))
		fileModels  = make([]*file_model.File, 0, len(fileHeaders))
		filePaths   = make([]string, 0, len(fileHeaders))
	)

	if len(fileHeaders) == 0 {
		return fileIDs, errors.New("no files provided")
	}

	ownerDirname := path.Join("./", FILES, category, owner)
	{
		if err := os.MkdirAll(ownerDirname, 0755); err != nil {
			return nil, err
		}
	}

	for _, fileHeader := range fileHeaders {
		ownerFilename := sharedutil.RndWithExt(fileHeader.Filename)
		ownerFullFilePath := path.Join(ownerDirname, ownerFilename)

		fileModel := &file_model.File{
			DirName:  FILES,
			Category: category,
			Filename: path.Base(ownerFilename),
			Name:     fileHeader.Filename,
			Sign:     sign,
			MimeType: path.Ext(fileHeader.Filename),
			Owner:    owner,
			Size:     fileHeader.Size,
		}

		fileModels = append(fileModels, fileModel)
		filePaths = append(filePaths, ownerFullFilePath)
	}

	err := pg.Transaction(s.db, func(tx *gorm.DB) error {
		for i, fileModel := range fileModels {
			if err := tx.WithContext(ctx.Request().Context()).Create(fileModel).Error; err != nil {
				return err
			}
			fileIDs = append(fileIDs, fileModel.Id)

			multipart, err := fileHeaders[i].Open()
			if err != nil {
				return err
			}
			defer multipart.Close()

			ownerFile, err := os.Create(filePaths[i])
			if err != nil {
				return err
			}
			defer ownerFile.Close()

			if _, err := io.Copy(ownerFile, multipart); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		for _, filePath := range filePaths {
			os.Remove(filePath)
		}
		return nil, err
	}

	return fileIDs, nil
}
