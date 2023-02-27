package repository

import "github.com/msterzhang/onelist/api/models"

type LastEpisodeToAirRepository interface {
	Save(models.LastEpisodeToAir) (models.LastEpisodeToAir, error)
	FindAll(page int, size int) ([]models.LastEpisodeToAir, int, error)
	FindByID(string) (models.LastEpisodeToAir, error)
	UpdateByID(string, models.LastEpisodeToAir) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.LastEpisodeToAir, int, error)
}
