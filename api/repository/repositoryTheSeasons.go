package repository

import "github.com/msterzhang/onelist/api/models"

type TheSeasonRepository interface {
	Save(models.TheSeason) (models.TheSeason, error)
	FindAll(page int, size int) ([]models.TheSeason, int, error)
	FindByID(string) (models.TheSeason, error)
	UpdateByID(string, models.TheSeason) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.TheSeason, int, error)
}
