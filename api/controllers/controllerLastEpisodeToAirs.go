package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"

	"github.com/gin-gonic/gin"
)

func CreateLastEpisodeToAir(c *gin.Context) {
	lastepisodetoair := models.LastEpisodeToAir{}
	err := c.ShouldBind(&lastepisodetoair)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": lastepisodetoair})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryLastEpisodeToAirsCRUD(db)
	func(lastepisodetoairRepository repository.LastEpisodeToAirRepository) {
		lastepisodetoair, err := lastepisodetoairRepository.Save(lastepisodetoair)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": lastepisodetoair})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": lastepisodetoair})
	}(repo)
}

func DeleteLastEpisodeToAirById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryLastEpisodeToAirsCRUD(db)
	func(lastepisodetoairRepository repository.LastEpisodeToAirRepository) {
		lastepisodetoair, err := lastepisodetoairRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": lastepisodetoair})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": lastepisodetoair})
	}(repo)
}

func UpdateLastEpisodeToAirById(c *gin.Context) {
	id := c.Query("id")
	lastepisodetoair := models.LastEpisodeToAir{}
	err := c.ShouldBind(&lastepisodetoair)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": lastepisodetoair})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryLastEpisodeToAirsCRUD(db)
	func(lastepisodetoairRepository repository.LastEpisodeToAirRepository) {
		lastepisodetoair, err := lastepisodetoairRepository.UpdateByID(id, lastepisodetoair)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": lastepisodetoair})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": lastepisodetoair})
	}(repo)
}

func GetLastEpisodeToAirById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryLastEpisodeToAirsCRUD(db)
	func(lastepisodetoairRepository repository.LastEpisodeToAirRepository) {
		lastepisodetoair, err := lastepisodetoairRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": lastepisodetoair})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": lastepisodetoair})
	}(repo)
}

func GetLastEpisodeToAirList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryLastEpisodeToAirsCRUD(db)
	func(lastepisodetoairRepository repository.LastEpisodeToAirRepository) {
		lastepisodetoairs, num, err := lastepisodetoairRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": lastepisodetoairs, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": lastepisodetoairs, "num": num})
	}(repo)
}

func SearchLastEpisodeToAir(c *gin.Context) {
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
	repo := crud.NewRepositoryLastEpisodeToAirsCRUD(db)
	func(lastepisodetoairRepository repository.LastEpisodeToAirRepository) {
		lastepisodetoairs, num, err := lastepisodetoairRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": lastepisodetoairs, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": lastepisodetoairs, "num": num})
	}(repo)
}
