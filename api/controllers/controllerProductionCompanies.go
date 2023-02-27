package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"

	"github.com/gin-gonic/gin"
)

func CreateProductionCompanie(c *gin.Context) {
	productioncompanie := models.ProductionCompanie{}
	err := c.ShouldBind(&productioncompanie)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": productioncompanie})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryProductionCompaniesCRUD(db)
	func(productioncompanieRepository repository.ProductionCompanieRepository) {
		productioncompanie, err := productioncompanieRepository.Save(productioncompanie)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": productioncompanie})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": productioncompanie})
	}(repo)
}

func DeleteProductionCompanieById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryProductionCompaniesCRUD(db)
	func(productioncompanieRepository repository.ProductionCompanieRepository) {
		productioncompanie, err := productioncompanieRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": productioncompanie})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": productioncompanie})
	}(repo)
}

func UpdateProductionCompanieById(c *gin.Context) {
	id := c.Query("id")
	productioncompanie := models.ProductionCompanie{}
	err := c.ShouldBind(&productioncompanie)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": productioncompanie})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryProductionCompaniesCRUD(db)
	func(productioncompanieRepository repository.ProductionCompanieRepository) {
		productioncompanie, err := productioncompanieRepository.UpdateByID(id, productioncompanie)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": productioncompanie})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": productioncompanie})
	}(repo)
}

func GetProductionCompanieById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryProductionCompaniesCRUD(db)
	func(productioncompanieRepository repository.ProductionCompanieRepository) {
		productioncompanie, err := productioncompanieRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": productioncompanie})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": productioncompanie})
	}(repo)
}

func GetProductionCompanieList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryProductionCompaniesCRUD(db)
	func(productioncompanieRepository repository.ProductionCompanieRepository) {
		productioncompanies, num, err := productioncompanieRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": productioncompanies, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": productioncompanies, "num": num})
	}(repo)
}

func SearchProductionCompanie(c *gin.Context) {
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
	repo := crud.NewRepositoryProductionCompaniesCRUD(db)
	func(productioncompanieRepository repository.ProductionCompanieRepository) {
		productioncompanies, num, err := productioncompanieRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": productioncompanies, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": productioncompanies, "num": num})
	}(repo)
}
