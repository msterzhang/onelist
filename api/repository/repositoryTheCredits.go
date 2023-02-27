package repository

import "github.com/msterzhang/onelist/api/models"

type TheCreditRepository interface {
	Save(models.TheCredit) (models.TheCredit, error)
	FindAll(page int, size int) ([]models.TheCredit, int, error)
	FindByID(string) (models.TheCredit, error)
	UpdateByID(string, models.TheCredit) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.TheCredit, int, error)
}
