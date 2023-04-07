package controllers

import (
	"errors"
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"
	"github.com/msterzhang/onelist/api/utils/dir"
	"github.com/msterzhang/onelist/plugins/alist"
	"github.com/msterzhang/onelist/plugins/thedb"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func SaveErrFile(file string, errMsg string, galleryUid string, workId uint, isTv bool) {
	db := database.NewDb()
	errFile := models.ErrFile{File: file, GalleryUid: galleryUid, WorkId: workId, IsTv: isTv, ErrMsg: errMsg}
	err := db.Model(&models.ErrFile{}).Create(&errFile).Error
	if err != nil {
		return
	}
}

// 开始刮削任务
func RunWork(files []string, work models.Work, gallery models.Gallery) {
	var err error
	db := database.NewDb()
	for _, file := range files {
		if gallery.GalleryType == "tv" {
			_, err = thedb.RunTheTvWork(file, gallery.GalleryUid)
			if err != nil {
				SaveErrFile(file, err.Error(), gallery.GalleryUid, work.Id, true)
			}
		} else {
			_, err = thedb.RunTheMovieWork(file, gallery.GalleryUid)
			if err != nil {
				SaveErrFile(file, err.Error(), gallery.GalleryUid, work.Id, false)
			}
		}
		work.Speed += 1
		db.Model(&models.Work{}).Where("id = ?", work.Id).Select("*").Updates(&work)
	}
	work.IsOk = true
	db.Model(&models.Work{}).Where("id = ?", work.Id).Select("*").Updates(&work)
}

// 只刮削目录中新增的文件
func RunWorkNew(files []string, work models.Work, gallery models.Gallery) {
	db := database.NewDb()
	var err error
	for _, file := range files {
		if gallery.GalleryType == "tv" {
			episode := models.Episode{}
			err := db.Model(&models.Episode{}).Where("url = ?", file).First(&episode).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				_, err = thedb.RunTheTvWork(file, gallery.GalleryUid)
				if err != nil {
					work.Speed += 1
					continue
				}
			}
		} else {
			themovie := models.TheMovie{}
			err = db.Model(&models.TheMovie{}).Where("url = ?", file).First(&themovie).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				_, err = thedb.RunTheMovieWork(file, gallery.GalleryUid)
				if err != nil {
					work.Speed += 1
					continue
				}
			}
		}
		work.Speed += 1
		db.Model(&models.Work{}).Where("id = ?", work.Id).Select("*").Updates(&work)
	}
	work.IsOk = true
	db.Model(&models.Work{}).Where("id = ?", work.Id).Select("*").Updates(&work)
}

func CreateWork(c *gin.Context) {
	work := models.Work{}
	err := c.ShouldBind(&work)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": work})
		return
	}
	gallery := models.Gallery{}
	db := database.NewDb()
	err = db.Model(&models.Gallery{}).Where("gallery_uid = ?", work.GalleryUid).First(&gallery).Error
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "Gallery not found!", "data": work})
		return
	}
	work.GalleryUid = gallery.GalleryUid
	var files = []string{}
	if gallery.IsAlist {
		files, err = alist.GetAlistFilesPath(work.Path, work.IsRef, gallery)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": err.Error(), "data": work})
			return
		}
	} else {
		files = dir.GetFilesByPath(work.Path)
	}
	if len(files) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": errors.New("files is 0"), "data": work})
		return
	}
	work.FileNumber = len(files)
	err = db.Model(&models.Work{}).Create(&work).Error
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": err.Error(), "data": work})
		return
	}
	go RunWork(files, work, gallery)
	c.JSON(200, gin.H{"code": 200, "msg": "创建刮削任务成功!", "data": work})
}

// 重新刮削
func ReNewWork(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	work := models.Work{}
	err := db.Model(&models.Work{}).Where("id = ?", id).First(&work).Error
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "Work not found!", "data": work})
		return
	}
	gallery := models.Gallery{}
	err = db.Model(&models.Gallery{}).Where("gallery_uid = ?", work.GalleryUid).First(&gallery).Error
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "Gallery not found!", "data": work})
		return
	}
	work.GalleryUid = gallery.GalleryUid
	var files = []string{}
	if gallery.IsAlist {
		files, err = alist.GetAlistFilesPath(work.Path, work.IsRef, gallery)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": err.Error(), "data": ""})
			return
		}
	} else {
		files = dir.GetFilesByPath(work.Path)
	}
	if len(files) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": errors.New("files is 0"), "data": ""})
		return
	}
	work.FileNumber = len(files)
	work.Speed = 0
	err = db.Model(&models.Work{}).Where("id = ?", work.Id).Select("*").Updates(&work).Error
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": err.Error(), "data": ""})
		return
	}
	mod := c.Query("mod")
	if mod == "new" {
		go RunWorkNew(files, work, gallery)
		c.JSON(200, gin.H{"code": 200, "msg": "重启刮削任务成功,只刮削新增文件!", "data": work})
		return
	}
	go RunWork(files, work, gallery)
	c.JSON(200, gin.H{"code": 200, "msg": "重启刮削任务成功!", "data": work})
}

func DeleteWorkById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryWorksCRUD(db)
	func(workRepository repository.WorkRepository) {
		work, err := workRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": work})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": work})
	}(repo)
}

func UpdateWorkById(c *gin.Context) {
	id := c.Query("id")
	work := models.Work{}
	err := c.ShouldBind(&work)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": work})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryWorksCRUD(db)
	func(workRepository repository.WorkRepository) {
		work, err := workRepository.UpdateByID(id, work)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": work})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": work})
	}(repo)
}

func GetWorkById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryWorksCRUD(db)
	func(workRepository repository.WorkRepository) {
		work, err := workRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": work})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": work})
	}(repo)
}

func GetWorkList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryWorksCRUD(db)
	func(workRepository repository.WorkRepository) {
		works, num, err := workRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": works, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": works, "num": num})
	}(repo)
}

func SearchWork(c *gin.Context) {
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
	repo := crud.NewRepositoryWorksCRUD(db)
	func(workRepository repository.WorkRepository) {
		works, num, err := workRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": works, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": works, "num": num})
	}(repo)
}

func GetWorkListByGalleryId(c *gin.Context) {
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
	repo := crud.NewRepositoryWorksCRUD(db)
	func(workRepository repository.WorkRepository) {
		works, num, err := workRepository.GetWorkListByGalleryId(id, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": works, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": works, "num": num})
	}(repo)
}
