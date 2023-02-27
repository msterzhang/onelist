package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"

	"github.com/gin-gonic/gin"
)

func CreateSeason(c *gin.Context) {
	season := models.Season{}
	err := c.ShouldBind(&season)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": season})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositorySeasonsCRUD(db)
	func(seasonRepository repository.SeasonRepository) {
		season, err := seasonRepository.Save(season)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": season})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": season})
	}(repo)
}

func DeleteSeasonById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositorySeasonsCRUD(db)
	func(seasonRepository repository.SeasonRepository) {
		season, err := seasonRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": season})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": season})
	}(repo)
}

func UpdateSeasonById(c *gin.Context) {
	id := c.Query("id")
	season := models.Season{}
	err := c.ShouldBind(&season)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": season})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositorySeasonsCRUD(db)
	func(seasonRepository repository.SeasonRepository) {
		season, err := seasonRepository.UpdateByID(id, season)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": season})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": season})
	}(repo)
}

func GetSeasonById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositorySeasonsCRUD(db)
	func(seasonRepository repository.SeasonRepository) {
		season, err := seasonRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": season})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": season})
	}(repo)
}

func GetSeasonList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositorySeasonsCRUD(db)
	func(seasonRepository repository.SeasonRepository) {
		seasons, num, err := seasonRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": seasons, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": seasons, "num": num})
	}(repo)
}

func SearchSeason(c *gin.Context) {
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
	repo := crud.NewRepositorySeasonsCRUD(db)
	func(seasonRepository repository.SeasonRepository) {
		seasons, num, err := seasonRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": seasons, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": seasons, "num": num})
	}(repo)
}
