package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositoryEpisodesCRUD is the struct for the Episode CRUD
type RepositoryEpisodesCRUD struct {
	db *gorm.DB
}

// NewRepositoryEpisodesCRUD returns a new repository with DB connection
func NewRepositoryEpisodesCRUD(db *gorm.DB) *RepositoryEpisodesCRUD {
	return &RepositoryEpisodesCRUD{db}
}

// Save returns a new episode created or an error
func (r *RepositoryEpisodesCRUD) Save(episode models.Episode) (models.Episode, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Episode{}).Create(&episode).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return episode, nil
	}
	return models.Episode{}, err
}

// UpdateByID update episode from the DB
func (r *RepositoryEpisodesCRUD) UpdateByID(id string, episode models.Episode) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Episode{}).Where("id = ?", id).Select("*").Updates(&episode)
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

// DeleteByID episode by the id
func (r *RepositoryEpisodesCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Episode{}).Where("id = ?", id).Delete(&models.Episode{})
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

// FindAll returns all the episodes from the DB
func (r *RepositoryEpisodesCRUD) FindAll(page int, size int) ([]models.Episode, int, error) {
	var err error
	var num int64
	episodes := []models.Episode{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Episode{}).Find(&episodes)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&episodes).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return episodes, int(num), nil
	}
	return nil, 0, err
}

// FindByID return episode from the DB
func (r *RepositoryEpisodesCRUD) FindByID(id string) (models.Episode, error) {
	var err error
	episode := models.Episode{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Episode{}).Where("id = ?", id).Take(&episode).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return episode, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Episode{}, errors.New("episode Not Found")
	}
	return models.Episode{}, err
}

// Search episode from the DB
func (r *RepositoryEpisodesCRUD) Search(q string, page int, size int) ([]models.Episode, int, error) {
	var err error
	var num int64
	episodes := []models.Episode{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Episode{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&episodes).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return episodes, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.Episode{}, 0, errors.New("episodes Not Found")
	}
	return []models.Episode{}, 0, err
}
