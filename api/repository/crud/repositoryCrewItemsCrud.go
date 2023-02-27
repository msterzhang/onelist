package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositoryCrewItemsCRUD is the struct for the CrewItem CRUD
type RepositoryCrewItemsCRUD struct {
	db *gorm.DB
}

// NewRepositoryCrewItemsCRUD returns a new repository with DB connection
func NewRepositoryCrewItemsCRUD(db *gorm.DB) *RepositoryCrewItemsCRUD {
	return &RepositoryCrewItemsCRUD{db}
}

// Save returns a new crewitem created or an error
func (r *RepositoryCrewItemsCRUD) Save(crewitem models.CrewItem) (models.CrewItem, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.CrewItem{}).Create(&crewitem).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return crewitem, nil
	}
	return models.CrewItem{}, err
}

// UpdateByID update crewitem from the DB
func (r *RepositoryCrewItemsCRUD) UpdateByID(id string, crewitem models.CrewItem) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.CrewItem{}).Where("id = ?", id).Select("*").Updates(&crewitem)
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

// DeleteByID crewitem by the id
func (r *RepositoryCrewItemsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.CrewItem{}).Where("id = ?", id).Delete(&models.CrewItem{})
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

// FindAll returns all the crewitems from the DB
func (r *RepositoryCrewItemsCRUD) FindAll(page int, size int) ([]models.CrewItem, int, error) {
	var err error
	var num int64
	crewitems := []models.CrewItem{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.CrewItem{}).Find(&crewitems)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&crewitems).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return crewitems, int(num), nil
	}
	return nil, 0, err
}

// FindByID return crewitem from the DB
func (r *RepositoryCrewItemsCRUD) FindByID(id string) (models.CrewItem, error) {
	var err error
	crewitem := models.CrewItem{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.CrewItem{}).Where("id = ?", id).Take(&crewitem).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return crewitem, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.CrewItem{}, errors.New("crewitem Not Found")
	}
	return models.CrewItem{}, err
}

// Search crewitem from the DB
func (r *RepositoryCrewItemsCRUD) Search(q string, page int, size int) ([]models.CrewItem, int, error) {
	var err error
	var num int64
	crewitems := []models.CrewItem{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.CrewItem{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&crewitems).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return crewitems, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.CrewItem{}, 0, errors.New("crewitems Not Found")
	}
	return []models.CrewItem{}, 0, err
}
