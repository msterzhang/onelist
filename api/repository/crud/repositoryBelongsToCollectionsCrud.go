package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositoryBelongsToCollectionsCRUD is the struct for the BelongsToCollection CRUD
type RepositoryBelongsToCollectionsCRUD struct {
	db *gorm.DB
}

// NewRepositoryBelongsToCollectionsCRUD returns a new repository with DB connection
func NewRepositoryBelongsToCollectionsCRUD(db *gorm.DB) *RepositoryBelongsToCollectionsCRUD {
	return &RepositoryBelongsToCollectionsCRUD{db}
}

// Save returns a new belongstocollection created or an error
func (r *RepositoryBelongsToCollectionsCRUD) Save(belongstocollection models.BelongsToCollection) (models.BelongsToCollection, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.BelongsToCollection{}).Create(&belongstocollection).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return belongstocollection, nil
	}
	return models.BelongsToCollection{}, err
}

// UpdateByID update belongstocollection from the DB
func (r *RepositoryBelongsToCollectionsCRUD) UpdateByID(id string, belongstocollection models.BelongsToCollection) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.BelongsToCollection{}).Where("id = ?", id).Select("*").Updates(&belongstocollection)
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

// DeleteByID belongstocollection by the id
func (r *RepositoryBelongsToCollectionsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.BelongsToCollection{}).Where("id = ?", id).Delete(&models.BelongsToCollection{})
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

// FindAll returns all the belongstocollections from the DB
func (r *RepositoryBelongsToCollectionsCRUD) FindAll(page int, size int) ([]models.BelongsToCollection, int, error) {
	var err error
	var num int64
	belongstocollections := []models.BelongsToCollection{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.BelongsToCollection{}).Find(&belongstocollections)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&belongstocollections).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return belongstocollections, int(num), nil
	}
	return nil, 0, err
}

// FindByID return belongstocollection from the DB
func (r *RepositoryBelongsToCollectionsCRUD) FindByID(id string) (models.BelongsToCollection, error) {
	var err error
	belongstocollection := models.BelongsToCollection{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.BelongsToCollection{}).Where("id = ?", id).Take(&belongstocollection).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return belongstocollection, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.BelongsToCollection{}, errors.New("belongstocollection Not Found")
	}
	return models.BelongsToCollection{}, err
}

// Search belongstocollection from the DB
func (r *RepositoryBelongsToCollectionsCRUD) Search(q string, page int, size int) ([]models.BelongsToCollection, int, error) {
	var err error
	var num int64
	belongstocollections := []models.BelongsToCollection{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.BelongsToCollection{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&belongstocollections).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return belongstocollections, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.BelongsToCollection{}, 0, errors.New("belongstocollections Not Found")
	}
	return []models.BelongsToCollection{}, 0, err
}
