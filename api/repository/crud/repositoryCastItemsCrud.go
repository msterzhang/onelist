package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositoryCastItemsCRUD is the struct for the CastItem CRUD
type RepositoryCastItemsCRUD struct {
	db *gorm.DB
}

// NewRepositoryCastItemsCRUD returns a new repository with DB connection
func NewRepositoryCastItemsCRUD(db *gorm.DB) *RepositoryCastItemsCRUD {
	return &RepositoryCastItemsCRUD{db}
}

// Save returns a new castitem created or an error
func (r *RepositoryCastItemsCRUD) Save(castitem models.CastItem) (models.CastItem, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.CastItem{}).Create(&castitem).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return castitem, nil
	}
	return models.CastItem{}, err
}

// UpdateByID update castitem from the DB
func (r *RepositoryCastItemsCRUD) UpdateByID(id string, castitem models.CastItem) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.CastItem{}).Where("id = ?", id).Select("*").Updates(&castitem)
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

// DeleteByID castitem by the id
func (r *RepositoryCastItemsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.CastItem{}).Where("id = ?", id).Delete(&models.CastItem{})
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

// FindAll returns all the castitems from the DB
func (r *RepositoryCastItemsCRUD) FindAll(page int, size int) ([]models.CastItem, int, error) {
	var err error
	var num int64
	castitems := []models.CastItem{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.CastItem{}).Find(&castitems)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&castitems).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return castitems, int(num), nil
	}
	return nil, 0, err
}

// FindByID return castitem from the DB
func (r *RepositoryCastItemsCRUD) FindByID(id string) (models.CastItem, error) {
	var err error
	castitem := models.CastItem{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.CastItem{}).Where("id = ?", id).Take(&castitem).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return castitem, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.CastItem{}, errors.New("castitem Not Found")
	}
	return models.CastItem{}, err
}

// Search castitem from the DB
func (r *RepositoryCastItemsCRUD) Search(q string, page int, size int) ([]models.CastItem, int, error) {
	var err error
	var num int64
	castitems := []models.CastItem{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.CastItem{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&castitems).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return castitems, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.CastItem{}, 0, errors.New("castitems Not Found")
	}
	return []models.CastItem{}, 0, err
}
