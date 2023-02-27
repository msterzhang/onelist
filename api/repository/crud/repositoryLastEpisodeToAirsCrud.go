package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositoryLastEpisodeToAirsCRUD is the struct for the LastEpisodeToAir CRUD
type RepositoryLastEpisodeToAirsCRUD struct {
	db *gorm.DB
}

// NewRepositoryLastEpisodeToAirsCRUD returns a new repository with DB connection
func NewRepositoryLastEpisodeToAirsCRUD(db *gorm.DB) *RepositoryLastEpisodeToAirsCRUD {
	return &RepositoryLastEpisodeToAirsCRUD{db}
}

// Save returns a new lastepisodetoair created or an error
func (r *RepositoryLastEpisodeToAirsCRUD) Save(lastepisodetoair models.LastEpisodeToAir) (models.LastEpisodeToAir, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.LastEpisodeToAir{}).Create(&lastepisodetoair).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return lastepisodetoair, nil
	}
	return models.LastEpisodeToAir{}, err
}

// UpdateByID update lastepisodetoair from the DB
func (r *RepositoryLastEpisodeToAirsCRUD) UpdateByID(id string, lastepisodetoair models.LastEpisodeToAir) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.LastEpisodeToAir{}).Where("id = ?", id).Select("*").Updates(&lastepisodetoair)
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

// DeleteByID lastepisodetoair by the id
func (r *RepositoryLastEpisodeToAirsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.LastEpisodeToAir{}).Where("id = ?", id).Delete(&models.LastEpisodeToAir{})
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

// FindAll returns all the lastepisodetoairs from the DB
func (r *RepositoryLastEpisodeToAirsCRUD) FindAll(page int, size int) ([]models.LastEpisodeToAir, int, error) {
	var err error
	var num int64
	lastepisodetoairs := []models.LastEpisodeToAir{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.LastEpisodeToAir{}).Find(&lastepisodetoairs)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&lastepisodetoairs).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return lastepisodetoairs, int(num), nil
	}
	return nil, 0, err
}

// FindByID return lastepisodetoair from the DB
func (r *RepositoryLastEpisodeToAirsCRUD) FindByID(id string) (models.LastEpisodeToAir, error) {
	var err error
	lastepisodetoair := models.LastEpisodeToAir{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.LastEpisodeToAir{}).Where("id = ?", id).Take(&lastepisodetoair).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return lastepisodetoair, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.LastEpisodeToAir{}, errors.New("lastepisodetoair Not Found")
	}
	return models.LastEpisodeToAir{}, err
}

// Search lastepisodetoair from the DB
func (r *RepositoryLastEpisodeToAirsCRUD) Search(q string, page int, size int) ([]models.LastEpisodeToAir, int, error) {
	var err error
	var num int64
	lastepisodetoairs := []models.LastEpisodeToAir{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.LastEpisodeToAir{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&lastepisodetoairs).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return lastepisodetoairs, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.LastEpisodeToAir{}, 0, errors.New("lastepisodetoairs Not Found")
	}
	return []models.LastEpisodeToAir{}, 0, err
}
