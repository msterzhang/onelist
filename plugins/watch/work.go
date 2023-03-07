package watch

import (
	"errors"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"gorm.io/gorm"
)

// 自动给影库添加封面
func UpdateGalleryImage() error {
	db := database.NewDb()
	gallerys := []models.Gallery{}
	err := db.Model(&models.Gallery{}).Find(&gallerys).Error
	if err != nil {
		return err
	}
	for _, gallery := range gallerys {
		if len(gallery.Image) == 0 {
			if gallery.GalleryType == "movie" {
				themovie := models.TheMovie{}
				err := db.Model(&models.TheMovie{}).Where("gallery_uid = ?", gallery.GalleryUid).First(&themovie).Error
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil
				}
				if err != nil {
					return err
				}
				gallery.Image = themovie.BackdropPath
				err = db.Model(&models.Gallery{}).Where("id = ?", gallery.Id).Select("*").Updates(&gallery).Error
				if err != nil {
					return err
				}
				continue
			}
			thetv := models.TheTv{}
			err := db.Model(&models.TheTv{}).Where("gallery_uid = ?", gallery.GalleryUid).First(&thetv).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			if err != nil {
				return err
			}
			gallery.Image = thetv.BackdropPath
			err = db.Model(&models.Gallery{}).Where("id = ?", gallery.Id).Select("*").Updates(&gallery).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}
