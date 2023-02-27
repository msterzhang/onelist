package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"

	"github.com/gin-gonic/gin"
)

func CreateNextEpisodeToAir(c *gin.Context) {
	nextepisodetoair := models.NextEpisodeToAir{}
	err := c.ShouldBind(&nextepisodetoair)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": nextepisodetoair})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryNextEpisodeToAirsCRUD(db)
	func(nextepisodetoairRepository repository.NextEpisodeToAirRepository) {
		nextepisodetoair, err := nextepisodetoairRepository.Save(nextepisodetoair)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": nextepisodetoair})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": nextepisodetoair})
	}(repo)
}

func DeleteNextEpisodeToAirById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryNextEpisodeToAirsCRUD(db)
	func(nextepisodetoairRepository repository.NextEpisodeToAirRepository) {
		nextepisodetoair, err := nextepisodetoairRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": nextepisodetoair})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": nextepisodetoair})
	}(repo)
}

func UpdateNextEpisodeToAirById(c *gin.Context) {
	id := c.Query("id")
	nextepisodetoair := models.NextEpisodeToAir{}
	err := c.ShouldBind(&nextepisodetoair)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": nextepisodetoair})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryNextEpisodeToAirsCRUD(db)
	func(nextepisodetoairRepository repository.NextEpisodeToAirRepository) {
		nextepisodetoair, err := nextepisodetoairRepository.UpdateByID(id, nextepisodetoair)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": nextepisodetoair})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": nextepisodetoair})
	}(repo)
}

func GetNextEpisodeToAirById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryNextEpisodeToAirsCRUD(db)
	func(nextepisodetoairRepository repository.NextEpisodeToAirRepository) {
		nextepisodetoair, err := nextepisodetoairRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": nextepisodetoair})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": nextepisodetoair})
	}(repo)
}

func GetNextEpisodeToAirList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryNextEpisodeToAirsCRUD(db)
	func(nextepisodetoairRepository repository.NextEpisodeToAirRepository) {
		nextepisodetoairs, num, err := nextepisodetoairRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": nextepisodetoairs, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": nextepisodetoairs, "num": num})
	}(repo)
}

func SearchNextEpisodeToAir(c *gin.Context) {
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
	repo := crud.NewRepositoryNextEpisodeToAirsCRUD(db)
	func(nextepisodetoairRepository repository.NextEpisodeToAirRepository) {
		nextepisodetoairs, num, err := nextepisodetoairRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": nextepisodetoairs, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": nextepisodetoairs, "num": num})
	}(repo)
}
