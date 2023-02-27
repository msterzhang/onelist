package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"
	"github.com/msterzhang/onelist/api/service"

	"github.com/gin-gonic/gin"
)

func CreateHeart(c *gin.Context) {
	heart := models.Heart{}
	err := c.ShouldBind(&heart)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": heart})
		return
	}
	if len(heart.UserId) == 0 {
		heart.UserId = c.GetString("UserId")
	}
	db := database.NewDb()
	repo := crud.NewRepositoryHeartsCRUD(db)
	func(heartRepository repository.HeartRepository) {
		heart, err := heartRepository.Save(heart)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": heart})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": heart})
	}(repo)
}

func DeleteHeartById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryHeartsCRUD(db)
	func(heartRepository repository.HeartRepository) {
		heart, err := heartRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": heart})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": heart})
	}(repo)
}

func UpdateHeartById(c *gin.Context) {
	id := c.Query("id")
	heart := models.Heart{}
	err := c.ShouldBind(&heart)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": heart})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryHeartsCRUD(db)
	func(heartRepository repository.HeartRepository) {
		heart, err := heartRepository.UpdateByID(id, heart)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": heart})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": heart})
	}(repo)
}

func ReNewHeartByHeart(c *gin.Context) {
	heart := models.Heart{}
	err := c.ShouldBind(&heart)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": heart})
		return
	}
	if len(heart.UserId) == 0 {
		heart.UserId = c.GetString("UserId")
	}
	db := database.NewDb()
	repo := crud.NewRepositoryHeartsCRUD(db)
	func(heartRepository repository.HeartRepository) {
		heart, err := heartRepository.ReNewHeartByHeart(heart)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "处理失败!", "data": heart})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "处理成功!", "data": heart})
	}(repo)
}

func GetHeartById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryHeartsCRUD(db)
	func(heartRepository repository.HeartRepository) {
		heart, err := heartRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": heart})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": heart})
	}(repo)
}

func GetHeartList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryHeartsCRUD(db)
	func(heartRepository repository.HeartRepository) {
		hearts, num, err := heartRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": hearts, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": hearts, "num": num})
	}(repo)
}

func SearchHeart(c *gin.Context) {
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
	repo := crud.NewRepositoryHeartsCRUD(db)
	func(heartRepository repository.HeartRepository) {
		hearts, num, err := heartRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": hearts, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": hearts, "num": num})
	}(repo)
}

func GetHeartDataList(c *gin.Context) {
	dataType := c.Query("data_type")
	if len(dataType) == 0 {
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
	UserId := c.GetString("UserId")
	heart := models.Heart{UserId: UserId, DataType: dataType}
	db := database.NewDb()
	repo := crud.NewRepositoryHeartsCRUD(db)
	func(heartRepository repository.HeartRepository) {
		hearts, num, err := heartRepository.FindAllByUser(heart, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": hearts, "num": num})
			return
		}
		if dataType == "tv" {
			thetvs := []models.TheTv{}
			for _, heartDb := range hearts {
				thetv := models.TheTv{}
				err = db.Model(&models.TheTv{}).Where("id = ?", heartDb.DataId).First(&thetv).Error
				if err != nil {
					continue
				}
				thetv = service.TheTvService(thetv, UserId)
				thetvs = append(thetvs, thetv)
			}
			c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": thetvs, "num": num})
			return
		}
		themovies := []models.TheMovie{}
		for _, heartDb := range hearts {
			themovie := models.TheMovie{}
			err = db.Model(&models.TheMovie{}).Where("id = ?", heartDb.DataId).First(&themovie).Error
			if err != nil {
				continue
			}
			themovie = service.TheMovieService(themovie, UserId)
			themovies = append(themovies, themovie)
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": themovies, "num": num})
	}(repo)
}
