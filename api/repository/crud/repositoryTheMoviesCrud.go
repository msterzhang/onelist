package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"
	"github.com/msterzhang/onelist/config"

	"gorm.io/gorm"
)

// RepositoryTheMoviesCRUD is the struct for the TheMovie CRUD
type RepositoryTheMoviesCRUD struct {
	db *gorm.DB
}

// NewRepositoryTheMoviesCRUD returns a new repository with DB connection
func NewRepositoryTheMoviesCRUD(db *gorm.DB) *RepositoryTheMoviesCRUD {
	return &RepositoryTheMoviesCRUD{db}
}

// Save returns a new themovie created or an error
func (r *RepositoryTheMoviesCRUD) Save(themovie models.TheMovie) (models.TheMovie, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.TheMovie{}).Create(&themovie).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return themovie, nil
	}
	return models.TheMovie{}, err
}

// UpdateByID update themovie from the DB
func (r *RepositoryTheMoviesCRUD) UpdateByID(id string, themovie models.TheMovie) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.TheMovie{}).Where("id = ?", id).Select("*").Updates(&themovie)
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

// DeleteByID themovie by the id
func (r *RepositoryTheMoviesCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.TheMovie{}).Where("id = ?", id).Delete(&models.TheMovie{})
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

// FindAll returns all the themovies from the DB
func (r *RepositoryTheMoviesCRUD) FindAll(page int, size int) ([]models.TheMovie, int, error) {
	var err error
	var num int64
	themovies := []models.TheMovie{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.TheMovie{}).Find(&themovies)
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&themovies).Error
		} else {
			err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&themovies).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return themovies, int(num), nil
	}
	return nil, 0, err
}

// FindByID return themovie from the DB
func (r *RepositoryTheMoviesCRUD) FindByID(id string) (models.TheMovie, error) {
	var err error
	themovie := models.TheMovie{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.TheMovie{}).Where("id = ?", id).Preload("ProductionCompanies").Preload("SpokenLanguages").Preload("ProductionCountries").Preload("Genres").Preload("ThePersons").Preload("BelongsToCollection").Preload("TheCredit").Take(&themovie).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return themovie, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.TheMovie{}, errors.New("themovie Not Found")
	}
	return models.TheMovie{}, err
}

// Search themovie from the DB
func (r *RepositoryTheMoviesCRUD) Search(q string, page int, size int) ([]models.TheMovie, int, error) {
	var err error
	var num int64
	themovies := []models.TheMovie{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.TheMovie{}).Where("title LIKE ?", "%"+q+"%")
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&themovies).Error
		} else {
			err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&themovies).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return themovies, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.TheMovie{}, 0, errors.New("themovies Not Found")
	}
	return []models.TheMovie{}, 0, err
}

// FindByGalleryId themovies from the DB
func (r *RepositoryTheMoviesCRUD) FindByGalleryId(id string, page int, size int) ([]models.TheMovie, int, error) {
	var err error
	var num int64
	themovies := []models.TheMovie{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.TheMovie{}).Where("gallery_uid = ?", id)
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&themovies).Error
		} else {
			err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&themovies).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return themovies, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.TheMovie{}, 0, errors.New("themovies Not Found")
	}
	return []models.TheMovie{}, 0, err
}
