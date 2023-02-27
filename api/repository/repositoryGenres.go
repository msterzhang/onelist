package repository

import "github.com/msterzhang/onelist/api/models"

type GenreRepository interface {
	Save(models.Genre) (models.Genre, error)
	FindAll(page int, size int) ([]models.Genre, int, error)
	FindByID(string) (models.Genre, error)
	UpdateByID(string, models.Genre) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.Genre, int, error)
}
