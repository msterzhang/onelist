package crud

import (
	"errors"

	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/channels"

	"gorm.io/gorm"
)

// RepositoryUsersCRUD is the struct for the User CRUD
type RepositoryUsersCRUD struct {
	db *gorm.DB
}

// NewRepositoryUsersCRUD returns a new repository with DB connection
func NewRepositoryUsersCRUD(db *gorm.DB) *RepositoryUsersCRUD {
	return &RepositoryUsersCRUD{db}
}

// Save returns a new user created or an error
func (r *RepositoryUsersCRUD) Save(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}

// UpdateByID update user from the DB
func (r *RepositoryUsersCRUD) UpdateByID(id string, user models.User) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.User{}).Where("id = ?", id).Select("*").Updates(&user)
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

// DeleteByID user by the id
func (r *RepositoryUsersCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.User{}).Where("id = ?", id).Delete(&models.User{})
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

// FindAll returns all the users from the DB
func (r *RepositoryUsersCRUD) FindAll(page int, size int) ([]models.User, int, error) {
	var err error
	var num int64
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.User{}).Find(&users)
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-ID").Scan(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return users, int(num), nil
	}
	return nil, 0, err
}

// FindByID return user from the DB
func (r *RepositoryUsersCRUD) FindByID(id string) (models.User, error) {
	var err error
	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.User{}).Where("id = ?", id).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("user Not Found")
	}
	return models.User{}, err
}

// Search user from the DB
func (r *RepositoryUsersCRUD) Search(q string, page int, size int) ([]models.User, int, error) {
	var err error
	var num int64
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result := r.db.Model(&models.User{}).Where("name LIKE ?", "%"+q+"%")
		result.Count(&num)
		err = result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return users, int(num), nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.User{}, 0, errors.New("users Not Found")
	}
	return []models.User{}, 0, err
}
