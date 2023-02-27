package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"
	"github.com/msterzhang/onelist/config"

	"gorm.io/gorm"
)

// RepositoryWorksCRUD is the struct for the Work CRUD
type RepositoryWorksCRUD struct {
	db *gorm.DB
}

// NewRepositoryWorksCRUD returns a new repository with DB connection
func NewRepositoryWorksCRUD(db *gorm.DB) *RepositoryWorksCRUD {
	return &RepositoryWorksCRUD{db}
}

// Save returns a new work created or an error
func (r *RepositoryWorksCRUD) Save(work models.Work) (models.Work, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Work{}).Create(&work).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return work, nil
	}
	return models.Work{}, err
}

// UpdateByID update work from the DB
func (r *RepositoryWorksCRUD) UpdateByID(id string, work models.Work) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Work{}).Where("id = ?", id).Select("*").Updates(&work)
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

// DeleteByID work by the id
func (r *RepositoryWorksCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Work{}).Where("id = ?", id).Delete(&models.Work{})
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

// FindAll returns all the works from the DB
func (r *RepositoryWorksCRUD) FindAll(page int, size int) ([]models.Work, int, error) {
	var err error
	var num int64
	works := []models.Work{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Work{}).Find(&works)
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&works).Error
		} else {
			err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&works).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return works, int(num), nil
	}
	return nil, 0, err
}

// FindByID return work from the DB
func (r *RepositoryWorksCRUD) FindByID(id string) (models.Work, error) {
	var err error
	work := models.Work{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Work{}).Where("id = ?", id).Take(&work).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return work, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Work{}, errors.New("work Not Found")
	}
	return models.Work{}, err
}

// Search work from the DB
func (r *RepositoryWorksCRUD) Search(q string, page int, size int) ([]models.Work, int, error) {
	var err error
	var num int64
	works := []models.Work{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Work{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&works).Error
		} else {
			err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&works).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return works, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.Work{}, 0, errors.New("works Not Found")
	}
	return []models.Work{}, 0, err
}

// Search work from the DB
func (r *RepositoryWorksCRUD) GetWorkListByGalleryId(id string, page int, size int) ([]models.Work, int, error) {
	var err error
	var num int64
	works := []models.Work{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Work{}).Where("gallery_uid = ?", id)
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&works).Error
		} else {
			err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&works).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return works, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.Work{}, 0, errors.New("works Not Found")
	}
	return []models.Work{}, 0, err
}
