package repository

import "github.com/msterzhang/onelist/api/models"

type ErrFileRepository interface {
	Save(models.ErrFile) (models.ErrFile, error)
	FindAll(page int, size int) ([]models.ErrFile, int, error)
	FindByID(string) (models.ErrFile, error)
	UpdateByID(string, models.ErrFile) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.ErrFile, int, error)
	GetErrFilesByWorkId(string, int, int) ([]models.ErrFile, int, error)
}
