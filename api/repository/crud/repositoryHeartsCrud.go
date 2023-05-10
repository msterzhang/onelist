package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"
	"github.com/msterzhang/onelist/config"

	"gorm.io/gorm"
)

// RepositoryHeartsCRUD is the struct for the Heart CRUD
type RepositoryHeartsCRUD struct {
	db *gorm.DB
}

// NewRepositoryHeartsCRUD returns a new repository with DB connection
func NewRepositoryHeartsCRUD(db *gorm.DB) *RepositoryHeartsCRUD {
	return &RepositoryHeartsCRUD{db}
}

// Save returns a new heart created or an error
func (r *RepositoryHeartsCRUD) Save(heart models.Heart) (models.Heart, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Heart{}).Create(&heart).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return heart, nil
	}
	return models.Heart{}, err
}

// ReNew heart by the Heart
func (r *RepositoryHeartsCRUD) ReNewHeartByHeart(heart models.Heart) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		heartDb := models.Heart{}
		err := r.db.Model(&models.Heart{}).Where("user_id = ? AND data_id = ? AND data_type = ?", heart.UserId, heart.DataId, heart.DataType).Take(&heartDb).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) && heartDb.Id != 0 {
			rs = r.db.Model(&models.Heart{}).Where("user_id = ? AND data_id = ? AND data_type = ?", heart.UserId, heart.DataId, heart.DataType).Delete(&models.Heart{})
		} else {
			rs = r.db.Model(&models.Heart{}).Create(&heart)
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

// UpdateByID update heart from the DB
func (r *RepositoryHeartsCRUD) UpdateByID(id string, heart models.Heart) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Heart{}).Where("id = ?", id).Select("*").Updates(&heart)
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

// DeleteByID heart by the id
func (r *RepositoryHeartsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Heart{}).Where("id = ?", id).Delete(&models.Heart{})
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

// FindAll returns all the hearts from the DB
func (r *RepositoryHeartsCRUD) FindAll(page int, size int) ([]models.Heart, int, error) {
	var err error
	var num int64
	hearts := []models.Heart{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Heart{}).Find(&hearts)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&hearts).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return hearts, int(num), nil
	}
	return nil, 0, err
}

// FindByID return heart from the DB
func (r *RepositoryHeartsCRUD) FindByID(id string) (models.Heart, error) {
	var err error
	heart := models.Heart{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Heart{}).Where("id = ?", id).Take(&heart).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return heart, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Heart{}, errors.New("heart Not Found")
	}
	return models.Heart{}, err
}

// Search heart from the DB
func (r *RepositoryHeartsCRUD) Search(q string, page int, size int) ([]models.Heart, int, error) {
	var err error
	var num int64
	hearts := []models.Heart{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Heart{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&hearts).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return hearts, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.Heart{}, 0, errors.New("hearts Not Found")
	}
	return []models.Heart{}, 0, err
}

func (r *RepositoryHeartsCRUD) FindAllByUser(heart models.Heart, page int, size int) ([]models.Heart, int, error) {
	var err error
	var num int64
	hearts := []models.Heart{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Heart{}).Where("user_id = ? AND data_type = ?", heart.UserId, heart.DataType).Find(&hearts)
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Order("datetime(updated_at) desc").Limit(size).Offset((page - 1) * size).Scan(&hearts).Error
		} else {
			err = result.Order("-updated_at").Limit(size).Offset((page - 1) * size).Scan(&hearts).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return hearts, int(num), nil
	}
	return nil, 0, err
}
