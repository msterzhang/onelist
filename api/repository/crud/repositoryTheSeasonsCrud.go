package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositoryTheSeasonsCRUD is the struct for the TheSeason CRUD
type RepositoryTheSeasonsCRUD struct {
	db *gorm.DB
}

// NewRepositoryTheSeasonsCRUD returns a new repository with DB connection
func NewRepositoryTheSeasonsCRUD(db *gorm.DB) *RepositoryTheSeasonsCRUD {
	return &RepositoryTheSeasonsCRUD{db}
}

// Save returns a new theseason created or an error
func (r *RepositoryTheSeasonsCRUD) Save(theseason models.TheSeason) (models.TheSeason, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.TheSeason{}).Create(&theseason).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return theseason, nil
	}
	return models.TheSeason{}, err
}

// UpdateByID update theseason from the DB
func (r *RepositoryTheSeasonsCRUD) UpdateByID(id string, theseason models.TheSeason) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.TheSeason{}).Where("id = ?", id).Select("*").Updates(&theseason)
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

// DeleteByID theseason by the id
func (r *RepositoryTheSeasonsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.TheSeason{}).Where("id = ?", id).Delete(&models.TheSeason{})
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

// FindAll returns all the theseasons from the DB
func (r *RepositoryTheSeasonsCRUD) FindAll(page int, size int) ([]models.TheSeason, int, error) {
	var err error
	var num int64
	theseasons := []models.TheSeason{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.TheSeason{}).Find(&theseasons)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&theseasons).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return theseasons, int(num), nil
	}
	return nil, 0, err
}

// FindByID return theseason from the DB
func (r *RepositoryTheSeasonsCRUD) FindByID(id string) (models.TheSeason, error) {
	var err error
	theseason := models.TheSeason{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.TheSeason{}).Where("id = ?", id).Preload("Episodes").Take(&theseason).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return theseason, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.TheSeason{}, errors.New("theseason Not Found")
	}
	return models.TheSeason{}, err
}

// Search theseason from the DB
func (r *RepositoryTheSeasonsCRUD) Search(q string, page int, size int) ([]models.TheSeason, int, error) {
	var err error
	var num int64
	theseasons := []models.TheSeason{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.TheSeason{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&theseasons).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return theseasons, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.TheSeason{}, 0, errors.New("theseasons Not Found")
	}
	return []models.TheSeason{}, 0, err
}
