package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"

	"github.com/gin-gonic/gin"
)

func CreateCastItem(c *gin.Context) {
	castitem := models.CastItem{}
	err := c.ShouldBind(&castitem)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": castitem})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryCastItemsCRUD(db)
	func(castitemRepository repository.CastItemRepository) {
		castitem, err := castitemRepository.Save(castitem)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": castitem})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": castitem})
	}(repo)
}

func DeleteCastItemById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryCastItemsCRUD(db)
	func(castitemRepository repository.CastItemRepository) {
		castitem, err := castitemRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": castitem})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": castitem})
	}(repo)
}

func UpdateCastItemById(c *gin.Context) {
	id := c.Query("id")
	castitem := models.CastItem{}
	err := c.ShouldBind(&castitem)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": castitem})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryCastItemsCRUD(db)
	func(castitemRepository repository.CastItemRepository) {
		castitem, err := castitemRepository.UpdateByID(id, castitem)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": castitem})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": castitem})
	}(repo)
}

func GetCastItemById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryCastItemsCRUD(db)
	func(castitemRepository repository.CastItemRepository) {
		castitem, err := castitemRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": castitem})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": castitem})
	}(repo)
}

func GetCastItemList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryCastItemsCRUD(db)
	func(castitemRepository repository.CastItemRepository) {
		castitems, num, err := castitemRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": castitems, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": castitems, "num": num})
	}(repo)
}

func SearchCastItem(c *gin.Context) {
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
	repo := crud.NewRepositoryCastItemsCRUD(db)
	func(castitemRepository repository.CastItemRepository) {
		castitems, num, err := castitemRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": castitems, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": castitems, "num": num})
	}(repo)
}
