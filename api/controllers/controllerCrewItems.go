package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"

	"github.com/gin-gonic/gin"
)

func CreateCrewItem(c *gin.Context) {
	crewitem := models.CrewItem{}
	err := c.ShouldBind(&crewitem)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": crewitem})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryCrewItemsCRUD(db)
	func(crewitemRepository repository.CrewItemRepository) {
		crewitem, err := crewitemRepository.Save(crewitem)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": crewitem})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": crewitem})
	}(repo)
}

func DeleteCrewItemById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryCrewItemsCRUD(db)
	func(crewitemRepository repository.CrewItemRepository) {
		crewitem, err := crewitemRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": crewitem})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": crewitem})
	}(repo)
}

func UpdateCrewItemById(c *gin.Context) {
	id := c.Query("id")
	crewitem := models.CrewItem{}
	err := c.ShouldBind(&crewitem)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": crewitem})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryCrewItemsCRUD(db)
	func(crewitemRepository repository.CrewItemRepository) {
		crewitem, err := crewitemRepository.UpdateByID(id, crewitem)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": crewitem})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": crewitem})
	}(repo)
}

func GetCrewItemById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryCrewItemsCRUD(db)
	func(crewitemRepository repository.CrewItemRepository) {
		crewitem, err := crewitemRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": crewitem})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": crewitem})
	}(repo)
}

func GetCrewItemList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryCrewItemsCRUD(db)
	func(crewitemRepository repository.CrewItemRepository) {
		crewitems, num, err := crewitemRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": crewitems, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": crewitems, "num": num})
	}(repo)
}

func SearchCrewItem(c *gin.Context) {
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
	repo := crud.NewRepositoryCrewItemsCRUD(db)
	func(crewitemRepository repository.CrewItemRepository) {
		crewitems, num, err := crewitemRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": crewitems, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": crewitems, "num": num})
	}(repo)
}
