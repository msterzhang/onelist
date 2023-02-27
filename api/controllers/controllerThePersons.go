package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"

	"github.com/gin-gonic/gin"
)

func CreateThePerson(c *gin.Context) {
	theperson := models.ThePerson{}
	err := c.ShouldBind(&theperson)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": theperson})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryThePersonsCRUD(db)
	func(thepersonRepository repository.ThePersonRepository) {
		theperson, err := thepersonRepository.Save(theperson)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": theperson})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": theperson})
	}(repo)
}

func DeleteThePersonById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryThePersonsCRUD(db)
	func(thepersonRepository repository.ThePersonRepository) {
		theperson, err := thepersonRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": theperson})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": theperson})
	}(repo)
}

func UpdateThePersonById(c *gin.Context) {
	id := c.Query("id")
	theperson := models.ThePerson{}
	err := c.ShouldBind(&theperson)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": theperson})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryThePersonsCRUD(db)
	func(thepersonRepository repository.ThePersonRepository) {
		theperson, err := thepersonRepository.UpdateByID(id, theperson)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": theperson})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": theperson})
	}(repo)
}

func GetThePersonById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryThePersonsCRUD(db)
	func(thepersonRepository repository.ThePersonRepository) {
		theperson, err := thepersonRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": theperson})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": theperson})
	}(repo)
}

func GetThePersonList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryThePersonsCRUD(db)
	func(thepersonRepository repository.ThePersonRepository) {
		thepersons, num, err := thepersonRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": thepersons, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": thepersons, "num": num})
	}(repo)
}

func SearchThePerson(c *gin.Context) {
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
	repo := crud.NewRepositoryThePersonsCRUD(db)
	func(thepersonRepository repository.ThePersonRepository) {
		thepersons, num, err := thepersonRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": thepersons, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": thepersons, "num": num})
	}(repo)
}
