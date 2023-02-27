package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"
	"github.com/msterzhang/onelist/config"

	"gorm.io/gorm"
)

// RepositoryTheTvsCRUD is the struct for the TheTv CRUD
type RepositoryTheTvsCRUD struct {
	db *gorm.DB
}

// NewRepositoryTheTvsCRUD returns a new repository with DB connection
func NewRepositoryTheTvsCRUD(db *gorm.DB) *RepositoryTheTvsCRUD {
	return &RepositoryTheTvsCRUD{db}
}

// Save returns a new thetv created or an error
func (r *RepositoryTheTvsCRUD) Save(thetv models.TheTv) (models.TheTv, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.TheTv{}).Create(&thetv).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return thetv, nil
	}
	return models.TheTv{}, err
}

// UpdateByID update thetv from the DB
func (r *RepositoryTheTvsCRUD) UpdateByID(id string, thetv models.TheTv) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.TheTv{}).Where("id = ?", id).Select("*").Updates(&thetv)
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

// DeleteByID thetv by the id
func (r *RepositoryTheTvsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.TheTv{}).Where("id = ?", id).Delete(&models.TheTv{})
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

// FindAll returns all the thetvs from the DB
func (r *RepositoryTheTvsCRUD) FindAll(page int, size int) ([]models.TheTv, int, error) {
	var err error
	var num int64
	thetvs := []models.TheTv{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.TheTv{}).Find(&thetvs)
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&thetvs).Error
		} else {
			err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&thetvs).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return thetvs, int(num), nil
	}
	return nil, 0, err
}

// FindByID return thetv from the DB
func (r *RepositoryTheTvsCRUD) FindByID(id string) (models.TheTv, error) {
	var err error
	thetv := models.TheTv{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.TheTv{}).Where("id = ?", id).Preload("ProductionCompanies").Preload("SpokenLanguages").Preload("ProductionCountries").Preload("Genres").Preload("TheSeasons").Preload("Seasons").Preload("ThePersons").Preload("Networks").Preload("TheCredit").Preload("LastEpisodeToAir").Preload("NextEpisodeToAir").Take(&thetv).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return thetv, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.TheTv{}, errors.New("thetv Not Found")
	}
	return models.TheTv{}, err
}

// Search thetv from the DB
func (r *RepositoryTheTvsCRUD) Search(q string, page int, size int) ([]models.TheTv, int, error) {
	var err error
	var num int64
	thetvs := []models.TheTv{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.TheTv{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&thetvs).Error
		} else {
			err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&thetvs).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return thetvs, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.TheTv{}, 0, errors.New("thetvs Not Found")
	}
	return []models.TheTv{}, 0, err
}

// FindByGalleryId thetv from the DB
func (r *RepositoryTheTvsCRUD) FindByGalleryId(id string, page int, size int) ([]models.TheTv, int, error) {
	var err error
	var num int64
	thetvs := []models.TheTv{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.TheTv{}).Where("gallery_uid = ?", id)
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&thetvs).Error
		} else {
			err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&thetvs).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return thetvs, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.TheTv{}, 0, errors.New("thetvs Not Found")
	}
	return []models.TheTv{}, 0, err
}
