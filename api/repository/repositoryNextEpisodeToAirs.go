package repository

import "github.com/msterzhang/onelist/api/models"

type NextEpisodeToAirRepository interface {
	Save(models.NextEpisodeToAir) (models.NextEpisodeToAir, error)
	FindAll(page int, size int) ([]models.NextEpisodeToAir, int, error)
	FindByID(string) (models.NextEpisodeToAir, error)
	UpdateByID(string, models.NextEpisodeToAir) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.NextEpisodeToAir, int, error)
}
