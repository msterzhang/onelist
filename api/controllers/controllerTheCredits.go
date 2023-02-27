package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"

	"github.com/gin-gonic/gin"
)

func CreateTheCredit(c *gin.Context) {
	thecredit := models.TheCredit{}
	err := c.ShouldBind(&thecredit)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": thecredit})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheCreditsCRUD(db)
	func(thecreditRepository repository.TheCreditRepository) {
		thecredit, err := thecreditRepository.Save(thecredit)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": thecredit})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": thecredit})
	}(repo)
}

func DeleteTheCreditById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryTheCreditsCRUD(db)
	func(thecreditRepository repository.TheCreditRepository) {
		thecredit, err := thecreditRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": thecredit})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": thecredit})
	}(repo)
}

func UpdateTheCreditById(c *gin.Context) {
	id := c.Query("id")
	thecredit := models.TheCredit{}
	err := c.ShouldBind(&thecredit)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": thecredit})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheCreditsCRUD(db)
	func(thecreditRepository repository.TheCreditRepository) {
		thecredit, err := thecreditRepository.UpdateByID(id, thecredit)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": thecredit})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": thecredit})
	}(repo)
}

func GetTheCreditById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryTheCreditsCRUD(db)
	func(thecreditRepository repository.TheCreditRepository) {
		thecredit, err := thecreditRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": thecredit})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": thecredit})
	}(repo)
}

func GetTheCreditList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheCreditsCRUD(db)
	func(thecreditRepository repository.TheCreditRepository) {
		thecredits, num, err := thecreditRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": thecredits, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": thecredits, "num": num})
	}(repo)
}

func SearchTheCredit(c *gin.Context) {
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
	repo := crud.NewRepositoryTheCreditsCRUD(db)
	func(thecreditRepository repository.TheCreditRepository) {
		thecredits, num, err := thecreditRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": thecredits, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": thecredits, "num": num})
	}(repo)
}
