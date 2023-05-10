package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"
	"github.com/msterzhang/onelist/config"

	"gorm.io/gorm"
)

// RepositoryPlayedsCRUD is the struct for the Played CRUD
type RepositoryPlayedsCRUD struct {
	db *gorm.DB
}

// NewRepositoryPlayedsCRUD returns a new repository with DB connection
func NewRepositoryPlayedsCRUD(db *gorm.DB) *RepositoryPlayedsCRUD {
	return &RepositoryPlayedsCRUD{db}
}

// Save returns a new played created or an error
func (r *RepositoryPlayedsCRUD) Save(played models.Played) (models.Played, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Played{}).Create(&played).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return played, nil
	}
	return models.Played{}, err
}

// UpdateByID update played from the DB
func (r *RepositoryPlayedsCRUD) UpdateByID(id string, played models.Played) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Played{}).Where("id = ?", id).Select("*").Updates(&played)
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

// DeleteByID played by the id
func (r *RepositoryPlayedsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Played{}).Where("id = ?", id).Delete(&models.Played{})
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

// ReNew played by the Played
func (r *RepositoryPlayedsCRUD) ReNewByPlayed(played models.Played) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		playedDb := models.Played{}
		err := r.db.Model(&models.Played{}).Where("user_id = ? AND data_id = ? AND data_type = ?", played.UserId, played.DataId, played.DataType).Take(&playedDb).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) && playedDb.Id != 0 {
			rs = r.db.Model(&models.Played{}).Where("user_id = ? AND data_id = ? AND data_type = ?", played.UserId, played.DataId, played.DataType).Delete(&models.Played{})
		} else {
			rs = r.db.Model(&models.Played{}).Create(&played)
		}
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

// FindAll returns all the playeds from the DB
func (r *RepositoryPlayedsCRUD) FindAll(page int, size int) ([]models.Played, int, error) {
	var err error
	var num int64
	playeds := []models.Played{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Played{}).Find(&playeds)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&playeds).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return playeds, int(num), nil
	}
	return nil, 0, err
}

// FindByID return played from the DB
func (r *RepositoryPlayedsCRUD) FindByID(id string) (models.Played, error) {
	var err error
	played := models.Played{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Played{}).Where("id = ?", id).Take(&played).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return played, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Played{}, errors.New("played Not Found")
	}
	return models.Played{}, err
}

// Search played from the DB
func (r *RepositoryPlayedsCRUD) Search(q string, page int, size int) ([]models.Played, int, error) {
	var err error
	var num int64
	playeds := []models.Played{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Played{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&playeds).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return playeds, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.Played{}, 0, errors.New("playeds Not Found")
	}
	return []models.Played{}, 0, err
}

func (r *RepositoryPlayedsCRUD) FindAllByUser(played models.Played, page int, size int) ([]models.Played, int, error) {
	var err error
	var num int64
	playeds := []models.Played{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Debug().Model(&models.Played{}).Where("user_id = ? AND data_type = ?", played.UserId, played.DataType).Find(&playeds)
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&playeds).Error
		} else {
			err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&playeds).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return playeds, int(num), nil
	}
	return nil, 0, err
}
