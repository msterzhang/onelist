package repository

import "github.com/msterzhang/onelist/api/models"

type EpisodeRepository interface {
	Save(models.Episode) (models.Episode, error)
	FindAll(page int, size int) ([]models.Episode, int, error)
	FindByID(string) (models.Episode, error)
	UpdateByID(string, models.Episode) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.Episode, int, error)
}
