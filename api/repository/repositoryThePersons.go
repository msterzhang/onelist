package repository

import "github.com/msterzhang/onelist/api/models"

type ThePersonRepository interface {
	Save(models.ThePerson) (models.ThePerson, error)
	FindAll(page int, size int) ([]models.ThePerson, int, error)
	FindByID(string) (models.ThePerson, error)
	UpdateByID(string, models.ThePerson) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.ThePerson, int, error)
}
