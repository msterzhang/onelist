package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 影库
type Gallery struct {
	Id          uint      `json:"id" gorm:"primaryKey"`         //ID
	Title       string    `json:"title" gorm:"not null;unique"` //标题
	GalleryType string    `json:"gallery_type"`                 //影库类型，电影或者电视
	IsTv        bool      `json:"is_tv"`                        //影库类型，是否是电视
	GalleryUid  string    `json:"gallery_uid"`                  //唯一uid
	Image       string    `json:"image"`                        //图片
	IsAlist     bool      `json:"is_alist"`                     //是否是alist
	AlistHost   string    `json:"alist_host"`                   //alist网站域名
	AlistUser   string    `json:"alist_user"`                   //alist网站管理账号
	AlistPwd    string    `json:"alist_pwd"`                    //alist网站管理密码
	Works       []Work    `json:"works"`                        //添加的目录列表
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (g *Gallery) BeforeCreate(tx *gorm.DB) (err error) {
	g.GalleryUid = uuid.New().String()
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
	return
}

func (g *Gallery) BeforeUpdate(tx *gorm.DB) (err error) {
	g.UpdatedAt = time.Now()
	return
}
