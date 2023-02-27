package repository

import "github.com/msterzhang/onelist/api/models"

type GalleryRepository interface {
	Save(models.Gallery) (models.Gallery, error)
	FindAll(page int, size int) ([]models.Gallery, int, error)
	FindByID(string) (models.Gallery, error)
	FindByUID(string) (models.Gallery, error)
	UpdateByID(string, models.Gallery) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.Gallery, int, error)
}
