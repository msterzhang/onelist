package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositoryTheCreditsCRUD is the struct for the TheCredit CRUD
type RepositoryTheCreditsCRUD struct {
	db *gorm.DB
}

// NewRepositoryTheCreditsCRUD returns a new repository with DB connection
func NewRepositoryTheCreditsCRUD(db *gorm.DB) *RepositoryTheCreditsCRUD {
	return &RepositoryTheCreditsCRUD{db}
}

// Save returns a new thecredit created or an error
func (r *RepositoryTheCreditsCRUD) Save(thecredit models.TheCredit) (models.TheCredit, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.TheCredit{}).Create(&thecredit).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return thecredit, nil
	}
	return models.TheCredit{}, err
}

// UpdateByID update thecredit from the DB
func (r *RepositoryTheCreditsCRUD) UpdateByID(id string, thecredit models.TheCredit) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.TheCredit{}).Where("id = ?", id).Select("*").Updates(&thecredit)
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

// DeleteByID thecredit by the id
func (r *RepositoryTheCreditsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.TheCredit{}).Where("id = ?", id).Delete(&models.TheCredit{})
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

// FindAll returns all the thecredits from the DB
func (r *RepositoryTheCreditsCRUD) FindAll(page int, size int) ([]models.TheCredit, int, error) {
	var err error
	var num int64
	thecredits := []models.TheCredit{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.TheCredit{}).Find(&thecredits)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&thecredits).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return thecredits, int(num), nil
	}
	return nil, 0, err
}

// FindByID return thecredit from the DB
func (r *RepositoryTheCreditsCRUD) FindByID(id string) (models.TheCredit, error) {
	var err error
	thecredit := models.TheCredit{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.TheCredit{}).Where("id = ?", id).Take(&thecredit).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return thecredit, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.TheCredit{}, errors.New("thecredit Not Found")
	}
	return models.TheCredit{}, err
}

// Search thecredit from the DB
func (r *RepositoryTheCreditsCRUD) Search(q string, page int, size int) ([]models.TheCredit, int, error) {
	var err error
	var num int64
	thecredits := []models.TheCredit{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.TheCredit{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&thecredits).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return thecredits, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.TheCredit{}, 0, errors.New("thecredits Not Found")
	}
	return []models.TheCredit{}, 0, err
}
