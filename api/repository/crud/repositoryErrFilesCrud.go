package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"
	"github.com/msterzhang/onelist/config"

	"gorm.io/gorm"
)

// RepositoryErrFilesCRUD is the struct for the ErrFile CRUD
type RepositoryErrFilesCRUD struct {
	db *gorm.DB
}

// NewRepositoryErrFilesCRUD returns a new repository with DB connection
func NewRepositoryErrFilesCRUD(db *gorm.DB) *RepositoryErrFilesCRUD {
	return &RepositoryErrFilesCRUD{db}
}

// Save returns a new errfile created or an error
func (r *RepositoryErrFilesCRUD) Save(errfile models.ErrFile) (models.ErrFile, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.ErrFile{}).Create(&errfile).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return errfile, nil
	}
	return models.ErrFile{}, err
}

// UpdateByID update errfile from the DB
func (r *RepositoryErrFilesCRUD) UpdateByID(id string, errfile models.ErrFile) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.ErrFile{}).Where("id = ?", id).Select("*").Updates(&errfile)
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

// DeleteByID errfile by the id
func (r *RepositoryErrFilesCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.ErrFile{}).Where("id = ?", id).Delete(&models.ErrFile{})
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

// FindAll returns all the errfiles from the DB
func (r *RepositoryErrFilesCRUD) FindAll(page int, size int) ([]models.ErrFile, int, error) {
	var err error
	var num int64
	errfiles := []models.ErrFile{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.ErrFile{}).Find(&errfiles)
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&errfiles).Error
		} else {
			err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&errfiles).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return errfiles, int(num), nil
	}
	return nil, 0, err
}

// FindByID return errfile from the DB
func (r *RepositoryErrFilesCRUD) FindByID(id string) (models.ErrFile, error) {
	var err error
	errfile := models.ErrFile{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.ErrFile{}).Where("id = ?", id).Take(&errfile).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return errfile, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.ErrFile{}, errors.New("errfile Not Found")
	}
	return models.ErrFile{}, err
}

// Search errfile from the DB
func (r *RepositoryErrFilesCRUD) Search(q string, page int, size int) ([]models.ErrFile, int, error) {
	var err error
	var num int64
	errfiles := []models.ErrFile{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.ErrFile{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&errfiles).Error
		} else {
			err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&errfiles).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return errfiles, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.ErrFile{}, 0, errors.New("errfiles Not Found")
	}
	return []models.ErrFile{}, 0, err
}

// GetErrFilesByWorkId errfile from the DB
func (r *RepositoryErrFilesCRUD) GetErrFilesByWorkId(id string, page int, size int) ([]models.ErrFile, int, error) {
	var err error
	var num int64
	errfiles := []models.ErrFile{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.ErrFile{}).Where("work_id = ?", id)
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&errfiles).Error
		} else {
			err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&errfiles).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return errfiles, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.ErrFile{}, 0, errors.New("errfiles Not Found")
	}
	return []models.ErrFile{}, 0, err
}
