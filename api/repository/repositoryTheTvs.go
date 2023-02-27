package repository

import "github.com/msterzhang/onelist/api/models"

type TheTvRepository interface {
	Save(models.TheTv) (models.TheTv, error)
	FindAll(page int, size int) ([]models.TheTv, int, error)
	FindByID(string) (models.TheTv, error)
	UpdateByID(string, models.TheTv) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.TheTv, int, error)
	FindByGalleryId(string, int, int) ([]models.TheTv, int, error)
}
