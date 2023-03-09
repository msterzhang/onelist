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
			files, err = alist.GetAlistFilesPath(work.Path, true, gallery)
			if err != nil {
				return
			}
		} else {
			files = dir.GetFilesByPath(work.Path)
		}
		for _, file := range files {
			if gallery.GalleryType == "tv" {
				episode := models.Episode{}
				err := db.Model(&models.Episode{}).Where("url = ?", file).First(&episode).Error
				if errors.Is(err, gorm.ErrRecordNotFound) {
					_, err = thedb.RunTheTvWork(file, gallery.GalleryUid)
					if err != nil {
						continue
					}
				}
			} else {
				themovie := models.TheMovie{}
				err = db.Model(&models.TheMovie{}).Where("url = ?", file).First(&themovie).Error
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
