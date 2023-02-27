package repository

import "github.com/msterzhang/onelist/api/models"

type PlayedRepository interface {
	Save(models.Played) (models.Played, error)
	ReNewByPlayed(models.Played) (int64, error)
	FindAll(page int, size int) ([]models.Played, int, error)
	FindAllByUser(played models.Played, page int, size int) ([]models.Played, int, error)
	FindByID(string) (models.Played, error)
	UpdateByID(string, models.Played) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.Played, int, error)
}
