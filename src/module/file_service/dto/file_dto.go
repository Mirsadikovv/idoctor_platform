package file_dto

import (
	"mime/multipart"
	"path/filepath"
	"time"

	"git.sriss.uz/shared/shared_service/response"
)

type FileCreate struct {
	FileHeader *multipart.FileHeader `form:"file"`
	Category   string                `form:"category"`
	Owner      string                `form:"owner"`
	Sign       string                `form:"sign"`
} // @name FileCreate

type FileCreateMany struct {
	FileHeaders []*multipart.FileHeader `form:"files"`
	Category    string                  `form:"category"`
	Owner       string                  `form:"owner"`
	Sign        string                  `form:"sign"`
} // @name FileCreateMany

type FilePage = response.PageData[FileDto]

type File struct {
	DirName  string
	Category string
	Owner    string
	Filename string
	Name     string
	MimeType string
	Size     int64
} // @name File

type FileReplace struct {
	FileHeader *multipart.FileHeader `form:"file"`
} // @name FileCreate

func (f *File) BuildFilePath() string {
	return filepath.Join("./", f.DirName, f.Category, f.Owner, f.Filename)
}

type FileParams struct {
	Sign     *string `query:"sign"`
	Category *string `query:"category"`
	Owner    *string `query:"owner"`
} // @name FileParams

type FileDto struct {
	Id        int64      `json:"id" gorm:"column:id"`
	DirName   string     `json:"dirName"`
	Category  string     `json:"category"`
	Owner     string     `json:"owner" gorm:"column:owner"`
	Sign      string     `json:"sign"`
	Filename  string     `json:"filename"`
	Name      string     `json:"name"`
	MimeType  string     `json:"mimeType"`
	Size      int64      `json:"size"`
	CreatedAt *time.Time `json:"createdAt"`
} // @name File

type EmployeeFileParams struct {
	Sign       *string `query:"sign"`
	EmployeeId *int64  `query:"employeeId"`
} // @name EmployeeFileParams
