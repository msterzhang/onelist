package models

import (
	"time"

	"gorm.io/gorm"
)

// 收集错误文件
type ErrFile struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	GalleryUid string    `json:"gallery_uid"`                 //用于关联电影电视到影库
	WorkId     uint       `json:"work_id"`                     //用于任务组
	File       string    `json:"file" gorm:"not null;unique"` //文件
	ErrMsg     string    `json:"err_msg"`                     //错误原因
	IsTv       bool      `json:"is_tv"`                       //是否为电视
	IsOk       bool      `json:"is_ok"`                       //是否刮削完毕
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (e *ErrFile) BeforeCreate(tx *gorm.DB) (err error) {
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return
}

func (e *ErrFile) BeforeUpdate(tx *gorm.DB) (err error) {
	e.UpdatedAt = time.Now()
	return
}
