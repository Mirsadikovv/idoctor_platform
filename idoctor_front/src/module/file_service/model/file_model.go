package file_model

import (
	"os"
	"path/filepath"
	"time"
)

type File struct {
	Id        int64      `json:"id" gorm:"primaryKey"`
	Owner     string     `json:"owner" gorm:"not null"`
	Category  string     `json:"category" gorm:"not null"`
	Sign      string     `json:"sign"`
	DirName   string     `json:"dirName"`
	Filename  string     `json:"filename"`
	Name      string     `json:"name"`
	MimeType  string     `json:"mimeType"`
	Size      int64      `json:"size"`
	CreatedAt *time.Time `json:"createdAt" gorm:"autoCreateTime:true"`
}

func (File) TableName() string {
	return "files"
}

func (f *File) BuildFilePath() string {
	return filepath.Join("./", f.DirName, f.Category, f.Owner, f.Filename)
}

func (f *File) CanDelete() bool {
	file, err := os.Stat(f.BuildFilePath())

	if err != nil {
		return true
	}

	return !file.IsDir()
}

func (f *File) Remove() error {
	return os.Remove(f.BuildFilePath())
}
