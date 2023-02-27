package repository

import "github.com/msterzhang/onelist/api/models"

type HeartRepository interface {
	Save(models.Heart) (models.Heart, error)
	FindAll(page int, size int) ([]models.Heart, int, error)
	FindAllByUser(heart models.Heart, page int, size int) ([]models.Heart, int, error)
	FindByID(string) (models.Heart, error)
	ReNewHeartByHeart(models.Heart) (int64, error)
	UpdateByID(string, models.Heart) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.Heart, int, error)
}
