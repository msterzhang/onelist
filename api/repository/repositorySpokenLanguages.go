package repository

import "github.com/msterzhang/onelist/api/models"

type SpokenLanguageRepository interface {
	Save(models.SpokenLanguage) (models.SpokenLanguage, error)
	FindAll(page int, size int) ([]models.SpokenLanguage, int, error)
	FindByID(string) (models.SpokenLanguage, error)
	UpdateByID(string, models.SpokenLanguage) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.SpokenLanguage, int, error)
}
