package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"

	"github.com/gin-gonic/gin"
)

func CreateBelongsToCollection(c *gin.Context) {
	belongstocollection := models.BelongsToCollection{}
	err := c.ShouldBind(&belongstocollection)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": belongstocollection})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryBelongsToCollectionsCRUD(db)
	func(belongstocollectionRepository repository.BelongsToCollectionRepository) {
		belongstocollection, err := belongstocollectionRepository.Save(belongstocollection)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": belongstocollection})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": belongstocollection})
	}(repo)
}

func DeleteBelongsToCollectionById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryBelongsToCollectionsCRUD(db)
	func(belongstocollectionRepository repository.BelongsToCollectionRepository) {
		belongstocollection, err := belongstocollectionRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": belongstocollection})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": belongstocollection})
	}(repo)
}

func UpdateBelongsToCollectionById(c *gin.Context) {
	id := c.Query("id")
	belongstocollection := models.BelongsToCollection{}
	err := c.ShouldBind(&belongstocollection)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": belongstocollection})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryBelongsToCollectionsCRUD(db)
	func(belongstocollectionRepository repository.BelongsToCollectionRepository) {
		belongstocollection, err := belongstocollectionRepository.UpdateByID(id, belongstocollection)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": belongstocollection})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": belongstocollection})
	}(repo)
}

func GetBelongsToCollectionById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryBelongsToCollectionsCRUD(db)
	func(belongstocollectionRepository repository.BelongsToCollectionRepository) {
		belongstocollection, err := belongstocollectionRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": belongstocollection})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": belongstocollection})
	}(repo)
}

func GetBelongsToCollectionList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryBelongsToCollectionsCRUD(db)
	func(belongstocollectionRepository repository.BelongsToCollectionRepository) {
		belongstocollections, num, err := belongstocollectionRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": belongstocollections, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": belongstocollections, "num": num})
	}(repo)
}

func SearchBelongsToCollection(c *gin.Context) {
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
	repo := crud.NewRepositoryBelongsToCollectionsCRUD(db)
	func(belongstocollectionRepository repository.BelongsToCollectionRepository) {
		belongstocollections, num, err := belongstocollectionRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": belongstocollections, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": belongstocollections, "num": num})
	}(repo)
}
