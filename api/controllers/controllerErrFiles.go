package controllers

import (
	"path/filepath"
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"
	"github.com/msterzhang/onelist/plugins/alist"
	"github.com/msterzhang/onelist/plugins/thedb"

	"github.com/gin-gonic/gin"
)

func CreateErrFile(c *gin.Context) {
	errfile := models.ErrFile{}
	err := c.ShouldBind(&errfile)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": errfile})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryErrFilesCRUD(db)
	func(errfileRepository repository.ErrFileRepository) {
		errfile, err := errfileRepository.Save(errfile)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": errfile})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": errfile})
	}(repo)
}

func DeleteErrFileById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryErrFilesCRUD(db)
	func(errfileRepository repository.ErrFileRepository) {
		errfile, err := errfileRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": errfile})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": errfile})
	}(repo)
}

func UpdateErrFileById(c *gin.Context) {
	id := c.Query("id")
	errfile := models.ErrFile{}
	err := c.ShouldBind(&errfile)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": errfile})
		return
	}
	db := database.NewDb()
	errfileDb := models.ErrFile{}
	err = db.Model(&models.ErrFile{}).Where("id = ?", errfile.Id).First(&errfileDb).Error
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": err, "data": errfile})
		return
	}
	fileName := filepath.Base(errfile.File)
	err = alist.AlistRnameFile(fileName, errfileDb)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": err, "data": errfile})
		return
	}
	repo := crud.NewRepositoryErrFilesCRUD(db)
	func(errfileRepository repository.ErrFileRepository) {
		errfile, err := errfileRepository.UpdateByID(id, errfile)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": errfile})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": errfile})
	}(repo)
}

func GetErrFileById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryErrFilesCRUD(db)
	func(errfileRepository repository.ErrFileRepository) {
		errfile, err := errfileRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": errfile})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": errfile})
	}(repo)
}

func GetErrFileList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryErrFilesCRUD(db)
	func(errfileRepository repository.ErrFileRepository) {
		errfiles, num, err := errfileRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": errfiles, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": errfiles, "num": num})
	}(repo)
}

func SearchErrFile(c *gin.Context) {
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
	repo := crud.NewRepositoryErrFilesCRUD(db)
	func(errfileRepository repository.ErrFileRepository) {
		errfiles, num, err := errfileRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": errfiles, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": errfiles, "num": num})
	}(repo)
}

func GetErrFilesByWorkId(c *gin.Context) {
	id := c.Query("id")
	if len(id) == 0 {
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
	repo := crud.NewRepositoryErrFilesCRUD(db)
	func(errfileRepository repository.ErrFileRepository) {
		errfiles, num, err := errfileRepository.GetErrFilesByWorkId(id, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": errfiles, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": errfiles, "num": num})
	}(repo)
}

func RefFiles(id string, errfiles []models.ErrFile) {
	db := database.NewDb()
	work := models.Work{}
	db.Model(&models.Work{}).Where("id = ?", id).First(&work)
	work.FileNumber = len(errfiles)
	work.Speed = 0
	err := db.Model(&models.Work{}).Where("id = ?", work.Id).Select("*").Updates(&work).Error
	if err != nil {
		return
	}
	for _, errfile := range errfiles {
		work.Speed++
		db.Model(&models.Work{}).Where("id = ?", work.Id).Select("*").Updates(&work)
		if errfile.IsTv {
			id, err := thedb.RunTheTvWork(errfile.File, errfile.GalleryUid)
			if err == nil && id != 0 {
				db.Model(&models.ErrFile{}).Where("file = ?", errfile.File).Delete(&errfile)
			}
			continue
		}
		id, err := thedb.RunTheMovieWork(errfile.File, errfile.GalleryUid)
		if err == nil && id != 0 {
			db.Model(&models.ErrFile{}).Where("file = ?", errfile.File).Delete(&errfile)
		}
	}
}

func RefErrFilesByWorkId(c *gin.Context) {
	id := c.Query("id")
	if len(id) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": "参数错误!", "data": ""})
		return
	}
	errfiles := []models.ErrFile{}
	db := database.NewDb()
	err := db.Model(&models.ErrFile{}).Where("work_id = ?", id).Find(&errfiles).Error
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": err, "data": ""})
		return
	}

	go RefFiles(id, errfiles)
	c.JSON(200, gin.H{"code": 200, "msg": "提交修复成功!", "data": errfiles})
}

