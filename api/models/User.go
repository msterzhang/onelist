package models

import (
	"time"

	"github.com/msterzhang/onelist/api/security"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id           uint      `json:"id" gorm:"primaryKey"`
	UserName     string    `json:"user_name"`
	UserId       string    `json:"user_id"`
	UserEmail    string    `json:"user_email"`
	UserPassword string    `json:"user_password"`
	IsAdmin      bool      `json:"is_admin"`
	IsLock       bool      `json:"is_lock"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashPwd, _ := security.Hash(u.UserPassword)
	u.UserPassword = hashPwd
	u.UserId = uuid.New().String()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
