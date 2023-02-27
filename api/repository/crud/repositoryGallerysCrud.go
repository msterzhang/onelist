package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"
	"github.com/msterzhang/onelist/config"

	"gorm.io/gorm"
)

// RepositoryGallerysCRUD is the struct for the Gallery CRUD
type RepositoryGallerysCRUD struct {
	db *gorm.DB
}

// NewRepositoryGallerysCRUD returns a new repository with DB connection
func NewRepositoryGallerysCRUD(db *gorm.DB) *RepositoryGallerysCRUD {
	return &RepositoryGallerysCRUD{db}
}

// Save returns a new gallery created or an error
func (r *RepositoryGallerysCRUD) Save(gallery models.Gallery) (models.Gallery, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Gallery{}).Create(&gallery).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return gallery, nil
	}
	return models.Gallery{}, err
}

// UpdateByID update gallery from the DB
func (r *RepositoryGallerysCRUD) UpdateByID(id string, gallery models.Gallery) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Gallery{}).Where("id = ?", id).Select("*").Updates(&gallery)
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

// DeleteByID gallery by the id
func (r *RepositoryGallerysCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Gallery{}).Where("id = ?", id).Delete(&models.Gallery{})
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

// FindAll returns all the gallerys from the DB
func (r *RepositoryGallerysCRUD) FindAll(page int, size int) ([]models.Gallery, int, error) {
	var err error
	var num int64
	gallerys := []models.Gallery{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Gallery{}).Find(&gallerys)
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&gallerys).Error
		} else {
			err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&gallerys).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return gallerys, int(num), nil
	}
	return nil, 0, err
}

// FindByID return gallery from the DB
func (r *RepositoryGallerysCRUD) FindByID(id string) (models.Gallery, error) {
	var err error
	gallery := models.Gallery{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Gallery{}).Where("id = ?", id).Take(&gallery).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return gallery, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Gallery{}, errors.New("gallery Not Found")
	}
	return models.Gallery{}, err
}

// FindByUID return gallery from the DB
func (r *RepositoryGallerysCRUD) FindByUID(id string) (models.Gallery, error) {
	var err error
	gallery := models.Gallery{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.Gallery{}).Where("gallery_uid = ?", id).Take(&gallery).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return gallery, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Gallery{}, errors.New("gallery Not Found")
	}
	return models.Gallery{}, err
}

// Search gallery from the DB
func (r *RepositoryGallerysCRUD) Search(q string, page int, size int) ([]models.Gallery, int, error) {
	var err error
	var num int64
	gallerys := []models.Gallery{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.Gallery{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		if config.DBDRIVER == "sqlite" {
			err = result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&gallerys).Error
		} else {
			err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&gallerys).Error
		}
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return gallerys, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.Gallery{}, 0, errors.New("gallerys Not Found")
	}
	return []models.Gallery{}, 0, err
}
