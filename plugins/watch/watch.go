package watch

import (
	"errors"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/dir"
	"github.com/msterzhang/onelist/plugins/alist"
	"github.com/msterzhang/onelist/plugins/thedb"
	"gorm.io/gorm"
)

// 重新查询挂载目录中所有文件，未在影库中找到的就开始刮削
func RunWork(work models.Work) {
	db := database.NewDb()
	if work.Watching {
		gallery := models.Gallery{}
		err := db.Model(&models.Gallery{}).Where("gallery_uid = ?", work.GalleryUid).First(&gallery).Error
		if err != nil {
			return
		}
		var files = []string{}
		if gallery.IsAlist {
			files, err = alist.GetAlistFilesPath(work, gallery)
			if err != nil {
				return
			}
		} else {
			files = dir.GetFilesByPath(work.Path)
		}
		for _, file := range files {
			if gallery.GalleryType == "tv" {
				err := db.Model(&models.Episode{}).Where("url = ?", file).Error
				if errors.Is(err, gorm.ErrRecordNotFound) {
					_, err = thedb.RunTheTvWork(file, gallery.GalleryUid)
					if err != nil {
						continue
					}
				}
			} else {
				err = db.Model(&models.TheTv{}).Where("url = ?", file).Error
				if errors.Is(err, gorm.ErrRecordNotFound) {
					_, err = thedb.RunTheMovieWork(file, gallery.GalleryUid)
					if err != nil {
						continue
					}
				}
			}
		}
	}
}

// 监控
func WatchPath() {
	db := database.NewDb()
	works := []models.Work{}
	err := db.Model(&models.Work{}).Find(&works).Error
	if err != nil {
		return
	}
	for _, work := range works {
		RunWork(work)
	}
}
