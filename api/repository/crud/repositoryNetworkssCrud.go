package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositoryNetworkssCRUD is the struct for the Networks CRUD
type RepositoryNetworkssCRUD struct {
	db *gorm.DB
}

// NewRepositoryNetworkssCRUD returns a new repository with DB connection
func NewRepositoryNetworkssCRUD(db *gorm.DB) *RepositoryNetworkssCRUD {
	return &RepositoryNetworkssCRUD{db}
}

// Save returns a new networks created or an error
func (r *RepositoryNetworkssCRUD) Save(networks models.Networks) (models.Networks, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Networks{}).Create(&networks).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return networks, nil
	}
	return models.Networks{}, err
}

// UpdateByID update networks from the DB
func (r *RepositoryNetworkssCRUD) UpdateByID(id string, networks models.Networks) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Networks{}).Where("id = ?", id).Select("*").Updates(&networks)
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

// DeleteByID networks by the id
func (r *RepositoryNetworkssCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Networks{}).Where("id = ?", id).Delete(&models.Networks{})
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

// FindAll returns all the networkss from the DB
func (r *RepositoryNetworkssCRUD) FindAll(page int, size int) ([]models.Networks, int, error) {
	var err error
	var num int64
	networkss := []models.Networks{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Networks{}).Find(&networkss)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&networkss).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return networkss, int(num), nil
	}
	return nil, 0, err
}

// FindByID return networks from the DB
func (r *RepositoryNetworkssCRUD) FindByID(id string) (models.Networks, error) {
	var err error
	networks := models.Networks{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Networks{}).Where("id = ?", id).Take(&networks).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return networks, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Networks{}, errors.New("networks Not Found")
	}
	return models.Networks{}, err
}

// Search networks from the DB
func (r *RepositoryNetworkssCRUD) Search(q string, page int, size int) ([]models.Networks, int, error) {
	var err error
	var num int64
	networkss := []models.Networks{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Networks{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&networkss).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return networkss, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.Networks{}, 0, errors.New("networkss Not Found")
	}
	return []models.Networks{}, 0, err
}
