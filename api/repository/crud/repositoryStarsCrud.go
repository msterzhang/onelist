package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositoryStarsCRUD is the struct for the Star CRUD
type RepositoryStarsCRUD struct {
	db *gorm.DB
}

// NewRepositoryStarsCRUD returns a new repository with DB connection
func NewRepositoryStarsCRUD(db *gorm.DB) *RepositoryStarsCRUD {
	return &RepositoryStarsCRUD{db}
}

// Save returns a new star created or an error
func (r *RepositoryStarsCRUD) Save(star models.Star) (models.Star, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Star{}).Create(&star).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return star, nil
	}
	return models.Star{}, err
}

// UpdateByID update star from the DB
func (r *RepositoryStarsCRUD) UpdateByID(id string, star models.Star) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Star{}).Where("id = ?", id).Select("*").Updates(&star)
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

// DeleteByID star by the id
func (r *RepositoryStarsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Star{}).Where("id = ?", id).Delete(&models.Star{})
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

// ReNew star by the Star
func (r *RepositoryStarsCRUD) ReNewStarByStar(star models.Star) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		starDb := models.Star{}
		err := r.db.Model(&models.Star{}).Where("user_id = ? AND data_id = ? AND data_type = ?", star.UserId, star.DataId, star.DataType).Take(&starDb).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) && starDb.Id != 0 {
			rs = r.db.Model(&models.Star{}).Where("user_id = ? AND data_id = ? AND data_type = ?", star.UserId, star.DataId, star.DataType).Delete(&models.Star{})
		} else {
			rs = r.db.Model(&models.Star{}).Create(&star)
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

// FindAll returns all the stars from the DB
func (r *RepositoryStarsCRUD) FindAll(page int, size int) ([]models.Star, int, error) {
	var err error
	var num int64
	stars := []models.Star{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Star{}).Find(&stars)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&stars).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return stars, int(num), nil
	}
	return nil, 0, err
}

// FindByID return star from the DB
func (r *RepositoryStarsCRUD) FindByID(id string) (models.Star, error) {
	var err error
	star := models.Star{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Star{}).Where("id = ?", id).Take(&star).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return star, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Star{}, errors.New("star Not Found")
	}
	return models.Star{}, err
}

// Search star from the DB
func (r *RepositoryStarsCRUD) Search(q string, page int, size int) ([]models.Star, int, error) {
	var err error
	var num int64
	stars := []models.Star{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Star{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&stars).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return stars, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.Star{}, 0, errors.New("stars Not Found")
	}
	return []models.Star{}, 0, err
}

func (r *RepositoryStarsCRUD) FindAllByUser(star models.Star, page int, size int) ([]models.Star, int, error) {
	var err error
	var num int64
	stars := []models.Star{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Star{}).Where("user_id = ? AND data_type = ?", star.UserId, star.DataType).Find(&stars)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&stars).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return stars, int(num), nil
	}
	return nil, 0, err
}
