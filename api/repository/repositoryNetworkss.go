package repository

import "github.com/msterzhang/onelist/api/models"

type NetworksRepository interface {
	Save(models.Networks) (models.Networks, error)
	FindAll(page int, size int) ([]models.Networks, int, error)
	FindByID(string) (models.Networks, error)
	UpdateByID(string, models.Networks) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.Networks, int, error)
}
