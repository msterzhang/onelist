package models

import (
	"time"

	"gorm.io/gorm"
)

// 刮削任务模型
type Work struct {
	Id         uint      `json:"id" gorm:"primaryKey"` //ID
	GalleryID  uint      `json:"gallery_id"`           //关联影库
	GalleryUid string    `json:"gallery_uid"`          //用于关联电影电视到影库
	Image      string    `json:"image"`                //图片
	Path       string    `json:"path"`                 //需要被刮削的目录
	FileNumber int       `json:"file_number"`          //文件总数
	Speed      int       `json:"speed"`                //刮削进度
	IsOk       bool      `json:"is_ok"`                //是否刮削完毕
	Watching   bool      `json:"watching"`             //是否监控目录，每天晚上2点自动扫描
	IsRef      bool      `json:"is_ref"`               //是否强制获取alist刷新后的文件列表，不走alist缓存，速度较慢
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (w *Work) BeforeCreate(tx *gorm.DB) (err error) {
	w.CreatedAt = time.Now()
	w.UpdatedAt = time.Now()
	return
}

func (w *Work) BeforeUpdate(tx *gorm.DB) (err error) {
	w.UpdatedAt = time.Now()
	return
}
