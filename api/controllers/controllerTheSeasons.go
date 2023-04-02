package controllers

import (
	"sort"
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"

	"github.com/gin-gonic/gin"
)

func CreateTheSeason(c *gin.Context) {
	theseason := models.TheSeason{}
	err := c.ShouldBind(&theseason)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": theseason})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheSeasonsCRUD(db)
	func(theseasonRepository repository.TheSeasonRepository) {
		theseason, err := theseasonRepository.Save(theseason)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": theseason})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": theseason})
	}(repo)
}

func DeleteTheSeasonById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryTheSeasonsCRUD(db)
	func(theseasonRepository repository.TheSeasonRepository) {
		theseason, err := theseasonRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": theseason})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": theseason})
	}(repo)
}

func UpdateTheSeasonById(c *gin.Context) {
	id := c.Query("id")
	theseason := models.TheSeason{}
	err := c.ShouldBind(&theseason)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": theseason})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheSeasonsCRUD(db)
	func(theseasonRepository repository.TheSeasonRepository) {
		theseason, err := theseasonRepository.UpdateByID(id, theseason)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": theseason})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": theseason})
	}(repo)
}

func GetTheSeasonById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryTheSeasonsCRUD(db)
	func(theseasonRepository repository.TheSeasonRepository) {
		theseason, err := theseasonRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": theseason})
			return
		}
		sort.SliceStable(theseason.Episodes, func(i, j int) bool { return theseason.Episodes[i].EpisodeNumber < theseason.Episodes[j].EpisodeNumber })
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": theseason})
	}(repo)
}

func GetTheSeasonList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheSeasonsCRUD(db)
	func(theseasonRepository repository.TheSeasonRepository) {
		theseasons, num, err := theseasonRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": theseasons, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": theseasons, "num": num})
	}(repo)
}

func SearchTheSeason(c *gin.Context) {
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
	repo := crud.NewRepositoryTheSeasonsCRUD(db)
	func(theseasonRepository repository.TheSeasonRepository) {
		theseasons, num, err := theseasonRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": theseasons, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": theseasons, "num": num})
	}(repo)
}
