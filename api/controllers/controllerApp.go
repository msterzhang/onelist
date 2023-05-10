package controllers

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/service"
	"github.com/msterzhang/onelist/api/utils/cache"
	"github.com/msterzhang/onelist/config"
)

func AppIndex(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	thedatas := []models.TheDataIndex{}
	cDb := cache.NewCache()
	VideoData, found := cDb.Get(string(config.SECRETKEY))
	if found {
		err := json.Unmarshal(VideoData.([]byte), &thedatas)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": err.Error(), "data": ""})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": thedatas})
		return
	}
	db := database.NewDb()
	gallerys := []models.Gallery{}
	result := db.Model(&models.Gallery{}).Find(&gallerys)
	if config.DBDRIVER == "sqlite" {
		err := result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&gallerys).Error
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": err.Error(), "data": ""})
			return
		}
	} else {
		err := result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&gallerys).Error
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": err.Error(), "data": ""})
			return
		}
	}
	for _, gallery := range gallerys {
		if gallery.GalleryType == "tv" {
			thetvs := []models.TheTv{}
			result := db.Model(&models.TheTv{}).Find(&thetvs)
			if config.DBDRIVER == "sqlite" {
				err := result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&thetvs).Error
				if err != nil {
					continue
				}
			} else {
				err := result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&thetvs).Error
				if err != nil {
					continue
				}
			}
			thetvsNew := service.TheTvsService(thetvs, c.GetString("UserId"))
			thedata := models.TheDataIndex{Title: gallery.Title, GalleryType: gallery.GalleryType, TheTvList: thetvsNew}
			thedatas = append(thedatas, thedata)
		} else {
			themovies := []models.TheMovie{}
			result := db.Model(&models.TheMovie{}).Find(&themovies)
			if config.DBDRIVER == "sqlite" {
				err := result.Limit(size).Offset((page - 1) * size).Order("datetime(updated_at) desc").Scan(&themovies).Error
				if err != nil {
					continue
				}
			} else {
				err := result.Limit(size).Offset((page - 1) * size).Order("-updated_at").Scan(&themovies).Error
				if err != nil {
					continue
				}
			}
			themoviesNew := service.TheMoviesService(themovies, c.GetString("UserId"))
			thedata := models.TheDataIndex{Title: gallery.Title, GalleryType: gallery.GalleryType, TheMovieList: themoviesNew}
			thedatas = append(thedatas, thedata)
		}
	}
	videoText, _ := json.Marshal(thedatas)
	cDb.Set(string(config.SECRETKEY), videoText, 1*time.Minute)
	c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": thedatas})
}
