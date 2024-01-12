package barrage

import (
	"github.com/gin-gonic/gin"
	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"
	"github.com/msterzhang/onelist/plugins/alist"
	"sort"
	"strconv"
	"strings"
)

type barrageGetRequest struct {
	Path string `json:"path"`
}

func Get(c *gin.Context) {
	db := database.NewDb()

	id := c.Query("id")
	season_id := c.Query("season_id")
	gallery_type := c.Query("gallery_type")
	if gallery_type == "tv" {
		tv, err := strconv.Atoi(c.Query("tv"))
		thetvDb := models.TheTv{}
		err = db.Model(&models.TheTv{}).Where("id = ?", id).First(&thetvDb).Error
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!"})
			return
		}

		gallery := models.Gallery{}
		err = db.Model(&models.Gallery{}).Where("gallery_uid = ?", thetvDb.GalleryUid).First(&gallery).Error

		repo := crud.NewRepositoryTheSeasonsCRUD(db)
		func(theseasonRepository repository.TheSeasonRepository) {
			theseason, err := theseasonRepository.FindByID(season_id)
			if err != nil {
				c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": theseason})
				return
			}
			sort.SliceStable(theseason.Episodes, func(i, j int) bool { return theseason.Episodes[i].EpisodeNumber < theseason.Episodes[j].EpisodeNumber })

			path := theseason.Episodes[tv].Url
			path = strings.Replace(path, "/d", "", -1)
			path = strings.Split(path, "?sign")[0]
			paths := strings.Split(path, "/")
			file_name := strings.Split(paths[len(paths)-1], ".")
			barrage_file_name := strings.Join(file_name[:len(file_name)-1], ".") + ".xml"
			path = strings.Join(paths[:len(paths)-1], "/") + "/弹幕/" + barrage_file_name
			file_data, err := alist.AlistFileUrl(gallery, path)
			if err != nil {
				c.JSON(200, gin.H{"code": 201, "msg": "获取弹幕文件出错或此文件无弹幕，" + err.Error(), "data": ""})
				return
			}
			c.JSON(200, gin.H{"code": 200, "msg": "获取弹幕文件成功", "data": file_data, "url": gallery.AlistHost + "/d" + path + "?sign=" + file_data.Data.Sign})
		}(repo)
	}
	if gallery_type == "movie" {
		theMovieDb := models.TheMovie{}
		err := db.Model(&models.TheMovie{}).Where("id = ?", id).First(&theMovieDb).Error
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!"})
			return
		}
		gallery := models.Gallery{}
		err = db.Model(&models.Gallery{}).Where("gallery_uid = ?", theMovieDb.GalleryUid).First(&gallery).Error
		path := theMovieDb.Url
		path = strings.Replace(path, "/d", "", -1)
		path = strings.Split(path, "?sign")[0]
		paths := strings.Split(path, "/")
		file_name := strings.Split(paths[len(paths)-1], ".")
		barrage_file_name := strings.Join(file_name[:len(file_name)-1], ".") + ".xml"
		path = strings.Join(paths[:len(paths)-1], "/") + "/弹幕/" + barrage_file_name
		file_data, err := alist.AlistFileUrl(gallery, path)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "获取弹幕文件出错或此文件无弹幕，" + err.Error(), "data": ""})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "获取弹幕文件成功", "data": file_data, "url": gallery.AlistHost + "/d" + path + "?sign=" + file_data.Data.Sign})

	}
}
