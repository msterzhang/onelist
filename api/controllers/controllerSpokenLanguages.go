package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"

	"github.com/gin-gonic/gin"
)

func CreateSpokenLanguage(c *gin.Context) {
	spokenlanguage := models.SpokenLanguage{}
	err := c.ShouldBind(&spokenlanguage)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": spokenlanguage})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositorySpokenLanguagesCRUD(db)
	func(spokenlanguageRepository repository.SpokenLanguageRepository) {
		spokenlanguage, err := spokenlanguageRepository.Save(spokenlanguage)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": spokenlanguage})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": spokenlanguage})
	}(repo)
}

func DeleteSpokenLanguageById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositorySpokenLanguagesCRUD(db)
	func(spokenlanguageRepository repository.SpokenLanguageRepository) {
		spokenlanguage, err := spokenlanguageRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": spokenlanguage})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": spokenlanguage})
	}(repo)
}

func UpdateSpokenLanguageById(c *gin.Context) {
	id := c.Query("id")
	spokenlanguage := models.SpokenLanguage{}
	err := c.ShouldBind(&spokenlanguage)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": spokenlanguage})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositorySpokenLanguagesCRUD(db)
	func(spokenlanguageRepository repository.SpokenLanguageRepository) {
		spokenlanguage, err := spokenlanguageRepository.UpdateByID(id, spokenlanguage)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": spokenlanguage})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": spokenlanguage})
	}(repo)
}

func GetSpokenLanguageById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositorySpokenLanguagesCRUD(db)
	func(spokenlanguageRepository repository.SpokenLanguageRepository) {
		spokenlanguage, err := spokenlanguageRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": spokenlanguage})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": spokenlanguage})
	}(repo)
}

func GetSpokenLanguageList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositorySpokenLanguagesCRUD(db)
	func(spokenlanguageRepository repository.SpokenLanguageRepository) {
		spokenlanguages, num, err := spokenlanguageRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": spokenlanguages, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": spokenlanguages, "num": num})
	}(repo)
}

func SearchSpokenLanguage(c *gin.Context) {
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
	repo := crud.NewRepositorySpokenLanguagesCRUD(db)
	func(spokenlanguageRepository repository.SpokenLanguageRepository) {
		spokenlanguages, num, err := spokenlanguageRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": spokenlanguages, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": spokenlanguages, "num": num})
	}(repo)
}
