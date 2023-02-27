package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositorySpokenLanguagesCRUD is the struct for the SpokenLanguage CRUD
type RepositorySpokenLanguagesCRUD struct {
	db *gorm.DB
}

// NewRepositorySpokenLanguagesCRUD returns a new repository with DB connection
func NewRepositorySpokenLanguagesCRUD(db *gorm.DB) *RepositorySpokenLanguagesCRUD {
	return &RepositorySpokenLanguagesCRUD{db}
}

// Save returns a new spokenlanguage created or an error
func (r *RepositorySpokenLanguagesCRUD) Save(spokenlanguage models.SpokenLanguage) (models.SpokenLanguage, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.SpokenLanguage{}).Create(&spokenlanguage).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return spokenlanguage, nil
	}
	return models.SpokenLanguage{}, err
}

// UpdateByID update spokenlanguage from the DB
func (r *RepositorySpokenLanguagesCRUD) UpdateByID(id string, spokenlanguage models.SpokenLanguage) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.SpokenLanguage{}).Where("id = ?", id).Select("*").Updates(&spokenlanguage)
		ch <- true
	}(done)

	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}

		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

// DeleteByID spokenlanguage by the id
func (r *RepositorySpokenLanguagesCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.SpokenLanguage{}).Where("id = ?", id).Delete(&models.SpokenLanguage{})
		ch <- true
	}(done)

	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

// FindAll returns all the spokenlanguages from the DB
func (r *RepositorySpokenLanguagesCRUD) FindAll(page int, size int) ([]models.SpokenLanguage, int, error) {
	var err error
	var num int64
	spokenlanguages := []models.SpokenLanguage{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.SpokenLanguage{}).Find(&spokenlanguages)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&spokenlanguages).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return spokenlanguages, int(num), nil
	}
	return nil, 0, err
}

// FindByID return spokenlanguage from the DB
func (r *RepositorySpokenLanguagesCRUD) FindByID(id string) (models.SpokenLanguage, error) {
	var err error
	spokenlanguage := models.SpokenLanguage{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.SpokenLanguage{}).Where("id = ?", id).Take(&spokenlanguage).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return spokenlanguage, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.SpokenLanguage{}, errors.New("spokenlanguage Not Found")
	}
	return models.SpokenLanguage{}, err
}

// Search spokenlanguage from the DB
func (r *RepositorySpokenLanguagesCRUD) Search(q string, page int, size int) ([]models.SpokenLanguage, int, error) {
	var err error
	var num int64
	spokenlanguages := []models.SpokenLanguage{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.SpokenLanguage{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&spokenlanguages).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return spokenlanguages, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.SpokenLanguage{}, 0, errors.New("spokenlanguages Not Found")
	}
	return []models.SpokenLanguage{}, 0, err
}
