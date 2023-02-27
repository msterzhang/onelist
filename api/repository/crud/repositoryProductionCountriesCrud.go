package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositoryProductionCountriesCRUD is the struct for the ProductionCountrie CRUD
type RepositoryProductionCountriesCRUD struct {
	db *gorm.DB
}

// NewRepositoryProductionCountriesCRUD returns a new repository with DB connection
func NewRepositoryProductionCountriesCRUD(db *gorm.DB) *RepositoryProductionCountriesCRUD {
	return &RepositoryProductionCountriesCRUD{db}
}

// Save returns a new productioncountrie created or an error
func (r *RepositoryProductionCountriesCRUD) Save(productioncountrie models.ProductionCountrie) (models.ProductionCountrie, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.ProductionCountrie{}).Create(&productioncountrie).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return productioncountrie, nil
	}
	return models.ProductionCountrie{}, err
}

// UpdateByID update productioncountrie from the DB
func (r *RepositoryProductionCountriesCRUD) UpdateByID(id string, productioncountrie models.ProductionCountrie) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.ProductionCountrie{}).Where("id = ?", id).Select("*").Updates(&productioncountrie)
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

// DeleteByID productioncountrie by the id
func (r *RepositoryProductionCountriesCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.ProductionCountrie{}).Where("id = ?", id).Delete(&models.ProductionCountrie{})
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

// FindAll returns all the productioncountries from the DB
func (r *RepositoryProductionCountriesCRUD) FindAll(page int, size int) ([]models.ProductionCountrie, int, error) {
	var err error
	var num int64
	productioncountries := []models.ProductionCountrie{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.ProductionCountrie{}).Find(&productioncountries)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&productioncountries).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return productioncountries, int(num), nil
	}
	return nil, 0, err
}

// FindByID return productioncountrie from the DB
func (r *RepositoryProductionCountriesCRUD) FindByID(id string) (models.ProductionCountrie, error) {
	var err error
	productioncountrie := models.ProductionCountrie{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.ProductionCountrie{}).Where("id = ?", id).Take(&productioncountrie).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return productioncountrie, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.ProductionCountrie{}, errors.New("productioncountrie Not Found")
	}
	return models.ProductionCountrie{}, err
}

// Search productioncountrie from the DB
func (r *RepositoryProductionCountriesCRUD) Search(q string, page int, size int) ([]models.ProductionCountrie, int, error) {
	var err error
	var num int64
	productioncountries := []models.ProductionCountrie{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.ProductionCountrie{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&productioncountries).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return productioncountries, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.ProductionCountrie{}, 0, errors.New("productioncountries Not Found")
	}
	return []models.ProductionCountrie{}, 0, err
}
