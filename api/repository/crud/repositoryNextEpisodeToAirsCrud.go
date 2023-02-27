package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositoryNextEpisodeToAirsCRUD is the struct for the NextEpisodeToAir CRUD
type RepositoryNextEpisodeToAirsCRUD struct {
	db *gorm.DB
}

// NewRepositoryNextEpisodeToAirsCRUD returns a new repository with DB connection
func NewRepositoryNextEpisodeToAirsCRUD(db *gorm.DB) *RepositoryNextEpisodeToAirsCRUD {
	return &RepositoryNextEpisodeToAirsCRUD{db}
}

// Save returns a new nextepisodetoair created or an error
func (r *RepositoryNextEpisodeToAirsCRUD) Save(nextepisodetoair models.NextEpisodeToAir) (models.NextEpisodeToAir, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.NextEpisodeToAir{}).Create(&nextepisodetoair).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return nextepisodetoair, nil
	}
	return models.NextEpisodeToAir{}, err
}

// UpdateByID update nextepisodetoair from the DB
func (r *RepositoryNextEpisodeToAirsCRUD) UpdateByID(id string, nextepisodetoair models.NextEpisodeToAir) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.NextEpisodeToAir{}).Where("id = ?", id).Select("*").Updates(&nextepisodetoair)
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

// DeleteByID nextepisodetoair by the id
func (r *RepositoryNextEpisodeToAirsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.NextEpisodeToAir{}).Where("id = ?", id).Delete(&models.NextEpisodeToAir{})
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

// FindAll returns all the nextepisodetoairs from the DB
func (r *RepositoryNextEpisodeToAirsCRUD) FindAll(page int, size int) ([]models.NextEpisodeToAir, int, error) {
	var err error
	var num int64
	nextepisodetoairs := []models.NextEpisodeToAir{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.NextEpisodeToAir{}).Find(&nextepisodetoairs)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&nextepisodetoairs).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return nextepisodetoairs, int(num), nil
	}
	return nil, 0, err
}

// FindByID return nextepisodetoair from the DB
func (r *RepositoryNextEpisodeToAirsCRUD) FindByID(id string) (models.NextEpisodeToAir, error) {
	var err error
	nextepisodetoair := models.NextEpisodeToAir{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.NextEpisodeToAir{}).Where("id = ?", id).Take(&nextepisodetoair).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return nextepisodetoair, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.NextEpisodeToAir{}, errors.New("nextepisodetoair Not Found")
	}
	return models.NextEpisodeToAir{}, err
}

// Search nextepisodetoair from the DB
func (r *RepositoryNextEpisodeToAirsCRUD) Search(q string, page int, size int) ([]models.NextEpisodeToAir, int, error) {
	var err error
	var num int64
	nextepisodetoairs := []models.NextEpisodeToAir{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.NextEpisodeToAir{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&nextepisodetoairs).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return nextepisodetoairs, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.NextEpisodeToAir{}, 0, errors.New("nextepisodetoairs Not Found")
	}
	return []models.NextEpisodeToAir{}, 0, err
}
