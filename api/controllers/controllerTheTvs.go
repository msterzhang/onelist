package controllers

import (
	"errors"
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"
	"github.com/msterzhang/onelist/api/service"
	"github.com/msterzhang/onelist/api/utils/dir"
	"github.com/msterzhang/onelist/plugins/alist"
	"github.com/msterzhang/onelist/plugins/thedb"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTheTv(c *gin.Context) {
	thetv := models.TheTv{}
	err := c.ShouldBind(&thetv)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": thetv})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheTvsCRUD(db)
	func(thetvRepository repository.TheTvRepository) {
		thetv, err := thetvRepository.Save(thetv)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": thetv})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": thetv})
	}(repo)
}

func DeleteTheTvById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryTheTvsCRUD(db)
	func(thetvRepository repository.TheTvRepository) {
		thetv, err := thetvRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": thetv})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": thetv})
	}(repo)
}

func UpdateTheTvById(c *gin.Context) {
	id := c.Query("id")
	thetv := models.TheTv{}
	err := c.ShouldBind(&thetv)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": thetv})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheTvsCRUD(db)
	func(thetvRepository repository.TheTvRepository) {
		thetv, err := thetvRepository.UpdateByID(id, thetv)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": thetv})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": thetv})
	}(repo)
}

func GetTheTvById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryTheTvsCRUD(db)
	func(thetvRepository repository.TheTvRepository) {
		thetv, err := thetvRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": thetv})
			return
		}
		thetvNew := service.TheTvService(thetv, c.GetString("UserId"))
		tag := "剧情"
		if len(thetvNew.Genres) > 1 {
			if thetvNew.Genres[0].Name == tag {
				tag = thetvNew.Genres[1].Name
			}
		}
		genre := models.Genre{}
		db.Where("name = ?", tag).Preload("TheTvs", func(db *gorm.DB) *gorm.DB {
			return db.Order("datetime(updated_at) desc").Limit(12)
		}).Find(&genre)
		genre.TheTvs = service.TheTvsService(genre.TheTvs, c.GetString("UserId"))
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": thetvNew, "like": genre.TheTvs})
	}(repo)
}

func GetTheTvList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheTvsCRUD(db)
	func(thetvRepository repository.TheTvRepository) {
		thetvs, num, err := thetvRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": thetvs, "num": num})
			return
		}
		thetvsNew := service.TheTvsService(thetvs, c.GetString("UserId"))
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": thetvsNew, "num": num})
	}(repo)
}

func GetTheTvListByGalleryId(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryTheTvsCRUD(db)
	func(thetvRepository repository.TheTvRepository) {
		thetvs, num, err := thetvRepository.FindByGalleryId(id, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": thetvs, "num": num})
			return
		}
		thetvsNew := service.TheTvsService(thetvs, c.GetString("UserId"))
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": thetvsNew, "num": num})
	}(repo)
}

func SearchTheTv(c *gin.Context) {
	q := c.Query("q")
	if len(q) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": "参数错误!", "data": ""})
		return
	}
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheTvsCRUD(db)
	func(thetvRepository repository.TheTvRepository) {
		thetvs, num, err := thetvRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": thetvs, "num": num})
			return
		}
		thetvsNew := service.TheTvsService(thetvs, c.GetString("UserId"))
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": thetvsNew, "num": num})
	}(repo)
}

func SortTheTv(c *gin.Context) {
	galleryUid := c.Query("gallery_uid")
	if len(galleryUid) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": "参数不足!", "data": ""})
		return
	}
	order := c.Query("order")
	if len(order) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": "参数不足!", "data": ""})
		return
	}
	mode := c.Query("mode")
	if len(mode) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": "参数不足!", "data": ""})
		return
	}
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheTvsCRUD(db)
	func(theTvRepository repository.TheTvRepository) {
		theTvs, num, err := theTvRepository.Sort(galleryUid, mode, order, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": theTvs, "num": num})
			return
		}
		theTvsNew := service.TheTvsService(theTvs, c.GetString("UserId"))
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": theTvsNew, "num": num})
	}(repo)
}

// 手动添加电视
func AddTheTv(c *gin.Context) {
	addVideo := models.AddVideo{}
	err := c.ShouldBind(&addVideo)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "添加资源失败,表单解析出错!", "data": ""})
		return
	}
	gallery := models.Gallery{}
	db := database.NewDb()
	err = db.Model(&models.Gallery{}).Where("gallery_uid = ?", addVideo.GalleryUid).First(&gallery).Error
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "Gallery not found!", "data": ""})
		return
	}
	var files = []string{}
	if gallery.IsAlist {
		files, err = alist.GetAlistFilesPath(addVideo.Path, true, gallery)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": err, "data": ""})
			return
		}
	} else {
		files = dir.GetFilesByPath(addVideo.Path)
	}
	if len(files) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": errors.New("files is 0"), "data": ""})
		return
	}
	go RunTheTvById(addVideo, files, gallery)
	c.JSON(200, gin.H{"code": 200, "msg": "刮削比较耗时，已添加到任务队列，后台运行中，请勿重复提交!", "data": addVideo.TheTvId})
}

func RunTheTvById(addVideo models.AddVideo, files []string, gallery models.Gallery) {
	for _, file := range files {
		_, err := thedb.TheTvDb(addVideo.TheTvId, file, gallery.GalleryUid)
		if err != nil {
			continue
		}
	}
}
