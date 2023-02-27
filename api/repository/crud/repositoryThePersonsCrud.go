package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositoryThePersonsCRUD is the struct for the ThePerson CRUD
type RepositoryThePersonsCRUD struct {
	db *gorm.DB
}

// NewRepositoryThePersonsCRUD returns a new repository with DB connection
func NewRepositoryThePersonsCRUD(db *gorm.DB) *RepositoryThePersonsCRUD {
	return &RepositoryThePersonsCRUD{db}
}

// Save returns a new theperson created or an error
func (r *RepositoryThePersonsCRUD) Save(theperson models.ThePerson) (models.ThePerson, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.ThePerson{}).Create(&theperson).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return theperson, nil
	}
	return models.ThePerson{}, err
}

// UpdateByID update theperson from the DB
func (r *RepositoryThePersonsCRUD) UpdateByID(id string, theperson models.ThePerson) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.ThePerson{}).Where("id = ?", id).Select("*").Updates(&theperson)
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

// DeleteByID theperson by the id
func (r *RepositoryThePersonsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.ThePerson{}).Where("id = ?", id).Delete(&models.ThePerson{})
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

// FindAll returns all the thepersons from the DB
func (r *RepositoryThePersonsCRUD) FindAll(page int, size int) ([]models.ThePerson, int, error) {
	var err error
	var num int64
	thepersons := []models.ThePerson{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.ThePerson{}).Find(&thepersons)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&thepersons).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return thepersons, int(num), nil
	}
	return nil, 0, err
}

// FindByID return theperson from the DB
func (r *RepositoryThePersonsCRUD) FindByID(id string) (models.ThePerson, error) {
	var err error
	theperson := models.ThePerson{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.ThePerson{}).Where("id = ?", id).Preload("TheMovies").Preload("TheTvs").Take(&theperson).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return theperson, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.ThePerson{}, errors.New("theperson Not Found")
	}
	return models.ThePerson{}, err
}

// Search theperson from the DB
func (r *RepositoryThePersonsCRUD) Search(q string, page int, size int) ([]models.ThePerson, int, error) {
	var err error
	var num int64
	thepersons := []models.ThePerson{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.ThePerson{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&thepersons).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return thepersons, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.ThePerson{}, 0, errors.New("thepersons Not Found")
	}
	return []models.ThePerson{}, 0, err
}
