package models

import (
	"time"

	"gorm.io/gorm"
)

// 已播放
type Played struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	UserId    string    `json:"user_id"`
	DataType  string    `json:"data_type"`
	DataId    int       `json:"data_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (d *Played) BeforeCreate(tx *gorm.DB) (err error) {
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
	return
}

func (d *Played) BeforeUpdate(tx *gorm.DB) (err error) {
	d.UpdatedAt = time.Now()
	return
}
