package repository

import "github.com/msterzhang/onelist/api/models"

type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll(page int, size int) ([]models.User, int, error)
	FindByID(string) (models.User, error)
	UpdateByID(string, models.User) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.User, int, error)
}
