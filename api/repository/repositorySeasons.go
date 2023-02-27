package repository

import "github.com/msterzhang/onelist/api/models"

type SeasonRepository interface {
	Save(models.Season) (models.Season, error)
	FindAll(page int, size int) ([]models.Season, int, error)
	FindByID(string) (models.Season, error)
	UpdateByID(string, models.Season) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.Season, int, error)
}
