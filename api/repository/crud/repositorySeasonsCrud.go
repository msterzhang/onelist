package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositorySeasonsCRUD is the struct for the Season CRUD
type RepositorySeasonsCRUD struct {
	db *gorm.DB
}

// NewRepositorySeasonsCRUD returns a new repository with DB connection
func NewRepositorySeasonsCRUD(db *gorm.DB) *RepositorySeasonsCRUD {
	return &RepositorySeasonsCRUD{db}
}

// Save returns a new season created or an error
func (r *RepositorySeasonsCRUD) Save(season models.Season) (models.Season, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Season{}).Create(&season).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return season, nil
	}
	return models.Season{}, err
}

// UpdateByID update season from the DB
func (r *RepositorySeasonsCRUD) UpdateByID(id string, season models.Season) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Season{}).Where("id = ?", id).Select("*").Updates(&season)
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

// DeleteByID season by the id
func (r *RepositorySeasonsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Season{}).Where("id = ?", id).Delete(&models.Season{})
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

// FindAll returns all the seasons from the DB
func (r *RepositorySeasonsCRUD) FindAll(page int, size int) ([]models.Season, int, error) {
	var err error
	var num int64
	seasons := []models.Season{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Season{}).Find(&seasons)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&seasons).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return seasons, int(num), nil
	}
	return nil, 0, err
}

// FindByID return season from the DB
func (r *RepositorySeasonsCRUD) FindByID(id string) (models.Season, error) {
	var err error
	season := models.Season{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Season{}).Where("id = ?", id).Take(&season).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return season, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Season{}, errors.New("season Not Found")
	}
	return models.Season{}, err
}

// Search season from the DB
func (r *RepositorySeasonsCRUD) Search(q string, page int, size int) ([]models.Season, int, error) {
	var err error
	var num int64
	seasons := []models.Season{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Season{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&seasons).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return seasons, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.Season{}, 0, errors.New("seasons Not Found")
	}
	return []models.Season{}, 0, err
}
