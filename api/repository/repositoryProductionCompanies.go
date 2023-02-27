package repository

import "github.com/msterzhang/onelist/api/models"

type ProductionCompanieRepository interface {
	Save(models.ProductionCompanie) (models.ProductionCompanie, error)
	FindAll(page int, size int) ([]models.ProductionCompanie, int, error)
	FindByID(string) (models.ProductionCompanie, error)
	UpdateByID(string, models.ProductionCompanie) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.ProductionCompanie, int, error)
}
