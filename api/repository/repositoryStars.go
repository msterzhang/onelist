package repository

import "github.com/msterzhang/onelist/api/models"

type StarRepository interface {
	Save(models.Star) (models.Star, error)
	FindAll(page int, size int) ([]models.Star, int, error)
	FindAllByUser(star models.Star, page int, size int) ([]models.Star, int, error)
	FindByID(string) (models.Star, error)
	UpdateByID(string, models.Star) (int64, error)
	ReNewStarByStar(models.Star) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.Star, int, error)
}
