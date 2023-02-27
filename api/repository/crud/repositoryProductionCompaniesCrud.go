package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositoryProductionCompaniesCRUD is the struct for the ProductionCompanie CRUD
type RepositoryProductionCompaniesCRUD struct {
	db *gorm.DB
}

// NewRepositoryProductionCompaniesCRUD returns a new repository with DB connection
func NewRepositoryProductionCompaniesCRUD(db *gorm.DB) *RepositoryProductionCompaniesCRUD {
	return &RepositoryProductionCompaniesCRUD{db}
}

// Save returns a new productioncompanie created or an error
func (r *RepositoryProductionCompaniesCRUD) Save(productioncompanie models.ProductionCompanie) (models.ProductionCompanie, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.ProductionCompanie{}).Create(&productioncompanie).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return productioncompanie, nil
	}
	return models.ProductionCompanie{}, err
}

// UpdateByID update productioncompanie from the DB
func (r *RepositoryProductionCompaniesCRUD) UpdateByID(id string, productioncompanie models.ProductionCompanie) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.ProductionCompanie{}).Where("id = ?", id).Select("*").Updates(&productioncompanie)
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

// DeleteByID productioncompanie by the id
func (r *RepositoryProductionCompaniesCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.ProductionCompanie{}).Where("id = ?", id).Delete(&models.ProductionCompanie{})
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

// FindAll returns all the productioncompanies from the DB
func (r *RepositoryProductionCompaniesCRUD) FindAll(page int, size int) ([]models.ProductionCompanie, int, error) {
	var err error
	var num int64
	productioncompanies := []models.ProductionCompanie{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.ProductionCompanie{}).Find(&productioncompanies)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&productioncompanies).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return productioncompanies, int(num), nil
	}
	return nil, 0, err
}

// FindByID return productioncompanie from the DB
func (r *RepositoryProductionCompaniesCRUD) FindByID(id string) (models.ProductionCompanie, error) {
	var err error
	productioncompanie := models.ProductionCompanie{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.ProductionCompanie{}).Where("id = ?", id).Take(&productioncompanie).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return productioncompanie, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.ProductionCompanie{}, errors.New("productioncompanie Not Found")
	}
	return models.ProductionCompanie{}, err
}

// Search productioncompanie from the DB
func (r *RepositoryProductionCompaniesCRUD) Search(q string, page int, size int) ([]models.ProductionCompanie, int, error) {
	var err error
	var num int64
	productioncompanies := []models.ProductionCompanie{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.ProductionCompanie{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&productioncompanies).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return productioncompanies, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.ProductionCompanie{}, 0, errors.New("productioncompanies Not Found")
	}
	return []models.ProductionCompanie{}, 0, err
}
