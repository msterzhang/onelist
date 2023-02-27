package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositoryGenresCRUD is the struct for the Genre CRUD
type RepositoryGenresCRUD struct {
	db *gorm.DB
}

// NewRepositoryGenresCRUD returns a new repository with DB connection
func NewRepositoryGenresCRUD(db *gorm.DB) *RepositoryGenresCRUD {
	return &RepositoryGenresCRUD{db}
}

// Save returns a new genre created or an error
func (r *RepositoryGenresCRUD) Save(genre models.Genre) (models.Genre, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Genre{}).Create(&genre).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return genre, nil
	}
	return models.Genre{}, err
}

// UpdateByID update genre from the DB
func (r *RepositoryGenresCRUD) UpdateByID(id string, genre models.Genre) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Genre{}).Where("id = ?", id).Select("*").Updates(&genre)
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

// DeleteByID genre by the id
func (r *RepositoryGenresCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Genre{}).Where("id = ?", id).Delete(&models.Genre{})
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

// FindAll returns all the genres from the DB
func (r *RepositoryGenresCRUD) FindAll(page int, size int) ([]models.Genre, int, error) {
	var err error
	var num int64
	genres := []models.Genre{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Genre{}).Find(&genres)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&genres).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return genres, int(num), nil
	}
	return nil, 0, err
}

// FindByID return genre from the DB
func (r *RepositoryGenresCRUD) FindByID(id string) (models.Genre, error) {
	var err error
	genre := models.Genre{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Genre{}).Where("id = ?", id).Take(&genre).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return genre, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Genre{}, errors.New("genre Not Found")
	}
	return models.Genre{}, err
}

// Search genre from the DB
func (r *RepositoryGenresCRUD) Search(q string, page int, size int) ([]models.Genre, int, error) {
	var err error
	var num int64
	genres := []models.Genre{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Genre{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&genres).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return genres, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.Genre{}, 0, errors.New("genres Not Found")
	}
	return []models.Genre{}, 0, err
}
