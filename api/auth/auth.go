package auth

import (
	"errors"

	"github.com/msterzhang/onelist/api/security"
	"github.com/msterzhang/onelist/api/utils/channels"

	"github.com/msterzhang/onelist/api/models"

	"github.com/msterzhang/onelist/api/database"

	"gorm.io/gorm"
)

// SignIn method
func Login(email, password string) (models.User, string, error) {
	user := models.User{}
	var err error
	var db *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		db = database.NewDb()
		if err != nil {
			ch <- false
			return
		}
		err = db.Model(&models.User{}).Where("user_email = ?", email).Take(&user).Error
		if err != nil {
			err = errors.New("用户不存在")
			ch <- false
			return
		}
		if user.IsLock {
			err = errors.New("账号已锁定")
			ch <- false
			return
		}
		err = security.VerifyPassword(user.UserPassword, password)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		user.UserPassword = ""
		err, token := GenerateJWT(user)
		return user, err, token
	}
	return models.User{}, "", err
}

// SignIn method
func LoginAdmin(email, password string) (string, error) {
	user := models.User{}
	var err error
	var db *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		db = database.NewDb()
		if err != nil {
			ch <- false
			return
		}
		err = db.Debug().Model(&models.User{}).Where("user_email = ?", email).Take(&user).Error
		if err != nil {
			err = errors.New("用户不存在")
			ch <- false
			return
		}
		if user.IsLock {
			err = errors.New("账号已锁定，请联系管理员解封")
			ch <- false
			return
		}
		if !user.IsAdmin {
			err = errors.New("非管理员，禁止登录")
			ch <- false
			return
		}
		err = security.VerifyPassword(user.UserPassword, password)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		user.UserPassword = ""
		return GenerateJWT(user)
	}
	return "", err
}
