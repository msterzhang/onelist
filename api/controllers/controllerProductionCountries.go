package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"

	"github.com/gin-gonic/gin"
)

func CreateProductionCountrie(c *gin.Context) {
	productioncountrie := models.ProductionCountrie{}
	err := c.ShouldBind(&productioncountrie)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": productioncountrie})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryProductionCountriesCRUD(db)
	func(productioncountrieRepository repository.ProductionCountrieRepository) {
		productioncountrie, err := productioncountrieRepository.Save(productioncountrie)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": productioncountrie})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": productioncountrie})
	}(repo)
}

func DeleteProductionCountrieById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryProductionCountriesCRUD(db)
	func(productioncountrieRepository repository.ProductionCountrieRepository) {
		productioncountrie, err := productioncountrieRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": productioncountrie})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": productioncountrie})
	}(repo)
}

func UpdateProductionCountrieById(c *gin.Context) {
	id := c.Query("id")
	productioncountrie := models.ProductionCountrie{}
	err := c.ShouldBind(&productioncountrie)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": productioncountrie})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryProductionCountriesCRUD(db)
	func(productioncountrieRepository repository.ProductionCountrieRepository) {
		productioncountrie, err := productioncountrieRepository.UpdateByID(id, productioncountrie)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": productioncountrie})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": productioncountrie})
	}(repo)
}

func GetProductionCountrieById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryProductionCountriesCRUD(db)
	func(productioncountrieRepository repository.ProductionCountrieRepository) {
		productioncountrie, err := productioncountrieRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": productioncountrie})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": productioncountrie})
	}(repo)
}

func GetProductionCountrieList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryProductionCountriesCRUD(db)
	func(productioncountrieRepository repository.ProductionCountrieRepository) {
		productioncountries, num, err := productioncountrieRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": productioncountries, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": productioncountries, "num": num})
	}(repo)
}

func SearchProductionCountrie(c *gin.Context) {
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
	repo := crud.NewRepositoryProductionCountriesCRUD(db)
	func(productioncountrieRepository repository.ProductionCountrieRepository) {
		productioncountries, num, err := productioncountrieRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": productioncountries, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": productioncountries, "num": num})
	}(repo)
}