func RefErrFileById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "id not number!", "data": ""})
		return
	}
	errfile := models.ErrFile{}
	err = c.ShouldBind(&errfile)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "表单解析出错!", "data": errfile})
		return
	}
	db := database.NewDb()
	gallery := models.Gallery{}
	err = db.Model(&models.Gallery{}).Where("gallery_uid = ?", errfile.GalleryUid).First(&gallery).Error
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "Gallery not found!", "data": ""})
		return
	}
	if gallery.GalleryType == "tv" {
		thetv, err := thedb.TheTvDb(id, errfile.File, errfile.GalleryUid)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": err})
			return
		}
		if thetv.ID != 0 {
			db.Model(&models.ErrFile{}).Where("file = ?", errfile.File).Delete(&errfile)
		}
		c.JSON(200, gin.H{"code": 200, "msg": "刮削节目成功!", "data": id})
		return
	}
	themovie, err := thedb.TheMovieDb(id, errfile.File, errfile.GalleryUid)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": err})
		return
	}
	if themovie.ID != 0 {
		db.Model(&models.ErrFile{}).Where("file = ?", errfile.File).Delete(&errfile)
	}
	c.JSON(200, gin.H{"code": 200, "msg": "刮削电影成功!", "data": id})
}

func RefErrFileSearch(c *gin.Context) {
	name := c.Query("name")
	dataType := c.Query("type")
	var tv = false
	if dataType == "tv" {
		tv = true
	}
	data, err := thedb.SearchTheDb(name, tv)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": err})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "查询成功!", "data": data})
}

func RefErrVideoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "id not number!", "data": ""})
		return
	}
	oldId, err := strconv.Atoi(c.Query("old_id"))
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "old_id not number!", "data": ""})
		return
	}
	db := database.NewDb()
	galleryType := c.Query("type")
	if galleryType == "movie" {
		themovieDb := models.TheMovie{}
		err = db.Model(&models.TheMovie{}).Where("id = ?", oldId).First(&themovieDb).Error
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有刮削到资源!", "data": err})
			return
		}
		themovieNew, err := thedb.TheMovieDb(id, themovieDb.Url, themovieDb.GalleryUid)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": err})
			return
		}
		if themovieNew.ID != 0 {
			db.Model(&models.TheMovie{}).Where("id = ?", oldId).Delete(&themovieDb)
		}
		c.JSON(200, gin.H{"code": 200, "msg": "刮削电影成功!", "data": themovieNew.ID})
		return
	}
	thetvDb := models.TheTv{}
	err = db.Model(&models.TheTv{}).Where("id = ?", oldId).Preload("TheSeasons").First(&thetvDb).Error
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "没有查询到电视资源!", "data": err})
		return
	}
	s := 0
	for _, season := range thetvDb.TheSeasons {
		seasonId := season.ID
		seasonDb := models.TheSeason{}
		err = db.Model(&models.TheSeason{}).Where("id = ?", seasonId).Preload("Episodes").First(&seasonDb).Error
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到分季资源!", "data": err})
			return
		}
		e := 0
		for _, episode := range seasonDb.Episodes {
			file := episode.Url
			thetv, err := thedb.TheTvDb(id, file, thetvDb.GalleryUid)
			if err != nil {
				continue
			}
			if thetv.ID != 0 {
				db.Model(&models.Episode{}).Where("id = ?", episode.ID).Delete(&models.Episode{})
				e++
			}
		}
		if e == len(seasonDb.Episodes) {
			db.Model(&models.TheSeason{}).Where("id = ?", seasonDb.ID).Delete(&seasonDb)
			s++
		}
	}
	if s == len(thetvDb.TheSeasons) {
		db.Model(&models.TheTv{}).Where("id = ?", thetvDb.ID).Delete(&thetvDb)
	}
	c.JSON(200, gin.H{"code": 200, "msg": "刮削节目成功!", "data": thetvDb.ID})
}
